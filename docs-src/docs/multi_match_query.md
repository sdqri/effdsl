# Multi Match Query

A multi match query searches multiple fields for the provided query text. It is useful when the same text should be matched against a set of fields with different weights or options.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    mmq "github.com/sdqri/effdsl/queries/multimatchquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        mmq.MultiMatchQuery(
            "quick brown fox",
            mmq.WithFields("title", "message"),
            mmq.WithType(mmq.BestFields),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Query (string)**
    _(Required, positional)_ The text to search for. This is a required parameter.

*   **WithFields (...string)**
    _(Optional, Functional option)_ Fields to search. Each field can include a boost using the `field^boost` syntax.

*   **WithType (MultiMatchType)**
    _(Optional, Functional option)_ Type of multi match query. Valid values are:
    * `best_fields`
    * `most_fields`
    * `cross_fields`
    * `phrase`
    * `phrase_prefix`
    * `bool_prefix`

*   **WithOperator (Operator)**
    _(Optional, Functional option)_ Boolean logic used to interpret text in the query. Valid values are:
    * `or` (Default)
    * `and`

*   **WithAnalyzer (string)**
    _(Optional, Functional option)_ Analyzer used to convert the text in the query value into tokens.

*   **WithSlop (int)**
    _(Optional, Functional option)_ Maximum number of positions allowed between matching tokens for phrases.

*   **WithFuzziness (string)**
    _(Optional, Functional option)_ Fuzziness used for fuzzy matching.

*   **WithPrefixLength (int)**
    _(Optional, Functional option)_ Number of beginning characters left unchanged for fuzzy matching.

*   **WithMaxExpansions (int)**
    _(Optional, Functional option)_ Maximum number of terms the query will expand to for fuzzy matching.

*   **WithMinimumShouldMatch (string)**
    _(Optional, Functional option)_ Minimum number of clauses that must match for a document to be returned.

*   **WithTieBreaker (float64)**
    _(Optional, Functional option)_ Tie breaker used when `best_fields` type is specified.

*   **WithLenient (bool)**
    _(Optional, Functional option)_ If true, format-based errors, such as providing text for a numeric field, are ignored.

*   **WithZeroTermsQuery (ZeroTerms)**
    _(Optional, Functional option)_ Indicates what to do when the query string contains no terms. Valid values are:
    * `none` (Default)
    * `all`

*   **WithAutoGenerateSynonymsPhraseQuery (bool)**
    _(Optional, Functional option)_ If true, match phrase queries are automatically created for multi-term synonyms.

### Additional Information

For more details on the multi match query and its parameters, refer to the [official Elasticsearch documentation on multi match queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html).
 
