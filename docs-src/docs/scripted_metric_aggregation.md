# Scripted Metric Aggregation

The scripted metric aggregation executes scripts to produce a metric output.

### Example

```go
import (
	"strings"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	scriptedmetric "github.com/sdqri/effdsl/v2/aggregations/metrics/scriptedmetric"
	"github.com/sdqri/effdsl/v2/aggregations"
)

query, err := effdsl.Define(
	effdsl.WithAggregation(
		scriptedmetric.ScriptedMetric(
			"profit",
			scriptedmetric.WithInitScript(aggregations.InlineScript("state.transactions = []")),
			scriptedmetric.WithMapScript(aggregations.InlineScript("state.transactions.add(doc.type.value == 'sale' ? doc.amount.value : -1 * doc.amount.value)")),
			scriptedmetric.WithCombineScript(aggregations.InlineScript("double profit = 0; for (t in state.transactions) { profit += t } return profit")),
			scriptedmetric.WithReduceScript(aggregations.InlineScript("double profit = 0; for (a in states) { profit += a } return profit")),
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

*   **WithInitScript (Script)**  
    _(Optional, functional option)_ Script executed before collection.

*   **WithMapScript (Script)**  
    _(Optional, functional option)_ Script executed per document.

*   **WithCombineScript (Script)**  
    _(Optional, functional option)_ Script executed per shard.

*   **WithReduceScript (Script)**  
    _(Optional, functional option)_ Script executed on the coordinating node.

*   **WithParams (map[string]any)**  
    _(Optional, functional option)_ Parameters passed into scripts.

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

For more details on the scripted metric aggregation and its parameters, refer to the [official Elasticsearch documentation on scripted metric aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-metrics-scripted-metric-aggregation).
