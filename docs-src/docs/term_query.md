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
        "user.id",
        "kimchy",
        tq.WithBoost(1.5),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field (string)**  
    _(Required, positional)_ The field you wish to search. This is a required parameter.

*   **Value (string)**  
    _(Required, positional)_ The term you wish to find in the provided field. This is a required parameter. The term must exactly match the field value, including whitespace and capitalization.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.

*   **WithCaseInsensitive (bool)**  
    _(Optional, Functional option)_ Allows ASCII case-insensitive matching of the value with the indexed field values when set to true. Defaults to false.
   

### Additional Information

For more details on the term query and its parameters, refer to the [official Elasticsearch documentation on term queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html).

