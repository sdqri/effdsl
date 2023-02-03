package objects

import "encoding/json"

type QueryStringS struct {
	Query  string   `json:"query"`  //(Required, string) Query string you wish to parse and use for search.
	Fields []string `json:"fields"` //(Optional, array of strings) Array of fields to search. Supports wildcards (*).
}

func (bq QueryStringS) QueryInfo() string {
	return "Query string query"
}

func (qs QueryStringS) MarshalJSON() ([]byte, error) {
	type QueryStringBase QueryStringS
	return json.Marshal(
		map[string]any{
			"query_string": (QueryStringBase)(qs),
		},
	)
}

type QueryStringOption func(*QueryStringS)

func (f DefineType) WithFields(fields ...string) QueryStringOption {
	return func(queryString *QueryStringS) {
		queryString.Fields = fields
	}
}

// Returns documents based on a provided query string, using a parser with a strict syntax
// [Query string query]: https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
func (f DefineType) QueryString(query string, opts ...QueryStringOption) QueryResult {
	queryString := QueryStringS{
		Query: query,
	}
	for _, opt := range opts {
		opt(&queryString)
	}
	return QueryResult{
		Ok:  queryString,
		Err: nil,
	}
}
