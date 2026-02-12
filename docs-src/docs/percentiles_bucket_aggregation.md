# Percentiles Bucket Aggregation

The percentiles bucket aggregation computes percentiles across buckets of a metric.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	percentilesbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/percentilesbucket"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		percentilesbucket.PercentilesBucket(
			"percentiles_monthly_sales",
			"sales_per_month>sales",
			percentilesbucket.WithPercents([]float64{25, 50, 75}),
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

*   **WithPercents ([]float64)**  
    _(Optional, functional option)_ List of percentiles to calculate.

*   **WithKeyed (bool)**  
    _(Optional, functional option)_ Return results as keyed object instead of array.

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

For more details on the percentiles bucket aggregation and its parameters, refer to the [official Elasticsearch documentation on percentiles bucket aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-percentiles-bucket-aggregation).
