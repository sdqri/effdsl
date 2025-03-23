package effdsl

import (
	"encoding/json"
	"fmt"

	"github.com/sdqri/effdsl/v2/utils"
)

type SortOrder string

const (
	SORT_DEFAULT SortOrder = "default"
	SORT_ASC     SortOrder = "asc"
	SORT_DESC    SortOrder = "desc"
)

const (
	SORT_MISSING_FIRST = "_first"
	SORT_MISSING_LAST  = "_last"
)

type SortMode string

const (
	SORT_MODE_MIN    SortMode = "min"
	SORT_MODE_MAX    SortMode = "max"
	SORT_MODE_SUM    SortMode = "sum"
	SORT_MODE_AVG    SortMode = "avg"
	SORT_MODE_MEDIAN SortMode = "median"
)

type NumericType string

const (
	SORT_NUMERIC_TYPE_DOUBLE     NumericType = "double"
	SORT_NUMERIC_TYPE_LONG       NumericType = "long"
	SORT_NUMERIC_TYPE_DATE       NumericType = "date"
	SORT_NUMERIC_TYPE_DATE_NANOS NumericType = "date_nanos"
)

type NestedSort struct {
	Path     string      `json:"path"`
	Filter   any         `json:"filter,omitempty"`
	MaxChild *int        `json:"max_children,omitempty"`
	Nested   *NestedSort `json:"nested,omitempty"`
}

type SortClauseS struct {
	Field        string       `json:"-"`
	Order        SortOrder    `json:"order,omitempty"`
	Format       *string      `json:"format,omitempty"`
	Missing      *string      `json:"missing,omitempty"`
	Mode         *SortMode    `json:"mode,omitempty"`
	NumericType  *NumericType `json:"numeric_type,omitempty"`
	Nested       *NestedSort  `json:"nested,omitempty"`
	UnmappedType *string      `json:"unmapped_type,omitempty"`
}

func (sq SortClauseS) SortClauseInfo() string {
	return "Sort clause"
}

func (sq SortClauseS) MarshalJSON() ([]byte, error) {
	if sq.Order == "default" && sq.Missing == nil && sq.Mode == nil && sq.NumericType == nil && sq.Nested == nil && sq.UnmappedType == nil && sq.Format == nil {
		return json.Marshal(sq.Field)
	}

	params := M{}

	if sq.Order != "default" {
		params["order"] = sq.Order
	}
	if sq.Missing != nil {
		params["missing"] = *sq.Missing
	}
	if sq.Mode != nil {
		params["mode"] = *sq.Mode
	}
	if sq.NumericType != nil {
		params["numeric_type"] = *sq.NumericType
	}
	if sq.Nested != nil {
		params["nested"] = sq.Nested
	}
	if sq.UnmappedType != nil {
		params["unmapped_type"] = *sq.UnmappedType
	}
	if sq.Format != nil {
		params["format"] = *sq.Format
	}

	tmpM := M{
		sq.Field: params,
	}
	return json.Marshal(tmpM)
}

type sortClauseParameters struct {
	Format       *string      `json:"format,omitempty"`
	Missing      *string      `json:"missing,omitempty"`
	Mode         *SortMode    `json:"mode,omitempty"`
	NumericType  *NumericType `json:"numeric_type,omitempty"`
	Nested       *NestedSort  `json:"nested,omitempty"`
	UnmappedType *string      `json:"unmapped_type,omitempty"`
}

type SortClauseParameter func(params *sortClauseParameters) error

func WithFormat(format string) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		params.Format = &format
		return nil
	}
}

func WithMissing(missing string) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		params.Missing = &missing
		return nil
	}
}

func WithSortMode(mode SortMode) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		params.Mode = &mode
		return nil
	}
}

func WithNumericType(numericType NumericType) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		params.NumericType = &numericType
		return nil
	}
}

type Nested struct {
	Path     string      `json:"path"`
	Filter   QueryResult `json:"filter,omitempty"`
	MaxChild *int        `json:"max_children,omitempty"`
	Nested   *Nested     `json:"nested,omitempty"`
}

const MaxNestedDepth = 5

func (ns Nested) ToNestedSort(currentDepth int) (*NestedSort, error) {
	var nestedSort NestedSort

	// Prevent infinite recursion by checking the current depth
	if currentDepth > MaxNestedDepth {
		return &nestedSort, fmt.Errorf("max nested depth exceeded")
	}

	nestedSort.Path = ns.Path

	if ns.Filter != (QueryResult{}) {
		if err := ns.Filter.Err; err != nil {
			return &nestedSort, err
		}
		nestedSort.Filter = ns.Filter.Ok
	}

	nestedSort.MaxChild = ns.MaxChild

	if ns.Nested != nil {
		subNested, err := ns.Nested.ToNestedSort(currentDepth + 1)
		if err != nil {
			return &nestedSort, err
		}
		nestedSort.Nested = subNested
	}

	return &nestedSort, nil
}

func NewNested(path string, filter QueryResult, maxChild *int, nested *Nested) *Nested {
	nestedProxy := Nested{}
	nestedProxy.Path = path
	if filter != (QueryResult{}) {
		nestedProxy.Filter = filter
	}
	nestedProxy.MaxChild = maxChild
	nestedProxy.Nested = nested
	return &nestedProxy
}

func WithNested(nested *Nested) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		nested, err := nested.ToNestedSort(0)
		if err != nil {
			return err
		}
		params.Nested = nested
		return nil
	}
}

func WithUnmappedType(unmappedType string) SortClauseParameter {
	return func(params *sortClauseParameters) error {
		params.UnmappedType = &unmappedType
		return nil
	}
}

func SortClause(field string, order SortOrder, opts ...SortClauseParameter) SortClauseResult {
	var parameters sortClauseParameters
	sortClause := SortClauseS{}
	for _, opt := range opts {
		err := opt(&parameters)
		if err != nil {
			return SortClauseResult{
				Ok:  nil,
				Err: err,
			}
		}
	}

	sortClause, err := utils.CastStruct[sortClauseParameters, SortClauseS](parameters)
	if err != nil {
		return SortClauseResult{
			Ok:  nil,
			Err: err,
		}
	}
	sortClause.Field = field
	sortClause.Order = order

	return SortClauseResult{
		Ok:  sortClause,
		Err: err,
	}
}
