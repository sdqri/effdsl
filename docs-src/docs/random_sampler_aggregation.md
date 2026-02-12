# Random Sampler Aggregation

The random sampler aggregation randomly samples documents in a bucket.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	randomsampler "github.com/sdqri/effdsl/v2/aggregations/bucket/randomsampler"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		randomsampler.RandomSampler(
			"sample",
			0.1,
			randomsampler.WithSeed(42),
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

*   **Probability (float64)**  
    _(Required, positional)_ Sampling probability.

*   **WithSeed (int)**  
    _(Optional, functional option)_ Random seed.

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

For more details on the random sampler aggregation and its parameters, refer to the [official Elasticsearch documentation on random sampler aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-random-sampler-aggregation).
