package existsquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

type ExistsQueryS struct {
	Field string `json:"field"` // (Required, object) Field you wish to search.
}

func (eq ExistsQueryS) QueryInfo() string {
	return "Exists query"
}

func (eq ExistsQueryS) MarshalJSON() ([]byte, error) {
	type ExistsQuerySBase ExistsQueryS
	return json.Marshal(
		effdsl.M{
			"exists": (ExistsQuerySBase)(eq),
		},
	)
}

// Returns documents that contain an indexed value for a field.
// [Exists query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
func ExistsQuery(field string) effdsl.QueryResult {
	existsQuery := ExistsQueryS{
		Field: field,
	}
	return effdsl.QueryResult{
		Ok:  existsQuery,
		Err: nil,
	}
}
