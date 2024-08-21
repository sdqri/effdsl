# Wildcard Query

A wildcard query returns documents that contain terms matching a wildcard pattern. 

### Example

```go
import (
	"github.com/sdqri/effdsl"
	wq "github.com/sdqri/effdsl/queries/wildcardquery"
)

query, err := effdsl.Define(
    wq.WildcardQuery(
        "field_name",
        "some match query",
        wq.WithBoost(1.0),
        wq.WithRewrite(wq.RewriteParameterConstantScore),
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
    The wildcard pattern for terms you wish to find in the provided field. This is a required parameter.

* **WithBoost(float64)**  
    Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.

* **WithCaseInsensitive()**  
    If true, the wildcard pattern is treated as case-insensitive.

* **WithRewrite(Rewrite)**  
    Method used to rewrite the query. For valid values and more information, see the rewrite parameter.

### Additional Information

For more details on the wildcard query and its parameters, refer to the [official Elasticsearch documentation on wildcard queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html).
