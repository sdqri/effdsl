# Top Metrics Aggregation

The top metrics aggregation selects metrics from the document with the largest or smallest sort value.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	topmetrics "github.com/sdqri/effdsl/v2/aggregations/metrics/topmetrics"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		topmetrics.TopMetrics(
			"tm",
			map[string]any{"field": "m"},
			map[string]any{"s": "desc"},
			topmetrics.WithSize(1),
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

*   **Metrics (any)**  
    _(Required, positional)_ Metrics configuration (single metric or list of metrics).

*   **Sort (any)**  
    _(Required, positional)_ Sort configuration.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of top documents to return.

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

For more details on the top metrics aggregation and its parameters, refer to the [official Elasticsearch documentation on top metrics aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-top-metrics).
