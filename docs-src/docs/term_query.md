# Term Query

A term query returns documents that contain an exact term in a provided field. The term must exactly match the field value, including whitespace and capitalization.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	tq "github.com/sdqri/effdsl/queries/termquery"
)

query, err := effdsl.Define(
    tq.TermQuery(
        "field_name",
        "exact term",
        tq.WithTQBoost(1.5),
        tq.WithTQCaseInsensitive(true),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field string**  
    The field you wish to search. This is a required parameter.
    
*   **Value string**  
    The term you wish to find in the provided field. This is a required parameter. The term must exactly match the field value, including whitespace and capitalization.
    
*   **WithTQBoost(float64)**  
    Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.
    
*   **WithTQCaseInsensitive(bool)**  
    Allows ASCII case insensitive matching of the value with the indexed field values when set to true. Defaults to false.
    

### Additional Information

For more details on the term query and its parameters, refer to the [official Elasticsearch documentation on term queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html).

