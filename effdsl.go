package effdsl

import (
	"encoding/json"
)

type M map[string]any

func (m M) MarshalJSON() ([]byte, error) {
	type MBase M
	return json.Marshal((MBase)(m))
}

type SearchBody struct {
	Source       json.Marshaler         `json:"_source,omitempty"`
	From         *uint64                `json:"from,omitempty"`
	Size         *uint64                `json:"size,omitempty"`
	Query        Query                  `json:"query,omitempty"`
	Sort         []SortClauseType       `json:"sort,omitempty"`
	TrackScore   bool                   `json:"track_scores,omitempty"`
	SearchAfter  SearchAfterType        `json:"search_after,omitempty"`
	Collapse     json.Marshaler         `json:"collapse,omitempty"`
	PIT          json.Marshaler         `json:"pit,omitempty"`
	Suggest      Suggest                `json:"suggest,omitempty"`
	Aggregations map[string]Aggregation `json:"aggs,omitempty"`
}

type BodyOption func(*SearchBody) error

// A search request by default executes against the most recent visible data of the target indices, which is called point in time. Elasticsearch pit (point in time) is a lightweight view into the state of the data as it existed when initiated. In some cases, it’s preferred to perform multiple search requests using the same point in time. For example, if refreshes happen between search_after requests, then the results of those requests might not be consistent as changes happening between searches are only visible to the more recent point in time.
// [Point in time]: https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
func WithPIT(id string, keepAlive string) BodyOption {
	pit := PIT(id, keepAlive)
	return func(sb *SearchBody) error {
		sb.PIT = pit
		return nil
	}
}

// By default, searches return the top 10 matching hits. To page through a larger set of results, you can use the search API's from and size parameters. The from parameter defines the number of hits to skip, defaulting to 0. The size parameter is the maximum number of hits to return. Together, these two parameters define a page of results.
// [Paginate search results]: https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html#paginate-search-results
func WithPaginate(from uint64, size uint64) BodyOption {
	return func(sb *SearchBody) error {
		if from != 0 {
			sb.From = &from
		}
		sb.Size = &size
		return nil
	}
}

type SearchAfterType interface {
	SearchAfterInfo() string
	json.Marshaler
}

type SearchAfterResult struct {
	Ok  SearchAfterType
	Err error
}

// By default, you cannot use from and size to page through more than 10,000 hits. This limit is a safeguard set by the index.max_result_window index setting. If you need to page through more than 10,000 hits, use the search_after parameter instead.
// You can use the search_after parameter to retrieve the next page of hits using a set of sort values from the previous page.
// [Search after]: https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html#search-after
func WithSearchAfter(sortValues ...any) BodyOption {
	SearchAfterResult := SearchAfter(sortValues...)
	searchAfter := SearchAfterResult.Ok
	err := SearchAfterResult.Err
	return func(sb *SearchBody) error {
		sb.SearchAfter = searchAfter
		return err
	}
}

type Query interface {
	QueryInfo() string
	json.Marshaler
}

type QueryResult struct {
	Ok  Query
	Err error
}

func WithQuery(queryResult QueryResult) BodyOption {
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
	type MBase M
	return json.Marshal((MBase)(q))
}

func MockQuery(m M) QueryResult {
	return QueryResult{
		Ok:  MockedQuery(m),
		Err: nil,
	}
}

type SortClauseType interface {
	SortClauseInfo() string
	json.Marshaler
}

type SortClauseResult struct {
	Ok  SortClauseType
	Err error
}

// Allows you to add one or more sorts on specific fields. Each sort can be reversed as well. The sort is defined on a per field level, with special field name for _score to sort by score, and _doc to sort by index order.
// [Sort search results]: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-search-results.html#sort-search-results
func WithSort(sortClauseResults ...SortClauseResult) BodyOption {
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

// If scores are not computed. By setting track_scores to true, scores will still be computed and tracked.
// [Track scores]: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-search-results.html#script-based-sortingfunc WithTrackScores() BodyOption {
func WithTrackScores() BodyOption {
	return func(sb *SearchBody) error {
		sb.TrackScore = true
		return nil
	}
}

// Deprecated: use WithCollapse
func WithCollpse(field string) BodyOption {
	return WithCollapse(field)
}

// You can use the collapse parameter to collapse search results based on field values. The collapsing is done by selecting only the top sorted document per collapse key.
// [Collapse search results]: https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html
func WithCollapse(field string) BodyOption {
	searchCollapse := Collapse(field)
	return func(sb *SearchBody) error {
		sb.Collapse = searchCollapse
		return nil
	}
}

// You can use the _source parameter to select what fields of the source are returned. This is called source filtering.
// The following search API request sets the _source request body parameter to false. The document source is not included in the response.
// [Source filtering]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#source-filtering
func WithSourceFilter(opts ...SourceFitlerOption) BodyOption {
	sourceFilter := SourceFilter(opts...)
	return func(sb *SearchBody) error {
		sb.Source = sourceFilter
		return nil
	}
}

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

type Suggest interface {
	SuggestInfo() string
	json.Marshaler
}

type SuggestResult struct {
	Ok  Suggest
	Err error
}

// WithSuggest - allows you to use suggest
func WithSuggest(suggestResult SuggestResult) BodyOption {
	suggest := suggestResult.Ok
	err := suggestResult.Err
	// Type assertion
	return func(b *SearchBody) error {
		b.Suggest = suggest
		return err
	}
}

type Aggregation interface {
	GetAggregationField() string
}

// An aggregation summarizes your data as metrics, statistics, or other analytics.
// [Aggregations]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html
func WithAggregations(aggregations ...Aggregation) BodyOption {
	return func(b *SearchBody) error {
		if b.Aggregations == nil {
			b.Aggregations = map[string]Aggregation{}
		}
		for _, agg := range aggregations {
			b.Aggregations[agg.GetAggregationField()] = agg
		}
		return nil
	}
}
