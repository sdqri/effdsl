# Avg Bucket Aggregation

The avg bucket aggregation calculates the average of a bucketed metric.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	avgbucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/avgbucket"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		avgbucket.AvgBucket(
			"avg_monthly_sales",
			"sales_per_month>sales",
			avgbucket.WithGapPolicy("skip"),
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
    _(Optional, functional option)_ Format string for the output value.

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

For more details on the avg bucket aggregation and its parameters, refer to the [official Elasticsearch documentation on avg bucket aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-avg-bucket-aggregation).
