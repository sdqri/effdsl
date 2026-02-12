# Extended Stats Aggregation

The extended stats aggregation is a multi-value metrics aggregation that computes statistics over numeric values from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	extendedstats "github.com/sdqri/effdsl/v2/aggregations/metrics/extendedstats"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		extendedstats.ExtendedStats(
			"grades_stats",
			"grade",
			extendedstats.WithSigma(3),
			extendedstats.WithMissing(0),
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
    _(Optional, positional)_ Field to compute stats for. Required when script is not set.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script to compute values for the aggregation. Cannot be used together with Field.

*   **WithFormat (string)**  
    _(Optional, functional option)_ Format string for the output values.

*   **WithSigma (float64)**  
    _(Optional, functional option)_ Sigma value used to compute standard deviation bounds.

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

For more details on the extended stats aggregation and its parameters, refer to the [official Elasticsearch documentation on extended stats aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-extendedstats-aggregation).
