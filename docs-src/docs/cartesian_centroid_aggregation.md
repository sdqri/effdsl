# Cartesian-centroid Aggregation

The cartesian-centroid aggregation computes the weighted centroid from all coordinate values for point or shape fields.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	cartesiancentroid "github.com/sdqri/effdsl/v2/aggregations/metrics/cartesiancentroid"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		cartesiancentroid.CartesianCentroid(
			"centroid",
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
    _(Optional, positional)_ Point or shape field to compute centroid for.

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

For more details on the cartesian-centroid aggregation and its parameters, refer to the [official Elasticsearch documentation on cartesian-centroid aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-cartesian-centroid-aggregation).
