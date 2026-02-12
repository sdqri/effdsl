# Bucket Selector Aggregation

The bucket selector aggregation filters buckets based on a script.

### Example

```go
import (
    "strings"

    es "github.com/elastic/go-elasticsearch/v8"

    "github.com/sdqri/effdsl/v2"
    "github.com/sdqri/effdsl/v2/aggregations"
    bucketselector "github.com/sdqri/effdsl/v2/aggregations/pipeline/bucketselector"
)

query, err := effdsl.Define(
    effdsl.WithAggregation(
        bucketselector.BucketSelector(
            "only_profitable",
            map[string]string{
                "total_sales": "sales_per_month>sales",
                "total_cost":  "sales_per_month>cost",
            },
            aggregations.Script{Source: "params.total_sales > params.total_cost"},
            bucketselector.WithGapPolicy("skip"),
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

*   **BucketsPath (map[string]string)**  
    _(Required, positional)_ Map of variables to bucketed metric paths.

*   **Script (Script)**  
    _(Required, positional)_ Script that returns a boolean for bucket selection.

*   **WithGapPolicy (string)**  
    _(Optional, functional option)_ Gap policy (`skip`, `insert_zeros`).

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

For more details on the bucket selector aggregation and its parameters, refer to the [official Elasticsearch documentation on bucket selector aggregation](https://www.elastic.co/docs/reference/aggregations/search-aggregations-pipeline-bucket-selector-aggregation).
