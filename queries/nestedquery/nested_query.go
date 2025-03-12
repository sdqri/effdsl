package nestedquery

import (
	"encoding/json"
	"fmt"

	"github.com/sdqri/effdsl/v2"
)

// NestedQueryS Wraps another query to search nested fields.
// The nested query searches nested field objects as if they were indexed as separate documents.
// If an object matches the search, the nested query returns the root parent document.
type NestedQueryS struct {
	Path           string       `json:"path"`                      // (Required, string) Path to the nested object you wish to search.
	Query          effdsl.Query `json:"query"`                     // (Required, effdsl.Query) Nested field query.
	ScoreMode      string       `json:"score_mode,omitempty"`      // (Optional, string) Indicates how scores for matching child objects affect the root parent document’s relevance score.
	IgnoreUnmapped bool         `json:"ignore_unmapped,omitempty"` // (Optional, Boolean) Indicates whether to ignore an unmapped path and not return any documents instead of an error.
}

func (nq NestedQueryS) QueryInfo() string {
	return "Nested query"
}

func (nq NestedQueryS) MarshalJSON() ([]byte, error) {
	type NestedQueryBase NestedQueryS
	return json.Marshal(
		effdsl.M{
			"nested": (NestedQueryBase)(nq),
		},
	)
}

type NestedQueryOption func(*NestedQueryS)

func WithNested(path string, qr effdsl.QueryResult, opts ...NestedQueryOption) effdsl.QueryResult {
	if qr.Err != nil {
		return effdsl.QueryResult{
			Ok:  nil,
			Err: fmt.Errorf("nested `%s`: %w", path, qr.Err),
		}
	}

	nq := NestedQueryS{
		Path:  path,
		Query: qr.Ok,
	}

	for _, opt := range opts {
		opt(&nq)
	}

	return effdsl.QueryResult{
		Ok:  nq,
		Err: nil,
	}
}

func WithScoreMode(scoreMode string) NestedQueryOption {
	return func(params *NestedQueryS) {
		params.ScoreMode = scoreMode
	}
}

func WithIgnoreUnmapped(ignoreUnmapped bool) NestedQueryOption {
	return func(params *NestedQueryS) {
		params.IgnoreUnmapped = ignoreUnmapped
	}
}
