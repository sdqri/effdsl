# Normalize Aggregation

The normalize aggregation rescales bucket values using a specified method.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	normalize "github.com/sdqri/effdsl/v2/aggregations/pipeline/normalize"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		normalize.Normalize(
			"percent_of_total_sales",
			"sales",
			"percent_of_sum",
			normalize.WithFormat("00.00%"),
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

*   **Method (string)**  
    _(Required, positional)_ Normalize method (`percent_of_sum`, `rescale_0_1`, `rescale_0_100`, `mean`, `z-score`, `softmax`).

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

For more details on the normalize aggregation and its parameters, refer to the [official Elasticsearch documentation on normalize aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-normalize-aggregation).
