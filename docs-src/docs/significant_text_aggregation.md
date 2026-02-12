# Significant Text Aggregation

The significant text aggregation returns significant terms from text fields.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	significanttext "github.com/sdqri/effdsl/v2/aggregations/bucket/significanttext"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		significanttext.SignificantText(
			"significant_description",
			"description",
			significanttext.WithFilterDuplicateText(true),
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

*   **Field (string)**  
    _(Required, positional)_ Field to analyze.

*   **WithFilterDuplicateText (bool)**  
    _(Optional, functional option)_ Filter duplicate text.

*   **WithMinDocCount (int)**  
    _(Optional, functional option)_ Minimum document count.

*   **WithShardMinDocCount (int)**  
    _(Optional, functional option)_ Shard-level minimum document count.

*   **WithSourceFields ([]string)**  
    _(Optional, functional option)_ Source fields to analyze.

*   **WithBackgroundFilter (any)**  
    _(Optional, functional option)_ Background filter.

*   **WithInclude (any)**  
    _(Optional, functional option)_ Include pattern or values.

*   **WithExclude (any)**  
    _(Optional, functional option)_ Exclude pattern or values.

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

For more details on the significant text aggregation and its parameters, refer to the [official Elasticsearch documentation on significant text aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-significanttext-aggregation).
