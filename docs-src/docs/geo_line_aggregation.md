# Geo-line Aggregation

The geo-line aggregation aggregates geo_point values within a bucket into a LineString ordered by a sort field.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	geoline "github.com/sdqri/effdsl/v2/aggregations/metrics/geoline"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		geoline.GeoLine(
			"line",
			"my_location",
			geoline.WithSortField("@timestamp"),
			geoline.WithIncludeSort(true),
			geoline.WithSortOrder("ASC"),
			geoline.WithSize(10000),
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

*   **Point Field (string)**  
    _(Optional, positional)_ Geo point field to build the line from.

*   **WithSortField (string)**  
    _(Optional, functional option)_ Field used to order points in the line.

*   **WithIncludeSort (bool)**  
    _(Optional, functional option)_ Include sort values in the feature properties.

*   **WithSortOrder (string)**  
    _(Optional, functional option)_ Sort order for the line, `ASC` or `DESC`.

*   **WithSize (int)**  
    _(Optional, functional option)_ Maximum number of points in the line.

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

For more details on the geo-line aggregation and its parameters, refer to the [official Elasticsearch documentation on geo-line aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-geo-line).
