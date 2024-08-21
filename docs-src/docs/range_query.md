# Range Query

A range query returns documents that contain terms within a specified range. It supports querying for values that are greater than, less than, or between certain values.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	rq "github.com/sdqri/effdsl/queries/rangequery"
)

query, err := effdsl.Define(
    rq.RangeQuery(
        "field_name",
        rq.WithGT(10),
        rq.WithLTE(100),
        rq.WithFormat("yyyy-MM-dd"),
        rq.WithBoost(1.5),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

* **WithGT(any)**  
    Greater than. If specified, the range query will include terms greater than this value.

* **WithGTE(any)**  
    Greater than or equal to. If specified, the range query will include terms greater than or equal to this value.

* **WithLT(any)**  
    Less than. If specified, the range query will include terms less than this value.

* **WithLTE(any)**  
    Less than or equal to. If specified, the range query will include terms less than or equal to this value.

* **WithFormat(string)**  
    Date format used to convert date values in the query. This is optional.

* **WithRelation(Relation)**  
    Indicates how the range query matches values for range fields. Valid values are:
      * INTERSECTS
      * CONTAINS
      * WITHIN

* **WithTimeZone(string)**  
    Coordinated Universal Time (UTC) offset or IANA time zone used to convert date values in the query to UTC. This is optional.

* **WithBoost(float64)**  
    Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.

### Additional Information

For more details on the range query and its parameters, refer to the [official Elasticsearch documentation on range queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html).

