package objects

import "encoding/json"

type QueryStringS struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}

func (bq QueryStringS) QueryInfo() string {
	return "Query string query"
}

func (qs QueryStringS) MarshalJSON() ([]byte, error) {
	type QueryStringAlias QueryStringS
	return json.Marshal(
		map[string]any{
			"query_string": (QueryStringAlias)(qs),
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
