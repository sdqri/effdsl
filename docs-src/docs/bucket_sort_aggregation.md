# Bucket Sort Aggregation

The bucket sort aggregation sorts and optionally truncates buckets from its parent aggregation.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	bucketsort "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketsort"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		bucketsort.BucketSort(
			"sales_bucket_sort",
			bucketsort.WithSort([]any{
				map[string]any{"total_sales": map[string]any{"order": "desc"}},
			}),
			bucketsort.WithSize(3),
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

*   **WithSort ([]any)**  
    _(Optional, functional option)_ List of sort clauses for the buckets.

*   **WithFrom (int)**  
    _(Optional, functional option)_ Buckets in positions prior to this value will be truncated.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of buckets to return.

*   **WithGapPolicy (string)**  
    _(Optional, functional option)_ Gap policy (`skip`, `insert_zeros`).

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

For more details on the bucket sort aggregation and its parameters, refer to the [official Elasticsearch documentation on bucket sort aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-bucket-sort-aggregation).
