# Terms Aggregation

The terms aggregation buckets documents by unique values of a field.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	terms "github.com/sdqri/effdsl/v2/aggregations/bucket/terms"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		terms.Terms(
			"categories",
			"category",
			terms.WithSize(10),
			terms.WithOrder(map[string]any{"_count": "desc"}),
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

*   **WithScript (Script)**  
    _(Optional, functional option)_ Script for term values.

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

*   **WithInclude (any)**  
    _(Optional, functional option)_ Include pattern or values.

*   **WithExclude (any)**  
    _(Optional, functional option)_ Exclude pattern or values.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Default value for missing field values.

*   **WithExecutionHint (string)**  
    _(Optional, functional option)_ Execution hint.

*   **WithCollectMode (string)**  
    _(Optional, functional option)_ Collect mode (`depth_first`, `breadth_first`).

*   **WithShowTermDocCountError (bool)**  
    _(Optional, functional option)_ Include doc count error.

*   **WithValueType (string)**  
    _(Optional, functional option)_ Value type hint.

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

For more details on the terms aggregation and its parameters, refer to the [official Elasticsearch documentation on terms aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-terms-aggregation).
