# Avg Aggregation

The avg aggregation is a single-value metrics aggregation that computes the average of numeric values extracted from the aggregated documents.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	avg "github.com/sdqri/effdsl/v2/aggregations/metrics/avg"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		avg.Avg(
			"avg_price",
			"price",
			avg.WithMissing(0),
			avg.WithFormat("0.00"),
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
    _(Optional, positional)_ Field to average. Required when script is not set.

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

For more details on the avg aggregation and its parameters, refer to the [official Elasticsearch documentation on avg aggregation](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-avg-aggregation.html).
