# Geo-bounds Aggregation

The geo-bounds aggregation computes the geographic bounding box that contains all values for a geo_point or geo_shape field.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	geobounds "github.com/sdqri/effdsl/v2/aggregations/metrics/geobounds"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		geobounds.GeoBounds(
			"viewport",
			"location",
			geobounds.WithWrapLongitude(true),
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
    _(Optional, positional)_ Geo field to compute bounds for.

*   **WithWrapLongitude (bool)**  
    _(Optional, functional option)_ Whether to wrap longitudes to the -180..180 range.

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

For more details on the geo-bounds aggregation and its parameters, refer to the [official Elasticsearch documentation on geo-bounds aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-geobounds-aggregation).
