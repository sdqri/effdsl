# Moving Percentiles Aggregation

The moving percentiles aggregation computes percentiles across a sliding window of percentile sketches.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	movingpercentiles "github.com/sdqri/effdsl/v2/aggregations/pipeline/movingpercentiles"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		movingpercentiles.MovingPercentiles(
			"the_movperc",
			"the_percentile",
			10,
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
    _(Required, positional)_ Path to the percentile metric.

*   **Window (int)**  
    _(Required, positional)_ Window size.

*   **WithShift (int)**  
    _(Optional, functional option)_ Shift of window position.

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

For more details on the moving percentiles aggregation and its parameters, refer to the [official Elasticsearch documentation on moving percentiles aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-moving-percentiles-aggregation).
