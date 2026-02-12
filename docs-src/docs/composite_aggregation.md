# Composite Aggregation

The composite aggregation creates buckets from multiple sources for pagination.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	composite "github.com/sdqri/effdsl/v2/aggregations/bucket/composite"
)

sources := []any{
	map[string]any{"date": map[string]any{"date_histogram": map[string]any{"field": "date", "calendar_interval": "month"}}},
	map[string]any{"type": map[string]any{"terms": map[string]any{"field": "type"}}},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		composite.Composite(
			"buckets",
			sources,
			composite.WithSize(1000),
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

*   **Sources ([]any)**  
    _(Required, positional)_ List of composite sources.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of buckets to return.

*   **WithAfter (map[string]any)**  
    _(Optional, functional option)_ Pagination key.

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

For more details on the composite aggregation and its parameters, refer to the [official Elasticsearch documentation on composite aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-composite-aggregation).
