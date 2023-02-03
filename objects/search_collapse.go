package objects

import "encoding/json"

type CollapseS struct {
	Field string `json:"field"`
}

func (sc CollapseS) MarshalJSON() ([]byte, error) {
	type SearchCollapseBase CollapseS
	return json.Marshal((SearchCollapseBase)(sc))
}

func Collapse(field string) CollapseS {
	searchCollapse := CollapseS{
		Field: field,
	}
	return searchCollapse
}
