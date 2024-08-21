# Exists Query

An exists query returns documents that contain an indexed value for a specified field. This query is useful for checking if a document contains a specific field.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	eq "github.com/sdqri/effdsl/queries/existsquery"
)

query, err := effdsl.Define(
    eq.ExistsQuery("field_name"),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

### Additional Information

For more details on the exists query and its parameters, refer to the [official Elasticsearch documentation on exists queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html).
