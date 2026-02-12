# Multi Terms Aggregation

The multi terms aggregation buckets documents by multiple terms.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	multiterms "github.com/sdqri/effdsl/v2/aggregations/bucket/multiterms"
)

terms := []any{
	map[string]any{"field": "category"},
	map[string]any{"field": "brand"},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		multiterms.MultiTerms(
			"category_brand",
			terms,
			multiterms.WithSize(10),
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

*   **Terms (any)**  
    _(Required, positional)_ List of term definitions.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of buckets to return.

*   **WithShardSize (int)**  
    _(Optional, functional option)_ Shard size for bucket collection.

*   **WithOrder (map[string]any)**  
    _(Optional, functional option)_ Bucket ordering.

*   **WithMinDocCount (int)**  
    _(Optional, functional option)_ Minimum document count.

*   **WithShardMinDocCount (int)**  
    _(Optional, functional option)_ Shard-level minimum document count.

*   **WithCollectMode (string)**  
    _(Optional, functional option)_ Collect mode.

*   **WithShowTermDocCountError (bool)**  
    _(Optional, functional option)_ Include doc count error.

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

For more details on the multi terms aggregation and its parameters, refer to the [official Elasticsearch documentation on multi terms aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-multi-terms-aggregation).
