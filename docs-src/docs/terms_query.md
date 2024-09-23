# Terms Query

A terms query returns documents that contain one or more exact terms in a provided field.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	tsq "github.com/sdqri/effdsl/queries/termsquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        tsq.TermsQuery(
            "user.id",
            []string{"kimchy", "elkbee"},
            tsq.WithBoost(1.0),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field (string)**  
    _(Required, positional)_ The field you wish to search. This is a required parameter.

*   **Values ([]string)**  
    _(Required, positional)_ The array of terms you wish to find in the provided field. This is a required parameter.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
   

### Additional Information

For more details on the terms query and its parameters, refer to the [official Elasticsearch documentation on terms queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html).

