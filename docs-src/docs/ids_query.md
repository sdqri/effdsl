# IDs Query

Returns documents based on their IDs.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	iq "github.com/sdqri/effdsl/queries/idsquery"
)

query, err := effdsl.Define(
    iq.IDsQuery("1", "4", "100"),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Values ...string**  
    An array of document IDs. This is a required parameter.

### Additional Information

For more details on the IDs query, see the [official Elasticsearch documentation on IDs queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html).

