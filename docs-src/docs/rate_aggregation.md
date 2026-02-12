# Rate Aggregation

The rate aggregation is a metrics aggregation that calculates a rate of documents or a field in each date-based bucket.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	rate "github.com/sdqri/effdsl/v2/aggregations/metrics/rate"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		rate.Rate(
			"avg_price",
			rate.WithField("price"),
			rate.WithUnit("day"),
			rate.WithMode("sum"),
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

*   **WithField (string)**  
    _(Optional, functional option)_ Field to compute the rate for.

*   **WithUnit (string)**  
    _(Optional, functional option)_ Rate unit (for example, `day`, `month`, `year`).

*   **WithMode (string)**  
    _(Optional, functional option)_ Mode for rate calculation (`sum` or `value_count`).

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

For more details on the rate aggregation and its parameters, refer to the [official Elasticsearch documentation on rate aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-rate-aggregation).
