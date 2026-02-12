# Moving Function Aggregation

The moving function aggregation applies a script over a sliding window of bucket values.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	movingfn "github.com/sdqri/effdsl/v2/aggregations/pipeline/movingfn"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		movingfn.MovingFn(
			"the_movfn",
			"the_sum",
			10,
			"MovingFunctions.unweightedAvg(values)",
			movingfn.WithGapPolicy("skip"),
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

*   **Window (int)**  
    _(Required, positional)_ Window size.

*   **Script (string)**  
    _(Required, positional)_ Script to run over the window values.

*   **WithGapPolicy (string)**  
    _(Optional, functional option)_ Gap policy (`skip`, `insert_zeros`).

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

For more details on the moving function aggregation and its parameters, refer to the [official Elasticsearch documentation on moving function aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-movfn-aggregation).
