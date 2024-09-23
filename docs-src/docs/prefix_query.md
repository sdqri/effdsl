# Prefix Query

A prefix query returns documents that contain terms starting with the specified prefix in a given field.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	pq "github.com/sdqri/effdsl/queries/prefixquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        pq.PrefixQuery(
            "name",
            "al",
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

*   **Value (string)**  
    _(Required, positional)_ The prefix you wish to match against terms in the provided field. This is a required parameter.

*   **WithRewrite (Rewrite)**  
    _(Optional, Functional option)_ The method used to rewrite the query. Valid values are:
    *   `constant_score`: Query is rewritten to a constant score query.
    *   `scoring_boolean`: Query is rewritten to a scoring boolean query.
    *   `constant_score_boolean`: Query is rewritten to a constant score boolean query.
    *   `top_terms_N`: Query is rewritten to match the top N scoring terms.
    *   `top_terms_boost_N`: Query is rewritten to match the top N scoring terms with boosting.
    *   `top_terms_blended_freqs_N`: Query is rewritten to match the top N scoring terms with blended frequencies.

*   **WithCaseInsensitive (bool)**  
    _(Optional, Functional option)_ Whether the query is case insensitive. Defaults to false.

### Additional Information

For more details on the prefix query and its parameters, refer to the [official Elasticsearch documentation on prefix queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html).

