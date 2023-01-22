package objects

import (
	"encoding/json"
)

type M map[string]any

func (m M) MarshalJSON() ([]byte, error) {
	type MAlias M
	return json.Marshal((MAlias)(m))
}

type SearchBody struct {
	Source   json.Marshaler   `json:"_source,omitempty"`
	Query    Query            `json:"query,omitempty"`
	Sort     []SortClauseType `json:"sort,omitempty"`
	Collapse json.Marshaler   `json:"collapse,omitempty"`
}

type BodyOption func(*SearchBody) error

//--------------------------------------------------------------------------------------//
//                                        Query                                         //
//--------------------------------------------------------------------------------------//

type Query interface {
	QueryInfo() string
	json.Marshaler
}

type QueryResult struct {
	Ok  Query
	Err error
}

func (f DefineType) WithQuery(queryResult QueryResult) BodyOption {
	query := queryResult.Ok
	err := queryResult.Err
	// Type assertion
	return func(b *SearchBody) error {
		b.Query = query
		return err
	}
}

type MockedQuery M

func (q MockedQuery) QueryInfo() string {
	return "Mock query"
}

func (q MockedQuery) MarshalJSON() ([]byte, error) {
	type MAlias M
	return json.Marshal((MAlias)(q))
}

func MockQuery(m M) QueryResult {
	return QueryResult{
		Ok:  MockedQuery(m),
		Err: nil,
	}
}

//--------------------------------------------------------------------------------------//
//                                         Sort                                         //
//--------------------------------------------------------------------------------------//

type SortClauseType interface {
	SortClauseInfo() string
	json.Marshaler
}

type SortClauseResult struct {
	Ok  SortClauseType
	Err error
}

func (f DefineType) WithSort(sortClauseResults ...SortClauseResult) BodyOption {
	sortClauses := make([]SortClauseType, 0)
	for _, scr := range sortClauseResults {
		if scr.Err != nil {
			return func(sb *SearchBody) error {
				return scr.Err
			}
		}
		sortClauses = append(sortClauses, scr.Ok)
	}
	return func(sb *SearchBody) error {
		if sb.Sort == nil {
			sb.Sort = sortClauses
		} else {
			sb.Sort = append(sb.Sort, sortClauses...)
		}
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                       Collapse                                       //
//--------------------------------------------------------------------------------------//

func (f DefineType) WithCollpse(field string) BodyOption {
	searchCollapse := Collapse(field)
	return func(sb *SearchBody) error {
		sb.Collapse = searchCollapse
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                   Source filtering                                   //
//--------------------------------------------------------------------------------------//

func (f DefineType) WithSourceFilter(opts ...SourceFitlerOption) BodyOption {
	sourceFilter := SourceFilter(opts...)
	return func(sb *SearchBody) error {
		sb.Source = sourceFilter
		return nil
	}
}

//--------------------------------------------------------------------------------------//
//                                                                                      //
//                                        Define                                        //
//                                                                                      //
//--------------------------------------------------------------------------------------//
type DefineType func(...BodyOption) (*SearchBody, error)

func Define(opts ...BodyOption) (body *SearchBody, err error) {
	body = new(SearchBody)
	for _, opt := range opts {
		err = opt(body)
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

var D DefineType = Define
