# Regexp Query

A regexp query returns documents that contain terms matching a specified regular expression. The regular expression can include additional options for controlling the match behavior.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	rq "github.com/sdqri/effdsl/queries/regexpquery"
)

query, err := effdsl.Define(
    rq.RegexpQuery(
        "field_name",
        "some_pattern",
        rq.WithFlags("some_flags"),
        rq.WithCaseInsensitive(),
        rq.WithMaxDeterminizedStates(50),
        rq.WithRQRewrite("scoring_boolean"),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

* **Value string**  
    The regular expression pattern to match against the field. This is a required parameter.

* **WithFlags(string)**  
    Additional matching options for the regular expression.

* **WithCaseInsensitive()**  
    If true, the regular expression is case-insensitive.

* **WithMaxDeterminizedStates(int)**  
    The maximum number of automaton states required for the query. Lower values will reduce memory usage but increase query time.

* **WithRQRewrite(string)**  
    The method used to rewrite the query. Valid values include:
    * "constant_score_boolean"
    * "constant_score_filter"
    * "scoring_boolean"
    * "top_terms_boost_N" (where N is the number of top terms)
    * "top_terms_N" (where N is the number of top terms)
    * "random_access_N" (where N is the maximum number of matching terms)

### Additional Information

For more details on the regexp query and its parameters, refer to the [official Elasticsearch documentation on regexp queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html).

