# Bucket Correlation Aggregation

The bucket correlation aggregation computes correlations using bucket counts and an indicator.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	bucketcorrelation "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketcorrelation"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		bucketcorrelation.BucketCorrelation(
			"bucket_correlation",
			"latency_ranges>_count",
			bucketcorrelation.CountCorrelation(
				[]float64{0, 52.5, 165, 335, 555},
				200,
				nil,
			),
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

*   **Function (BucketCorrelationFunction)**  
    _(Required, positional)_ Correlation function configuration.

*   **CountCorrelation (expectations []float64, docCount int, fractions []float64)**  
    _(Helper)_ Builds a `count_correlation` function with indicator expectations, total doc count, and optional fractions.

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

For more details on the bucket correlation aggregation and its parameters, refer to the [official Elasticsearch documentation on bucket correlation aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-correlation-aggregation).
