# Prefix Query

A prefix query returns documents that contain terms starting with the specified prefix in a given field.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	pq "github.com/sdqri/effdsl/queries/prefixquery"
)

query, err := effdsl.Define(
    pq.PrefixQuery(
        "field_name",
        "prefix_value",
        pq.WithRewrite(pq.ConstantScore),
        pq.WithCaseInsensitive(true),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

* **Value string**  
    The prefix you wish to match against terms in the provided field. This is a required parameter.

* **WithRewrite(Rewrite)**  
    The method used to rewrite the query. Valid values are:
      * constant_score: Query is rewritten to a constant score query.
      * scoring_boolean: Query is rewritten to a scoring boolean query.
      * constant_score_boolean: Query is rewritten to a constant score boolean query.
      * top_terms_N: Query is rewritten to match the top N scoring terms.
      * top_terms_boost_N: Query is rewritten to match the top N scoring terms with boosting.
      * top_terms_blended_freqs_N: Query is rewritten to match the top N scoring terms with blended frequencies.

* **WithCaseInsensitive(bool)**  
    Whether the query is case insensitive. Defaults to false.

### Additional Information

For more details on the prefix query and its parameters, refer to the [official Elasticsearch documentation on prefix queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html).

