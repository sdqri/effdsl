# Inference Bucket Aggregation

The inference bucket aggregation runs a trained model over bucketed values.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	inferencebucket "github.com/sdqri/effdsl/v2/aggregations/pipeline/inferencebucket"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		inferencebucket.InferenceBucket(
			"malicious_client_ip",
			"malicious_clients_model",
			map[string]string{
				"response_count": "responses_total",
				"url_dc":         "url_dc",
				"bytes_sum":      "bytes_sum",
			},
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

*   **ModelID (string)**  
    _(Required, positional)_ Model ID or alias.

*   **BucketsPath (map[string]string)**  
    _(Required, positional)_ Map of model input fields to bucketed metrics.

*   **WithInferenceConfig (any)**  
    _(Optional, functional option)_ Overrides model inference config.

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

For more details on the inference bucket aggregation and its parameters, refer to the [official Elasticsearch documentation on inference bucket aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-inference-bucket-aggregation).
