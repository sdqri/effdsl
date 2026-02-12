# Cumulative Sum Aggregation

The cumulative sum aggregation computes the cumulative sum of a metric across buckets.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	cumulativesum "github.com/sdqri/effdsl/v2/aggregations/pipeline/cumulativesum"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		cumulativesum.CumulativeSum(
			"cumulative_sales",
			"sales",
			cumulativesum.WithFormat("0.00"),
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

*   **WithFormat (string)**  
    _(Optional, functional option)_ Decimal format pattern for the output value.

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

For more details on the cumulative sum aggregation and its parameters, refer to the [official Elasticsearch documentation on cumulative sum aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-cumulative-sum-aggregation).
