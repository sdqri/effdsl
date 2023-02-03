package objects

import (
	"encoding/json"
)

type MatchQueryS struct {
	Field string `json:"-"`     //(Required, object) Field you wish to search.
	Query string `json:"query"` //(Required) Text, number, boolean value or date you wish to find in the provided <field>.
}

func (mq MatchQueryS) QueryInfo() string {
	return "Match query"
}

func (mq MatchQueryS) MarshalJSON() ([]byte, error) {
	type MatchQueryBase MatchQueryS
	return json.Marshal(
		M{
			"match": M{
				mq.Field: (MatchQueryBase)(mq),
			},
		},
	)
}

// Returns documents that match a provided text, number, date or boolean value. The provided text is analyzed before matching.
// [Match query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
func (f DefineType) MatchQuery(field string, query string) QueryResult {
	matchQuery := MatchQueryS{
		Field: field,
		Query: query,
	}
	return QueryResult{
		Ok:  matchQuery,
		Err: nil,
	}
}
