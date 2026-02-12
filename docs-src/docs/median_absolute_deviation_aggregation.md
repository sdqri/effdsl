# Median Absolute Deviation Aggregation

The median absolute deviation aggregation approximates the median absolute deviation of numeric values in the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	medianabsolutedeviation "github.com/sdqri/effdsl/v2/aggregations/metrics/medianabsolutedeviation"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		medianabsolutedeviation.MedianAbsoluteDeviation(
			"review_variability",
			"rating",
			medianabsolutedeviation.WithCompression(100),
			medianabsolutedeviation.WithMissing(5),
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
    _(Optional, positional)_ Field to compute the median absolute deviation for.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithCompression (int)**  
    _(Optional, functional option)_ Compression factor controlling accuracy vs. memory usage.

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

For more details on the median absolute deviation aggregation and its parameters, refer to the [official Elasticsearch documentation on median absolute deviation aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-median-absolute-deviation-aggregation).
