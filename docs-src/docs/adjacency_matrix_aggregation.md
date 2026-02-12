# Adjacency Matrix Aggregation

The adjacency matrix aggregation returns a matrix of co-occurring filters.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	adjacencymatrix "github.com/sdqri/effdsl/v2/aggregations/bucket/adjacencymatrix"
)

filters := map[string]any{
	"group_a": map[string]any{"term": map[string]any{"tags": "foo"}},
	"group_b": map[string]any{"term": map[string]any{"tags": "bar"}},
}

query, err := effdsl.Define(
	effdsl.WithAggregation(
		adjacencymatrix.AdjacencyMatrix(
			"interactions",
			filters,
			adjacencymatrix.WithSeparator("&"),
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

*   **Filters (map[string]any)**  
    _(Required, positional)_ Map of named filters.

*   **WithSeparator (string)**  
    _(Optional, functional option)_ Separator for filter names.

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

For more details on the adjacency matrix aggregation and its parameters, refer to the [official Elasticsearch documentation on adjacency matrix aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-adjacency-matrix-aggregation).
