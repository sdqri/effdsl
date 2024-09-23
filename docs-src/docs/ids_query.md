# IDs Query

Returns documents based on their IDs.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	iq "github.com/sdqri/effdsl/queries/idsquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        iq.IDsQuery("1", "4", "100"),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Values (...string)**  
    _(Required, positional)_ An array of document IDs. This is a required parameter.

### Additional Information

For more details on the IDs query, see the [official Elasticsearch documentation on IDs queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html).

