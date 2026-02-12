# Auto Date Histogram Aggregation

The auto date histogram aggregation creates a date histogram with an automatically selected interval to target a specified number of buckets.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	autodatehistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/autodatehistogram"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		autodatehistogram.AutoDateHistogram(
			"sales_over_time",
			"timestamp",
			10,
			autodatehistogram.WithMinimumInterval("day"),
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

*   **Buckets (int)**  
    _(Required, positional)_ Target number of buckets.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format for bucket keys.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Default value for missing field values.

*   **WithTimeZone (string)**  
    _(Optional, functional option)_ Time zone for rounding.

*   **WithMinimumInterval (string)**  
    _(Optional, functional option)_ Minimum interval to use.

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

For more details on the auto date histogram aggregation and its parameters, refer to the [official Elasticsearch documentation on auto date histogram aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-auto-date-histogram-aggregation).
