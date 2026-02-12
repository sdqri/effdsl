# String Stats Aggregation

The string stats aggregation computes statistics over string values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	stringstats "github.com/sdqri/effdsl/v2/aggregations/metrics/stringstats"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		stringstats.StringStats(
			"message_stats",
			"message.keyword",
			stringstats.WithShowDistribution(true),
			stringstats.WithMissing("[empty message]"),
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
    _(Optional, positional)_ Field to compute string stats for. Required when script is not set.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script to compute values for the aggregation. Cannot be used together with Field.

*   **WithShowDistribution (bool)**  
    _(Optional, functional option)_ Include character distribution.

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

For more details on the string stats aggregation and its parameters, refer to the [official Elasticsearch documentation on string stats aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-string-stats-aggregation).
