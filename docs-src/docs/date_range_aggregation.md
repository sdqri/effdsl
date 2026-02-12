# Date Range Aggregation

The date range aggregation buckets documents into date ranges.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	daterange "github.com/sdqri/effdsl/v2/aggregations/bucket/daterange"
)

ranges := []daterange.DateRangeItem{
	{To: "now-1M/M"},
	{From: "now-1M/M"},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		daterange.DateRange(
			"recent",
			"date",
			ranges,
			daterange.WithTimeZone("UTC"),
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

*   **Ranges ([]DateRangeItem)**  
    _(Required, positional)_ Ranges to bucket into.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format for bucket keys.

*   **WithTimeZone (string)**  
    _(Optional, functional option)_ Time zone for rounding.

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

For more details on the date range aggregation and its parameters, refer to the [official Elasticsearch documentation on date range aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-daterange-aggregation).
