# Regexp Query

A regexp query returns documents that contain terms matching a specified regular expression. The regular expression can include additional options for controlling the match behavior.

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
	rq "github.com/sdqri/effdsl/v2/queries/regexpquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        rq.RegexpQuery(
            "user.id",
            "k.*y",
            rq.WithFlags("ALL"),
            rq.WithCaseInsensitive(),
            rq.WithMaxDeterminizedStates(10000),
            rq.WithRQRewrite(rq.ConstantScore),
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
    _(Required, positional)_ The regular expression pattern to match against the field. This is a required parameter.

*   **WithFlags (string)**  
    _(Optional, Functional option)_ Additional matching options for the regular expression.

*   **WithCaseInsensitive (bool)**  
    _(Optional, Functional option)_ If true, the regular expression is case-insensitive.

*   **WithMaxDeterminizedStates (int)**  
    _(Optional, Functional option)_ The maximum number of automaton states required for the query. Lower values will reduce memory usage but increase query time.

*   **WithRewrite (Rewrite)**  
    _(Optional, Functional option)_ The method used to rewrite the query. Valid values are:
    *   `constant_score`: Query is rewritten to a constant score query.
    *   `scoring_boolean`: Query is rewritten to a scoring boolean query.
    *   `constant_score_boolean`: Query is rewritten to a constant score boolean query.
    *   `top_terms_N`: Query is rewritten to match the top N scoring terms.
    *   `top_terms_boost_N`: Query is rewritten to match the top N scoring terms with boosting.
    *   `top_terms_blended_freqs_N`: Query is rewritten to match the top N scoring terms with blended frequencies.

### Additional Information

For more details on the regexp query and its parameters, refer to the [official Elasticsearch documentation on regexp queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html).

