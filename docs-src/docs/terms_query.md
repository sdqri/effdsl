# Terms Query

A terms query returns documents that contain one or more exact terms in a provided field.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	tsq "github.com/sdqri/effdsl/queries/termsquery"
)

query, err := effdsl.Define(
    tsq.TermsQuery(
        "field_name",
        []string{"term1", "term2"},
        tsq.WithBoost(1.5),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field string**  
    The field you wish to search. This is a required parameter.
    
*   **Values []string**  
    The array of terms you wish to find in the provided field. This is a required parameter.
    
*   **WithBoost(float64)**  
    Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
    

### Additional Information

For more details on the terms query and its parameters, refer to the [official Elasticsearch documentation on terms queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html).
