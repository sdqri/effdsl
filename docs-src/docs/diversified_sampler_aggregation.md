# Diversified Sampler Aggregation

The diversified sampler aggregation limits documents per value to reduce bias.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	diversifiedsampler "github.com/sdqri/effdsl/v2/aggregations/bucket/diversifiedsampler"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		diversifiedsampler.DiversifiedSampler(
			"sample",
			diversifiedsampler.WithField("user_id"),
			diversifiedsampler.WithShardSize(200),
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

*   **WithShardSize (int)**  
    _(Optional, functional option)_ Shard size for sampling.

*   **WithField (string)**  
    _(Optional, functional option)_ Field used to diversify results.

*   **WithMaxDocsPerValue (int)**  
    _(Optional, functional option)_ Max docs per value.

*   **WithExecutionHint (string)**  
    _(Optional, functional option)_ Execution hint.

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

For more details on the diversified sampler aggregation and its parameters, refer to the [official Elasticsearch documentation on diversified sampler aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-diversified-sampler-aggregation).
