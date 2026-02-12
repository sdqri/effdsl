# Change Point Aggregation

The change point aggregation detects significant changes in bucketed values.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	changepoint "github.com/sdqri/effdsl/v2/aggregations/pipeline/changepoint"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		changepoint.ChangePoint(
			"change_points_avg",
			"date>avg",
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

*   **BucketsPath (string)**  
    _(Required, positional)_ Path to the bucketed metric.

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

For more details on the change point aggregation and its parameters, refer to the [official Elasticsearch documentation on change point aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-change-point-aggregation).
