# Geo-tile Grid Aggregation

The geo-tile grid aggregation buckets geo points into map tile grid cells.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	geotilegrid "github.com/sdqri/effdsl/v2/aggregations/bucket/geotilegrid"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		geotilegrid.GeoTileGrid(
			"tiles",
			"location",
			7,
			geotilegrid.WithSize(10000),
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
    _(Required, positional)_ Geo field to bucket on.

*   **Precision (int)**  
    _(Required, positional)_ Tile precision.

*   **WithSize (int)**  
    _(Optional, functional option)_ Maximum number of buckets.

*   **WithShardSize (int)**  
    _(Optional, functional option)_ Shard size for bucket collection.

*   **WithBounds (map[string]any)**  
    _(Optional, functional option)_ Bounding box for grid.

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

For more details on the geo-tile grid aggregation and its parameters, refer to the [official Elasticsearch documentation on geo-tile grid aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-geotilegrid-aggregation).
