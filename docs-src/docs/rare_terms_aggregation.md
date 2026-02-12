# Rare Terms Aggregation

The rare terms aggregation returns buckets for terms that are rare in the data set.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	rareterms "github.com/sdqri/effdsl/v2/aggregations/bucket/rareterms"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		rareterms.RareTerms(
			"rare_categories",
			"category",
			rareterms.WithMaxDocCount(2),
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
    _(Required, positional)_ Field to bucket on.

*   **WithMaxDocCount (int)**  
    _(Optional, functional option)_ Maximum document count.

*   **WithInclude (any)**  
    _(Optional, functional option)_ Include pattern or values.

*   **WithExclude (any)**  
    _(Optional, functional option)_ Exclude pattern or values.

*   **WithMissing (any)**  
    _(Optional, functional option)_ Default value for missing field values.

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

For more details on the rare terms aggregation and its parameters, refer to the [official Elasticsearch documentation on rare terms aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-bucket-rare-terms-aggregation).
