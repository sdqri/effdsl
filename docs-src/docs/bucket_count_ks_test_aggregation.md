# Bucket Count K-S Test Aggregation

The bucket count K-S test aggregation compares bucket counts to a provided distribution.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	bucketcountkstest "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketcountkstest"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		bucketcountkstest.BucketCountKSTest(
			"ks_test",
			"latency_ranges>_count",
			bucketcountkstest.WithAlternatives([]string{"less", "greater", "two_sided"}),
			bucketcountkstest.WithSamplingMethod("upper_tail"),
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

*   **BucketsPath (string)**  
    _(Required, positional)_ Path to the `_count` buckets.

*   **WithAlternatives ([]string)**  
    _(Optional, functional option)_ Alternatives (`greater`, `less`, `two_sided`).

*   **WithFractions ([]float64)**  
    _(Optional, functional option)_ Expected distribution fractions.

*   **WithSamplingMethod (string)**  
    _(Optional, functional option)_ Sampling method (`upper_tail`, `uniform`, `lower_tail`).

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

For more details on the bucket count K-S test aggregation and its parameters, refer to the [official Elasticsearch documentation on bucket count K-S test aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-count-ks-test-aggregation).
