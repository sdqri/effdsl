package objects

import "encoding/json"

type CollapseS struct {
	Field string `json:"field"`
}

func (sc CollapseS) MarshalJSON() ([]byte, error) {
	type SearchCollapseAlias CollapseS
	return json.Marshal((SearchCollapseAlias)(sc))
}

func Collapse(field string) CollapseS {
	searchCollapse := CollapseS{
		Field: field,
	}
	return searchCollapse
}
