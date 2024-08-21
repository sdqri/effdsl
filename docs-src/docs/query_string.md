# Query String Query

A query string query parses and executes a search query based on a query string syntax. It allows for flexible and complex query expressions.

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
	qs "github.com/sdqri/effdsl/v2/queries/querystring"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        qs.QueryString(
            "alice",
            qs.WithFields("first_name", "last_name")
            qs.WithBoost(1.5),
            qs.WithFuzziness("AUTO"),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Query (string)**  
    _(Required, positional)_ The query string to parse and use for search. This is a required parameter.

*   **WithDefaultField (string)**  
    _(Optional, Functional option)_ Default field to search if no field is provided in the query string.

*   **WithAllowLeadingWildcard ()**  
    _(Optional, Functional option)_ If true, wildcard characters `*` and `?` are allowed as the first character in the query string. Defaults to true.

*   **WithAnalyzeWildcard ()**  
    _(Optional, Functional option)_ If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.

*   **WithAnalyzer (string)**  
    _(Optional, Functional option)_ Analyzer used to convert the text in the query string into tokens.

*   **WithAutoGenerateSynonymsPhrase (bool)**  
    _(Optional, Functional option)_ If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to adjust the relevance scores of the query.

*   **WithDefaultOperator (Operator)**  
    _(Optional, Functional option)_ Default boolean logic used to interpret text in the query string. Valid values are:
    *   `OR`: Logical OR.
    *   `AND`: Logical AND.

*   **WithEnablePositionIncrements (bool)**  
    _(Optional, Functional option)_ If true, enable position increments in queries constructed from the query string search.

*   **WithFields (...string)**  
    _(Optional, Functional option)_ Array of fields to search. Supports wildcards `*`.

*   **WithFuzziness (string)**  
    _(Optional, Functional option)_ Maximum edit distance allowed for fuzzy matching.

*   **WithFuzzyMaxExpansions (int)**  
    _(Optional, Functional option)_ Maximum number of terms for fuzzy matching expansion.

*   **WithFuzzyPrefixLength (int)**  
    _(Optional, Functional option)_ Number of beginning characters left unchanged for fuzzy matching.

*   **WithFuzzyTranspositions (bool)**  
    _(Optional, Functional option)_ If true, edits for fuzzy matching include transpositions of adjacent characters.

*   **WithLenient (bool)**  
    _(Optional, Functional option)_ If true, format-based errors are ignored.

*   **WithMaxDeterminizedStates (int)**  
    _(Optional, Functional option)_ Maximum number of automaton states required for the query.

*   **WithMinimumShouldMatch (string)**  
    _(Optional, Functional option)_ Minimum number of clauses that must match for a document to be returned.

*   **WithQuoteAnalyzer (string)**  
    _(Optional, Functional option)_ Analyzer used to convert quoted text in the query string into tokens.

*   **WithPhraseSlop (int)**  
    _(Optional, Functional option)_ Maximum number of positions allowed between matching tokens for phrases.

*   **WithQuoteFieldSuffix (string)**  
    _(Optional, Functional option)_ Suffix appended to quoted text in the query string.

*   **WithRewrite (Rewrite)**  
    _(Optional, Functional option)_ Method used to rewrite the query. Valid values are:
    *   `constant_score`
    *   `scoring_boolean`
    *   `constant_score_boolean`
    *   `top_terms_N`
    *   `top_terms_boost_N`
    *   `top_terms_blended_freqs_N`

*   **WithTimeZone (string)**  
    _(Optional, Functional option)_ UTC offset or IANA time zone used to convert date values in the query string to UTC.
   

### Additional Information

For more details on the query string query and its parameters, refer to the [official Elasticsearch documentation on query string queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html).

