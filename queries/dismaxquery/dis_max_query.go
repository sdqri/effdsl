package dismaxquery

import (
	"encoding/json"
	"fmt"

	"github.com/sdqri/effdsl/v2"
)

type DisMaxQueryS struct {
	Queries    []effdsl.Query `json:"queries"`     // (Required) Array of query objects.
	TieBreaker *float64       `json:"tie_breaker"` // (Optional) Floating point number to adjust the relevance scores.
}

func (d DisMaxQueryS) QueryInfo() string {
	return "Disjunction Max query"
}

func (d DisMaxQueryS) MarshalJSON() ([]byte, error) {
	type DisMaxQueryBase DisMaxQueryS
	return json.Marshal(
		map[string]any{
			"dis_max": (DisMaxQueryBase)(d),
		},
	)
}

type DisMaxOption func(*DisMaxQueryS)

func WithTieBreaker(tieBreaker float64) DisMaxOption {
	return func(params *DisMaxQueryS) {
		params.TieBreaker = &tieBreaker
		return
	}
}

// DisMaxQuery returns a query that uses the dis_max query type to find documents matching multiple query clauses.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html
func DisMaxQuery(queryResults []effdsl.QueryResult, opts ...DisMaxOption) effdsl.QueryResult {
	disMaxQuery := DisMaxQueryS{
		Queries: []effdsl.Query{},
	}

	for i, qr := range queryResults {
		if qr.Err != nil {
			return effdsl.QueryResult{
				Ok:  nil,
				Err: fmt.Errorf("error in %dth query: %w", i, qr.Err),
			}
		}
		disMaxQuery.Queries = append(disMaxQuery.Queries, qr.Ok)
	}

	for _, opt := range opts {
		opt(&disMaxQuery)
	}

	return effdsl.QueryResult{
		Ok:  disMaxQuery,
		Err: nil,
	}
}
