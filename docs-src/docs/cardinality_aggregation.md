# Cardinality Aggregation

The cardinality aggregation is a single-value metrics aggregation that calculates an approximate count of distinct values.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	cardinality "github.com/sdqri/effdsl/v2/aggregations/metrics/cardinality"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		cardinality.Cardinality(
			"type_count",
			"type",
			cardinality.WithPrecisionThreshold(100),
			cardinality.WithExecutionHint("save_time_heuristic"),
			cardinality.WithMissing("N/A"),
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

*   **Field (string)**  
    _(Optional, positional)_ Field to count distinct values for.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithPrecisionThreshold (int)**  
    _(Optional, functional option)_ Precision threshold that trades memory for accuracy.

*   **WithExecutionHint (string)**  
    _(Optional, functional option)_ Execution mode hint (for example, `save_time_heuristic`).

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

For more details on the cardinality aggregation and its parameters, refer to the [official Elasticsearch documentation on cardinality aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-cardinality-aggregation).
