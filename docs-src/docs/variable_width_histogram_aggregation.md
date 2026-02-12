# Variable Width Histogram Aggregation

The variable width histogram aggregation creates buckets with dynamic ranges.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	variablewidthhistogram "github.com/sdqri/effdsl/v2/aggregations/bucket/variablewidthhistogram"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		variablewidthhistogram.VariableWidthHistogram(
			"price_histogram",
			"price",
			5,
			variablewidthhistogram.WithShardSize(100),
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
    _(Required, positional)_ Field to bucket on.

*   **Buckets (int)**  
    _(Required, positional)_ Number of buckets.

*   **WithShardSize (int)**  
    _(Optional, functional option)_ Shard size for bucket collection.

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

For more details on the variable width histogram aggregation and its parameters, refer to the [official Elasticsearch documentation on variable width histogram aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-variablewidthhistogram-aggregation).
