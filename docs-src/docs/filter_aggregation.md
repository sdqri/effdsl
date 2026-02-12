# Filter Aggregation

The filter aggregation is a single-bucket aggregation that matches documents using a filter.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	filter "github.com/sdqri/effdsl/v2/aggregations/bucket/filter"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		filter.Filter(
			"errors",
			map[string]any{"term": map[string]any{"status": "error"}},
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

*   **Filter (any)**  
    _(Required, positional)_ Filter query.

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

For more details on the filter aggregation and its parameters, refer to the [official Elasticsearch documentation on filter aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-filter-aggregation).
