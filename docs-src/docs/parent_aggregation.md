# Parent Aggregation

The parent aggregation selects parent documents from child documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	parent "github.com/sdqri/effdsl/v2/aggregations/bucket/parent"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		parent.Parent(
			"to_parent",
			"parent",
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

*   **Parent Type (string)**  
    _(Required, positional)_ Parent type.

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

For more details on the parent aggregation and its parameters, refer to the [official Elasticsearch documentation on parent aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-parent-aggregation).
