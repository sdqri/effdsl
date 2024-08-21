# Range Query

A range query returns documents that contain terms within a specified range. It supports querying for values that are greater than, less than, or between certain values.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"

	"github.com/sdqri/effdsl/v2"
	mppq "github.com/sdqri/effdsl/v2/queries/matchphraseprefixquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        mppq.MatchPhrasePrefixQuery(
            "field_name",
            "some phrase prefix query",
            mppq.WithAnalyzer("my_analyzer"),
            mppq.WithSlop(2),
            mppq.WithMaxExpansions(10),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```
### Parameters

*   **Field (string)**  
    _(Required, positional)_ The field to search. This is a required parameter.

*   **Query (string)**  
    _(Required, positional)_ The text to search for in the provided field. This is a required parameter.

*   **WithAnalyzer (string)**  
    _(Optional, Functional option)_ Analyzer used to convert the text in the query value into tokens. Defaults to the index-time analyzer mapped for the field. If no analyzer is mapped, the index’s default analyzer is used.

*   **WithSlop (int)**  
    _(Optional, Functional option)_ Maximum number of positions allowed between matching tokens for phrases. Defaults to 0.

*   **WithMaxExpansions (int)**  
    _(Optional, Functional option)_ Maximum number of terms to which the last provided term will expand. Defaults to not expanding terms.

*   **WithZeroTermsQuery (ZeroTerms)**  
    _(Optional, Functional option)_ Indicates what to do when the analyzed text contains no terms. Valid values are:
    
    *   `none` (Default): No documents are returned if the analyzer removes all tokens.
    *   `all`: Returns all documents, similar to a match_all query.

### Additional Information

For more details on the match phrase prefix query and its parameters, refer to the [official Elasticsearch documentation on match phrase prefix queries](https://elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html).

	"github.com/sdqri/effdsl/v2"
	rq "github.com/sdqri/effdsl/v2/queries/rangequery"
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

