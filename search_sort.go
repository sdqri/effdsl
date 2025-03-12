package effdsl

import "encoding/json"

type SortOrder string

const (
	SORT_DEFAULT SortOrder = "default"
	SORT_ASC     SortOrder = "asc"
	SORT_DESC    SortOrder = "desc"
)

type SortClauseS struct {
	Field string    `json:"-"`
	Order SortOrder `json:"-"`
}

func (sq SortClauseS) SortClauseInfo() string {
	return "Sort clause"
}

func (sq SortClauseS) MarshalJSON() ([]byte, error) {
	if sq.Order == SORT_DEFAULT {
		return json.Marshal(sq.Field)
	}
	tmpM := M{
		sq.Field: sq.Order,
	}
	return json.Marshal(tmpM)
}

func SortClause(field string, order SortOrder) SortClauseResult {
	sortClause := SortClauseS{
		Field: field,
		Order: order,
	}
	return SortClauseResult{
		Ok:  sortClause,
		Err: nil,
	}
}
