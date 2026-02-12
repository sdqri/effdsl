# Cartesian-bounds Aggregation

The cartesian-bounds aggregation computes the spatial bounding box containing all values for a point or shape field.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	cartesianbounds "github.com/sdqri/effdsl/v2/aggregations/metrics/cartesianbounds"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		cartesianbounds.CartesianBounds(
			"viewport",
			"location",
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
    _(Optional, positional)_ Point or shape field to compute bounds for.

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

For more details on the cartesian-bounds aggregation and its parameters, refer to the [official Elasticsearch documentation on cartesian-bounds aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-cartesian-bounds-aggregation).
