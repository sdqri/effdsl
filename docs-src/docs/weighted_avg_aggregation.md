# Weighted Avg Aggregation

The weighted avg aggregation computes the weighted average of numeric values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	weightedavg "github.com/sdqri/effdsl/v2/aggregations/metrics/weightedavg"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		weightedavg.WeightedAvg(
			"weighted_grade",
			"grade",
			"weight",
			weightedavg.WithValueMissing(2),
			weightedavg.WithWeightMissing(3),
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

*   **Value Field (string)**  
    _(Optional, positional)_ Field to compute values from.

*   **Weight Field (string)**  
    _(Optional, positional)_ Field to compute weights from.

*   **WithValueField (string)**  
    _(Optional, functional option)_ Field for the value component.

*   **WithValueMissing (any)**  
    _(Optional, functional option)_ Missing value for the value component.

*   **WithValueScript (Script)**  
    _(Optional, functional option)_ Script for the value component.

*   **WithWeightField (string)**  
    _(Optional, functional option)_ Field for the weight component.

*   **WithWeightMissing (any)**  
    _(Optional, functional option)_ Missing value for the weight component.

*   **WithWeightScript (Script)**  
    _(Optional, functional option)_ Script for the weight component.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format string for the output value.

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

For more details on the weighted avg aggregation and its parameters, refer to the [official Elasticsearch documentation on weighted avg aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-weight-avg-aggregation).
