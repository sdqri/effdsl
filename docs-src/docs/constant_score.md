# Constant Score Query

A constant score query wraps a filter query and returns every matching document with a relevance score equal to the boost parameter value. This query type is useful when you want to apply a uniform score to all documents that match a specific filter query.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	csq "github.com/sdqri/effdsl/queries/constantscore"
	tq "github.com/sdqri/effdsl/queries/termquery"
)

query, err := effdsl.Define(
    csq.ConstantScoreQuery(
        tq.TermQuery("user.id", "kimchy"),
        1.2
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Filter (effdsl.Query)**  
    _(Required, positional)_ The query object that documents must match. This is a required parameter.

*   **Boost (float64)**  
    _(Required, positional)_ A floating-point number used as the constant relevance score for every document matching the filter query. This is a required parameter and defaults to 1.0 if not specified.

### Additional Information

For more details on the constant score query and its parameters, refer to the [official Elasticsearch documentation on constant score queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html).

