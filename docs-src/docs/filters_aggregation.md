# Filters Aggregation

The filters aggregation is a multi-bucket aggregation where each bucket is defined by a filter.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	filters "github.com/sdqri/effdsl/v2/aggregations/bucket/filters"
)

filtersMap := map[string]any{
	"errors": map[string]any{"term": map[string]any{"status": "error"}},
	"warnings": map[string]any{"term": map[string]any{"status": "warning"}},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		filters.Filters(
			"statuses",
			filtersMap,
			filters.WithOtherBucket(true),
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

*   **Filters (any)**  
    _(Required, positional)_ Filters definition (map or array).

*   **WithOtherBucket (bool)**  
    _(Optional, functional option)_ Include an "other" bucket.

*   **WithOtherBucketKey (string)**  
    _(Optional, functional option)_ Key for the "other" bucket.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return buckets as a keyed object.

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

For more details on the filters aggregation and its parameters, refer to the [official Elasticsearch documentation on filters aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-filters-aggregation).
