package objects

import "encoding/json"

// FuzzyQueryS represents a query that matches documents containing terms
// similar to the specified term with a given degree of similarity.
type FuzzyQueryS struct {
	Field          string `json:"-"`                        // (Required, string) The field to search.
	Value          string `json:"value"`                    // (Required, string) The term to search for.
	Fuzziness      string `json:"fuzziness,omitempty"`      // (Optional, string) The degree of fuzziness allowed for the search term (e.g. "AUTO", "1", "2", etc.).
	PrefixLength   int    `json:"prefix_length,omitempty"`  // (Optional, int) The number of initial characters that must match exactly.
	MaxExpansions  int    `json:"max_expansions,omitempty"` // (Optional, int) The maximum number of terms to match.
	Rewrite        string `json:"rewrite,omitempty"`        // (Optional, string) The method used to rewrite the query (e.g. "constant_score_boolean").
	Transpositions bool   `json:"transpositions,omitempty"` // (Optional, bool) Whether to allow transpositions of two adjacent characters.
}

func (fq FuzzyQueryS) QueryInfo() string {
	return "Fuzzy query"
}

func (fq FuzzyQueryS) MarshalJSON() ([]byte, error) {
	type FuzzyQueryBase FuzzyQueryS
	return json.Marshal(
		M{
			"fuzzy": M{
				fq.Field: (FuzzyQueryBase)(fq),
			},
		},
	)
}

type FuzzyQueryOption func(*FuzzyQueryS)

func WithFuzziness(fuzziness string) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Fuzziness = fuzziness
	}
}

func WithPrefixLength(prefixLength int) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.PrefixLength = prefixLength
	}
}

func WithMaxExpansions(maxExpansions int) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.MaxExpansions = maxExpansions
	}
}

func WithFQRewrite(rewrite string) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Rewrite = rewrite
	}
}

func WithTranspositions(transpositions bool) FuzzyQueryOption {
	return func(fuzzyQuery *FuzzyQueryS) {
		fuzzyQuery.Transpositions = transpositions
	}
}

// Returns documents that contain terms similar to the specified term
// with a given degree of similarity.
// [Fuzzy query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
func FuzzyQuery(field string, value string, opts ...FuzzyQueryOption) QueryResult {
	fuzzyQuery := FuzzyQueryS{
		Field: field,
		Value: value,
	}
	for _, opt := range opts {
		opt(&fuzzyQuery)
	}
	return QueryResult{
		Ok:  fuzzyQuery,
		Err: nil,
	}
}
