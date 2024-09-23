# Simple Query String

A simple query string query parses a provided query string and searches for documents using the specified fields and options. It supports a variety of query options to refine the search.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	sqs "github.com/sdqri/effdsl/queries/simplequerystring"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        sqs.SimpleQueryString(
            `"fried eggs" +(eggplant | potato) -frittata`,
            sqs.WithFields("title^5", "body"),
            sqs.WithDefaultOperator(sqs.AND),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Query (string)**  
    _(Required, positional)_ The query string you wish to parse and use for search. This is a required parameter.

*   **WithFields (...string)**  
    _(Optional, Functional option)_ Array of fields to search. Supports wildcards `*`.

*   **WithDefaultOperator (Operator)**  
    _(Optional, Functional option)_ Default boolean logic used to interpret text in the query string. Valid values are:
    *   `OR`: For example, a query value of "capital of Hungary" is interpreted as "capital OR of OR Hungary".
    *   `AND`: For example, a query value of "capital of Hungary" is interpreted as "capital AND of AND Hungary".

*   **WithAnalyzeWildcard ()**  
    _(Optional, Functional option)_ If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.

*   **WithAnalyzer (string)**  
    _(Optional, Functional option)_ Analyzer used to convert text in the query string into tokens.

*   **WithAutoGenerateSynonymsPhrase (bool)**  
    _(Optional, Functional option)_ If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.

*   **WithFlags (string)**  
    _(Optional, Functional option)_ List of enabled operators for the simple query string syntax. Defaults to ALL (all operators). See Limit operators for valid values.

*   **WithFuzzyMaxExpansions (int)**  
    _(Optional, Functional option)_ Maximum number of terms for fuzzy matching expansion.

*   **WithFuzzyPrefixLength (int)**  
    _(Optional, Functional option)_ Number of beginning characters left unchanged for fuzzy matching.

*   **WithFuzzyTranspositions (bool)**  
    _(Optional, Functional option)_ If true, edits for fuzzy matching include transpositions of adjacent characters.

*   **WithLenient (bool)**  
    _(Optional, Functional option)_ If true, format-based errors are ignored.

*   **WithMinimumShouldMatch (string)**  
    _(Optional, Functional option)_ Minimum number of clauses that must match for a document to be returned.

*   **WithQuoteFieldSuffix (string)**  
    _(Optional, Functional option)_ Suffix appended to quoted text in the query string.

### Additional Information

For more details on the simple query string query and its parameters, refer to the [official Elasticsearch documentation on simple query string queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html).

