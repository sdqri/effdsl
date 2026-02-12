# Date Histogram Aggregation

The date histogram aggregation buckets documents by date intervals.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	datehistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/datehistogram"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		datehistogram.DateHistogram(
			"sales_over_time",
			"date",
			datehistogram.WithCalendarInterval("month"),
			datehistogram.WithMinDocCount(0),
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
    _(Required, positional)_ Date field to bucket on.

*   **WithCalendarInterval (string)**  
    _(Optional, functional option)_ Calendar interval (`day`, `month`, etc.).

*   **WithFixedInterval (string)**  
    _(Optional, functional option)_ Fixed interval (`1h`, `10m`, etc.).

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format for bucket keys.

*   **WithTimeZone (string)**  
    _(Optional, functional option)_ Time zone for rounding.

*   **WithOffset (string)**  
    _(Optional, functional option)_ Offset for bucket boundaries.

*   **WithOrder (map[string]any)**  
    _(Optional, functional option)_ Order for buckets.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return buckets as a keyed object.

*   **WithMinDocCount (int)**  
    _(Optional, functional option)_ Minimum document count.

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

For more details on the date histogram aggregation and its parameters, refer to the [official Elasticsearch documentation on date histogram aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-datehistogram-aggregation).
