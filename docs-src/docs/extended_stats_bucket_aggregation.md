# Extended Stats Bucket Aggregation

The extended stats bucket aggregation computes extended statistics across buckets.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	extendedstatsbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/extendedstatsbucket"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		extendedstatsbucket.ExtendedStatsBucket(
			"stats_monthly_sales",
			"sales_per_month>sales",
			extendedstatsbucket.WithSigma(2.0),
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
    _(Required, positional)_ Path to the bucketed metric.

*   **WithGapPolicy (string)**  
    _(Optional, functional option)_ Gap policy (`skip`, `insert_zeros`).

*   **WithFormat (string)**  
    _(Optional, functional option)_ Decimal format pattern for the output value.

*   **WithSigma (float64)**  
    _(Optional, functional option)_ Number of standard deviations above/below the mean.

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

For more details on the extended stats bucket aggregation and its parameters, refer to the [official Elasticsearch documentation on extended stats bucket aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-extended-stats-bucket-aggregation).
