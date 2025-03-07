package effdsl

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2/utils"
)

type SortOrder string

const (
	SORT_DEFAULT SortOrder = "default"
	SORT_ASC     SortOrder = "asc"
	SORT_DESC    SortOrder = "desc"
)

const (
	MISSING_FIRST = "_first"
	MISSING_LAST  = "_last"
)

type SortClauseS struct {
	Field   string    `json:"-"`
	Order   SortOrder `json:"order,omitempty"`
	Missing *string   `json:"missing,omitempty"`
}

func (sq SortClauseS) SortClauseInfo() string {
	return "Sort clause"
}

type sortClauseParameters struct {
	/*
			(Optional, string) The missing parameter specifies how docs which are missing the sort field should be treated:
		 The missing value can be set to _last, _first, or a custom value (that will be used for missing docs as the sort value). The default is _last.
	*/
	Missing *string `json:"missing,omitempty"`
}
type SortClauseParameter func(params *sortClauseParameters)

func (sq SortClauseS) MarshalJSON() ([]byte, error) {
	if sq.Order == SORT_DEFAULT && sq.Missing == nil {
		return json.Marshal(sq.Field)
	}
	params := M{}
	if sq.Order != SORT_DEFAULT {
		params["order"] = sq.Order
	}
	if sq.Missing != nil {
		params["missing"] = *sq.Missing
	}
	tmpM := M{
		sq.Field: params,
	}
	return json.Marshal(tmpM)
}

func WithMissing(missing string) SortClauseParameter {
	return func(params *sortClauseParameters) {
		params.Missing = &missing
	}
}

func SortClause(field string, order SortOrder, params ...SortClauseParameter) SortClauseResult {
	var parameters sortClauseParameters
	sortClause := SortClauseS{}
	for _, prm := range params {
		prm(&parameters)
	}

	sortClause, err := utils.CastStruct[sortClauseParameters, SortClauseS](parameters)
	sortClause.Field = field
	sortClause.Order = order
	return SortClauseResult{
		Ok:  sortClause,
		Err: err,
	}
}
