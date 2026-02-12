# Children Aggregation

The children aggregation selects child documents from parent documents.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	children "github.com/sdqri/effdsl/v2/aggregations/bucket/children"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		children.Children(
			"to_children",
			"child",
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

*   **Child Type (string)**  
    _(Required, positional)_ Child type.

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

For more details on the children aggregation and its parameters, refer to the [official Elasticsearch documentation on children aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-children-aggregation).
