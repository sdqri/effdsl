# Percentile Ranks Aggregation

The percentile ranks aggregation is a multi-value metrics aggregation that calculates one or more percentile ranks over numeric values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	percentileranks "github.com/sdqri/effdsl/v2/aggregations/metrics/percentileranks"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		percentileranks.PercentileRanks(
			"load_time_ranks",
			"load_time",
			[]float64{500, 600},
			percentileranks.WithKeyed(false),
			percentileranks.WithHDR(3),
			percentileranks.WithMissing(10),
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
    _(Optional, positional)_ Field to compute percentile ranks for. Required when script is not set.

*   **Values ([]float64)**  
    _(Optional, positional)_ Values to compute percentile ranks for.

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

For more details on the percentile ranks aggregation and its parameters, refer to the [official Elasticsearch documentation on percentile ranks aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-percentile-rank-aggregation).
