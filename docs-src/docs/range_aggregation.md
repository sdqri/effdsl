# Range Aggregation

The range aggregation buckets documents into numeric ranges.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	rangeagg "github.com/sdqri/effdsl/v2/aggregations/bucket/rangeagg"
)

ranges := []rangeagg.RangeItem{
	{To: 50},
	{From: 50, To: 100},
	{From: 100},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		rangeagg.Range(
			"price_ranges",
			"price",
			ranges,
			rangeagg.WithKeyed(true),
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

*   **Ranges ([]RangeItem)**  
    _(Required, positional)_ Ranges to bucket into.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return buckets as a keyed object.

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

For more details on the range aggregation and its parameters, refer to the [official Elasticsearch documentation on range aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-range-aggregation).
