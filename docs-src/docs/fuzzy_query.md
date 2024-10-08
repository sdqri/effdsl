# Fuzzy Query

Returns documents that contain terms similar to the search term, as measured by a Levenshtein edit distance.

An edit distance is the number of one-character changes needed to turn one term into another. These changes can include:

- Changing a character (e.g., `box` → `fox`)
- Removing a character (e.g., `black` → `lack`)
- Inserting a character (e.g., `sic` → `sick`)
- Transposing two adjacent characters (e.g., `act` → `cat`)

To find similar terms, the fuzzy query creates a set of all possible variations, or expansions, of the search term within a specified edit distance. The query then returns exact matches for each expansion.

## Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	eq "github.com/sdqri/effdsl/queries/existsquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        fq.FuzzyQuery(
            "user.id",
            "ki",
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
    _(Required, positional)_ The term you wish to find in the provided field. This is a required parameter.

*   **WithFuzziness (string)**  
    _(Optional, Functional option)_ The degree of fuzziness allowed for the search term (e.g., "AUTO", "1", "2", etc.). Defaults to no fuzziness.

*   **WithMaxExpansions (int)**  
    _(Optional, Functional option)_ Maximum number of terms to match. Defaults to 50.

*   **WithPrefixLength (int)**  
    _(Optional, Functional option)_ Number of initial characters that must match exactly. Defaults to 0.

*   **WithTranspositions (bool)**  
    _(Optional, Functional option)_ If true, allows transpositions of two adjacent characters. Defaults to true.

*   **WithRewrite (Rewrite)**  
    _(Optional, Functional option)_ Method used to rewrite the query. Valid values are:
    *   `constant_score`: Query is rewritten to a constant score query.
    *   `scoring_boolean`: Query is rewritten to a scoring boolean query.
    *   `constant_score_boolean`: Query is rewritten to a constant score boolean query.
    *   `top_terms_N`: Query is rewritten to match the top N scoring terms.
    *   `top_terms_boost_N`: Query is rewritten to match the top N scoring terms with boosting.
    *   `top_terms_blended_freqs_N`: Query is rewritten to match the top N scoring terms with blended frequencies.

### Additional Information

For more details on the fuzzy query and its parameters, refer to the [official Elasticsearch documentation on fuzzy queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html).

