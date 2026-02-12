# Percentiles Aggregation

The percentiles aggregation is a multi-value metrics aggregation that calculates one or more percentiles over numeric values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	percentiles "github.com/sdqri/effdsl/v2/aggregations/metrics/percentiles"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		percentiles.Percentiles(
			"load_time_outlier",
			"load_time",
			[]float64{95, 99, 99.9},
			percentiles.WithKeyed(false),
			percentiles.WithHDR(3),
			percentiles.WithMissing(10),
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
    _(Optional, positional)_ Field to compute percentiles for. Required when script is not set.

*   **Percents ([]float64)**  
    _(Optional, positional)_ Percentiles to calculate.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script to compute values for the aggregation. Cannot be used together with Field.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format string for output values.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Whether to return a keyed response.

*   **WithHDR (int)**  
    _(Optional, functional option)_ Use HDR histogram with the specified significant digits.

*   **WithTDigest (map[string]any)**  
    _(Optional, functional option)_ TDigest configuration options.

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

For more details on the percentiles aggregation and its parameters, refer to the [official Elasticsearch documentation on percentiles aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-percentile-aggregation).
