package objects

import (
	"encoding/json"
)

type ExistsQueryS struct {
	Field string `json:"field"` //(Required, object) Field you wish to search.
}

func (eq ExistsQueryS) QueryInfo() string {
	return "Exists query"
}

func (eq ExistsQueryS) MarshalJSON() ([]byte, error) {
	type ExistsQuerySAlias ExistsQueryS
	return json.Marshal(
		M{
			"exists": (ExistsQuerySAlias)(eq),
		},
	)
}

func (f DefineType) ExistsQuery(field string) QueryResult {
	existsQuery := ExistsQueryS{
		Field: field,
	}
	return QueryResult{
		Ok:  existsQuery,
		Err: nil,
	}
}
