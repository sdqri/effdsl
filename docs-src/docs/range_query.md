# Range Query

A range query returns documents that contain terms within a specified range. It supports querying for values that are greater than, less than, or between certain values.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	rq "github.com/sdqri/effdsl/queries/rangequery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        rq.RangeQuery(
            "age",
            rq.WithGT(10),
            rq.WithLTE(20),
            rq.WithBoost(2.0),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Field (string)**  
    _(Required, positional)_ The field you wish to search. This is a required parameter.

*   **WithGT (any)**  
    _(Optional, Functional option)_ Greater than. If specified, the range query will include terms greater than this value.

*   **WithGTE (any)**  
    _(Optional, Functional option)_ Greater than or equal to. If specified, the range query will include terms greater than or equal to this value.

*   **WithLT (any)**  
    _(Optional, Functional option)_ Less than. If specified, the range query will include terms less than this value.

*   **WithLTE (any)**  
    _(Optional, Functional option)_ Less than or equal to. If specified, the range query will include terms less than or equal to this value.

*   **WithFormat (string)**  
    _(Optional, Functional option)_ Date format used to convert date values in the query.

*   **WithRelation (Relation)**  
    _(Optional, Functional option)_ Indicates how the range query matches values for range fields. Valid values are:
    *   `INTERSECTS`
    *   `CONTAINS`
    *   `WITHIN`

*   **WithTimeZone (string)**  
    _(Optional, Functional option)_ Coordinated Universal Time (UTC) offset or IANA time zone used to convert date values in the query to UTC.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.

### Additional Information

For more details on the range query and its parameters, refer to the [official Elasticsearch documentation on range queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html).

