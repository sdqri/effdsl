# Matrix Stats Aggregation

The matrix stats aggregation computes a set of statistics over multiple numeric fields.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	matrixstats "github.com/sdqri/effdsl/v2/aggregations/metrics/matrixstats"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		matrixstats.MatrixStats(
			"statistics",
			[]string{"poverty", "income"},
			matrixstats.WithMode("avg"),
			matrixstats.WithMissing(map[string]any{"income": 50000}),
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

*   **Fields ([]string)**  
    _(Optional, positional)_ Fields to compute statistics for.

*   **WithMissing (map[string]any)**  
    _(Optional, functional option)_ Default values for missing fields.

*   **WithMode (string)**  
    _(Optional, functional option)_ How to handle multi-value fields (`avg`, `min`, `max`, `sum`, `median`).

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

For more details on the matrix stats aggregation and its parameters, refer to the [official Elasticsearch documentation on matrix stats aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-matrix-stats-aggregation).
