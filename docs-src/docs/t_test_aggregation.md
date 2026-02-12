# T-test Aggregation

The t-test aggregation performs a statistical hypothesis test on numeric values extracted from the aggregated documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	ttest "github.com/sdqri/effdsl/v2/aggregations/metrics/ttest"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		ttest.TTest(
			"startup_time_ttest",
			"startup_time_before",
			"startup_time_after",
			ttest.WithType("paired"),
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

*   **A Field (string)**  
    _(Optional, positional)_ Field for population A.

*   **B Field (string)**  
    _(Optional, positional)_ Field for population B.

*   **WithAFilter (map[string]any)**  
    _(Optional, functional option)_ Filter for population A.

*   **WithBFilter (map[string]any)**  
    _(Optional, functional option)_ Filter for population B.

*   **WithType (string)**  
    _(Optional, functional option)_ Test type (`paired`, `homoscedastic`, `heteroscedastic`).

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

For more details on the t-test aggregation and its parameters, refer to the [official Elasticsearch documentation on t-test aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-ttest-aggregation).
