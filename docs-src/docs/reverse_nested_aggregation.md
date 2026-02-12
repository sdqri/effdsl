# Reverse Nested Aggregation

The reverse nested aggregation moves from nested documents back to the root document.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	reversenested "github.com/sdqri/effdsl/v2/aggregations/bucket/reversenested"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		reversenested.ReverseNested(
			"back_to_root",
			reversenested.WithPath("comments"),
		),
	),
)

res, err := es.Search(
	es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Name (string)**  
    _(Required, positional)_ Aggregation name.

*   **WithPath (string)**  
    _(Optional, functional option)_ Path to reverse to.

*   **WithSubAggregation (AggregationResult)**  
    _(Optional, functional option)_ Adds an unnamed sub-aggregation.

*   **WithNamedSubAggregation (string, AggregationResult)**  
    _(Optional, functional option)_ Adds a named sub-aggregation.

*   **WithSubAggregationsMap (map[string]AggregationResult)**  
    _(Optional, functional option)_ Adds a map of named sub-aggregations.

*   **WithMetaField (string, any)**  
    _(Optional, functional option)_ Adds a metadata field.

*   **WithMetaMap (map[string]any)**  
    _(Optional, functional option)_ Replaces the metadata map.

### Additional Information

For more details on the reverse nested aggregation and its parameters, refer to the [official Elasticsearch documentation on reverse nested aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-reverse-nested-aggregation).
