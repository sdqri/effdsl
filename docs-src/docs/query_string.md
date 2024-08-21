# Query String Query

A query string query parses and executes a search query based on a query string syntax. It allows for flexible and complex query expressions.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	qs "github.com/sdqri/effdsl/queries/querystring"
)

query, err := effdsl.Define(
    qs.QueryString(
        "field1:value1 AND field2:value2",
        qs.WithDefaultField("field_name"),
        qs.WithAllowLeadingWildcard(),
        qs.WithBoost(1.5),
        qs.WithFuzziness("AUTO"),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Query string**  
    The query string to parse and use for search. This is a required parameter.
    
* **WithDefaultField(string)**  
    Default field to search if no field is provided in the query string.
    
* **WithAllowLeadingWildcard()**  
    If true, wildcard characters `*` and `?` are allowed as the first character in the query string. Defaults to true.
    
* **WithAnalyzeWildcard()**  
    If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.
    
* **WithAnalyzer(string)**  
    Analyzer used to convert the text in the query string into tokens.
    
* **WithAutoGenerateSynonymsPhrase(bool)**  
    If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.
    
* **WithBoost(float64)**  
    Floating point number used to adjust the relevance scores of the query.
    
* **WithDefaultOperator(Operator)**  
    Default boolean logic used to interpret text in the query string. Valid values are:
    
    *   `OR`: Logical OR.
    *   `AND`: Logical AND.
    
* **WithEnablePositionIncrements(bool)**  
    If true, enable position increments in queries constructed from the query string search.
    
* **WithFields(...string)**  
    Array of fields to search. Supports wildcards `*`.
    
* **WithFuzziness(string)**  
    Maximum edit distance allowed for fuzzy matching.
    
* **WithFuzzyMaxExpansions(int)**  
    Maximum number of terms for fuzzy matching expansion.
    
* **WithFuzzyPrefixLength(int)**  
    Number of beginning characters left unchanged for fuzzy matching.
    
* **WithFuzzyTranspositions(bool)**  
    If true, edits for fuzzy matching include transpositions of adjacent characters.
    
* **WithLenient(bool)**  
    If true, format-based errors are ignored.
    
* **WithMaxDeterminizedStates(int)**  
    Maximum number of automaton states required for the query.
    
* **WithMinimumShouldMatch(string)**  
    Minimum number of clauses that must match for a document to be returned.
    
* **WithQuoteAnalyzer(string)**  
    Analyzer used to convert quoted text in the query string into tokens.
    
* **WithPhraseSlop(int)**  
    Maximum number of positions allowed between matching tokens for phrases.
    
* **WithQuoteFieldSuffix(string)**  
    Suffix appended to quoted text in the query string.
    
* **WithRewrite(Rewrite)**  
    Method used to rewrite the query. Valid values are:
    
    *   `constant_score`
    *   `scoring_boolean`
    *   `constant_score_boolean`
    *   `top_terms_N`
    *   `top_terms_boost_N`
    *   `top_terms_blended_freqs_N`
    
* **WithTimeZone(string)**  
    UTC offset or IANA time zone used to convert date values in the query string to UTC.
    

### Additional Information

For more details on the query string query and its parameters, refer to the [official Elasticsearch documentation on query string queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html).

