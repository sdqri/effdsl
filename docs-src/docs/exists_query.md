# Exists Query

An exists query returns documents that contain an indexed value for a specified field. This query is useful for checking if a document contains a specific field.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	eq "github.com/sdqri/effdsl/queries/existsquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        eq.ExistsQuery("field_name"),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field (string)**  
    _(Required, positional)_ The field you wish to search. This is a required parameter.

### Additional Information

For more details on the exists query and its parameters, refer to the [official Elasticsearch documentation on exists queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html).

