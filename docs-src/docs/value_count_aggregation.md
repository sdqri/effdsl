# Value Count Aggregation

The value count aggregation counts the number of values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	valuecount "github.com/sdqri/effdsl/v2/aggregations/metrics/valuecount"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		valuecount.ValueCount(
			"types_count",
			"type",
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
    _(Optional, positional)_ Field to count values for. Required when script is not set.

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script to compute values for the aggregation. Cannot be used together with Field.

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

For more details on the value count aggregation and its parameters, refer to the [official Elasticsearch documentation on value count aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-valuecount-aggregation).
