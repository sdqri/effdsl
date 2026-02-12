# Cumulative Cardinality Aggregation

The cumulative cardinality aggregation computes the cumulative cardinality of a metric across buckets.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	cumulativecardinality "github.com/sdqri/effdsl/v2/aggregations/pipeline/cumulativecardinality"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		cumulativecardinality.CumulativeCardinality(
			"total_new_users",
			"distinct_users",
			cumulativecardinality.WithFormat("0"),
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
    _(Required, positional)_ Path to the cardinality aggregation.

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

For more details on the cumulative cardinality aggregation and its parameters, refer to the [official Elasticsearch documentation on cumulative cardinality aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-cumulative-cardinality-aggregation).
