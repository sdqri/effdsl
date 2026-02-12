# Geo Distance Aggregation

The geo distance aggregation buckets documents into distance ranges from an origin point.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	geodistance "github.com/sdqri/effdsl/v2/aggregations/bucket/geodistance"
)

to100 := 100.0
from100 := 100.0
to300 := 300.0
from300 := 300.0

ranges := []geodistance.GeoDistanceRange{
	{To: &to100},
	{From: &from100, To: &to300},
	{From: &from300},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		geodistance.GeoDistance(
			"rings",
			"location",
			"52.3760, 4.894",
			ranges,
			geodistance.WithUnit("km"),
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

*   **Origin (any)**  
    _(Required, positional)_ Origin point for distance calculation.

*   **Ranges ([]GeoDistanceRange)**  
    _(Required, positional)_ Distance ranges.

*   **WithUnit (string)**  
    _(Optional, functional option)_ Distance unit (e.g., `km`).

*   **WithDistanceType (string)**  
    _(Optional, functional option)_ Distance type (e.g., `arc`).

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return buckets as a keyed object.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Default value for missing field values.

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

For more details on the geo distance aggregation and its parameters, refer to the [official Elasticsearch documentation on geo distance aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-geo-distance-aggregation).
