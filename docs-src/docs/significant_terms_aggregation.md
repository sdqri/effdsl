# Significant Terms Aggregation

The significant terms aggregation returns terms that are significant in the result set.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	significantterms "github.com/sdqri/effdsl/v2/aggregations/bucket/significantterms"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		significantterms.SignificantTerms(
			"significant_categories",
			"category",
			significantterms.WithHeuristic("chi_square", map[string]any{"include_negatives": false}),
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
    _(Required, positional)_ Field to analyze.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of buckets to return.

*   **WithMinDocCount (int)**  
    _(Optional, functional option)_ Minimum document count.

*   **WithShardMinDocCount (int)**  
    _(Optional, functional option)_ Shard-level minimum document count.

*   **WithInclude (any)**  
    _(Optional, functional option)_ Include pattern or values.

*   **WithExclude (any)**  
    _(Optional, functional option)_ Exclude pattern or values.

*   **WithExecutionHint (string)**  
    _(Optional, functional option)_ Execution hint.

*   **WithBackgroundFilter (any)**  
    _(Optional, functional option)_ Background filter.

*   **WithHeuristic (string, map[string]any)**  
    _(Optional, functional option)_ Significance heuristic.

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

For more details on the significant terms aggregation and its parameters, refer to the [official Elasticsearch documentation on significant terms aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-significantterms-aggregation).
