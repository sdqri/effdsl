# Serial Differencing Aggregation

The serial differencing aggregation computes differences between values separated by a lag.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	serialdiff "github.com/sdqri/effdsl/v2/aggregations/pipeline/serialdiff"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		serialdiff.SerialDiff(
			"thirtieth_difference",
			"the_sum",
			serialdiff.WithLag(30),
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

*   **WithLag (int)**  
    _(Optional, functional option)_ Lag to use for differencing.

*   **WithGapPolicy (string)**  
    _(Optional, functional option)_ Gap policy (`insert_zeros`, `skip`).

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

For more details on the serial differencing aggregation and its parameters, refer to the [official Elasticsearch documentation on serial differencing aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-serialdiff-aggregation).
