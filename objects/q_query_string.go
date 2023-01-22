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
