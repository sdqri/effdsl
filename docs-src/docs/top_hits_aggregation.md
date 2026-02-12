# Top Hits Aggregation

The top hits aggregation keeps track of the most relevant document being aggregated and returns regular search hits per bucket.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	tophits "github.com/sdqri/effdsl/v2/aggregations/metrics/tophits"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		tophits.TopHits(
			"top_sales_hits",
			tophits.WithSize(1),
			tophits.WithSort([]any{map[string]any{"date": map[string]any{"order": "desc"}}}),
			tophits.WithSource(map[string]any{"includes": []string{"date", "price"}}),
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

*   **WithFrom (int)**  
    _(Optional, functional option)_ Offset from the first result.

*   **WithSize (int)**  
    _(Optional, functional option)_ Number of top hits to return.

*   **WithSort ([]any)**  
    _(Optional, functional option)_ Sort definition for top hits.

*   **WithSource (any)**  
    _(Optional, functional option)_ Source filtering definition.

*   **WithStoredFields ([]string)**  
    _(Optional, functional option)_ Stored fields to return.

*   **WithDocvalueFields ([]any)**  
    _(Optional, functional option)_ Docvalue fields to return.

*   **WithScriptFields (map[string]any)**  
    _(Optional, functional option)_ Script fields to return.

*   **WithHighlight (map[string]any)**  
    _(Optional, functional option)_ Highlighting configuration.

*   **WithExplain (bool)**  
    _(Optional, functional option)_ Include explanation of scoring.

*   **WithTrackScores (bool)**  
    _(Optional, functional option)_ Track scores for each hit.

*   **WithVersion (bool)**  
    _(Optional, functional option)_ Include document version.

*   **WithSeqNoPrimaryTerm (bool)**  
    _(Optional, functional option)_ Include sequence number and primary term.

*   **WithFields ([]any)**  
    _(Optional, functional option)_ Fields to return.

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

For more details on the top hits aggregation and its parameters, refer to the [official Elasticsearch documentation on top hits aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-top-hits-aggregation).
