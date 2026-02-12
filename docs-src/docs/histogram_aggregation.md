# Histogram Aggregation

The histogram aggregation buckets numeric values into fixed-size intervals.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	histogram "github.com/sdqri/effdsl/v2/aggregations/bucket/histogram"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		histogram.Histogram(
			"prices",
			"price",
			50,
			histogram.WithMinDocCount(0),
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
    _(Required, positional)_ Field to bucket on.

*   **Interval (float64)**  
    _(Required, positional)_ Bucket interval.

*   **WithOffset (float64)**  
    _(Optional, functional option)_ Offset for bucket boundaries.

*   **WithMinDocCount (int)**  
    _(Optional, functional option)_ Minimum document count.

*   **WithOrder (map[string]any)**  
    _(Optional, functional option)_ Order for buckets.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return buckets as a keyed object.

*   **WithExtendedBounds (map[string]any)**  
    _(Optional, functional option)_ Extended bounds for buckets.

*   **WithHardBounds (map[string]any)**  
    _(Optional, functional option)_ Hard bounds for buckets.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Default value for missing field values.

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

For more details on the histogram aggregation and its parameters, refer to the [official Elasticsearch documentation on histogram aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-histogram-aggregation).
