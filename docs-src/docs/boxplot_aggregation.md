# Boxplot Aggregation

The boxplot aggregation is a metrics aggregation that computes boxplot values (min, max, median, and quartiles) from numeric values in the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	boxplot "github.com/sdqri/effdsl/v2/aggregations/metrics/boxplot"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		boxplot.Boxplot(
			"load_time_boxplot",
			"load_time",
			boxplot.WithCompression(200),
			boxplot.WithExecutionHint("high_accuracy"),
			boxplot.WithMissing(10),
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
    _(Optional, positional)_ Field to compute the boxplot on.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Value to use when the field is missing.

*   **WithCompression (int)**  
    _(Optional, functional option)_ Compression factor controlling accuracy vs. memory usage.

*   **WithExecutionHint (string)**  
    _(Optional, functional option)_ Execution hint for the TDigest implementation (for example, `high_accuracy`).

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

For more details on the boxplot aggregation and its parameters, refer to the [official Elasticsearch documentation on boxplot aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-boxplot-aggregation).
