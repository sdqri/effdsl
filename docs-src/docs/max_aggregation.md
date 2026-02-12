# Max Aggregation

The max aggregation is a single-value metrics aggregation that computes the maximum of numeric values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	max "github.com/sdqri/effdsl/v2/aggregations/metrics/max"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		max.Max(
			"max_price",
			"price",
			max.WithMissing(0),
			max.WithFormat("0.00"),
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
    _(Optional, positional)_ Field to compute the maximum from. Required when script is not set.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script to compute values for the aggregation. Cannot be used together with Field.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format string for the output value.

*   **WithValueType (string)**  
    _(Optional, functional option)_ Hint about the value type for scripted metrics.

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

For more details on the max aggregation and its parameters, refer to the [official Elasticsearch documentation on max aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-max-aggregation).
