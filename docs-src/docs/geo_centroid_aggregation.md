# Geo-centroid Aggregation

The geo-centroid aggregation computes the weighted centroid from all coordinate values for geo_point or geo_shape fields.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	geocentroid "github.com/sdqri/effdsl/v2/aggregations/metrics/geocentroid"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		geocentroid.GeoCentroid(
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
    _(Optional, positional)_ Geo field to compute centroid for.

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

For more details on the geo-centroid aggregation and its parameters, refer to the [official Elasticsearch documentation on geo-centroid aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-geocentroid-aggregation).
