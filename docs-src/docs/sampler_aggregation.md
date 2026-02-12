# Sampler Aggregation

The sampler aggregation limits the number of documents in a bucket for sub-aggregations.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	sampler "github.com/sdqri/effdsl/v2/aggregations/bucket/sampler"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		sampler.Sampler(
			"sample",
			sampler.WithShardSize(200),
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

For more details on the sampler aggregation and its parameters, refer to the [official Elasticsearch documentation on sampler aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-sampler-aggregation).
