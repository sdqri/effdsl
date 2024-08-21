# Simple Query String

A simple query string query parses a provided query string and searches for documents using the specified fields and options. It supports a variety of query options to refine the search.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	sqs "github.com/sdqri/effdsl/queries/simplequerystring"
)

query, err := effdsl.Define(
    sqs.SimpleQueryString(
        "\"fried eggs\" +(eggplant | potato) -frittata",
        sqs.WithFields("title^5", "body"),
        sqs.WithDefaultOperator(sqs.AND),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Query string**  
    The query string you wish to parse and use for search. This is a required parameter.
    
* **WithFields(...string)**  
    Array of fields to search. Supports wildcards (*).
    
* **WithDefaultOperator(Operator)**  
    Default boolean logic used to interpret text in the query string. Valid values are:
    *   OR: For example, a query value of "capital of Hungary" is interpreted as "capital OR of OR Hungary".
    *   AND: For example, a query value of "capital of Hungary" is interpreted as "capital AND of AND Hungary".

* **WithAnalyzeWildcard()**  
    If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.
    
* **WithAnalyzer(string)**  
    Analyzer used to convert text in the query string into tokens.
    
* **WithAutoGenerateSynonymsPhrase(bool)**  
    If true, match phrase queries are automatically created for multi-term synonyms. Defaults to true.
    
* **WithFlags(string)**  
    List of enabled operators for the simple query string syntax. Defaults to ALL (all operators). See Limit operators for valid values.
    
* **WithFuzzyMaxExpansions(int)**  
    Maximum number of terms for fuzzy matching expansion.
    
* **WithFuzzyPrefixLength(int)**  
    Number of beginning characters left unchanged for fuzzy matching.
    
* **WithFuzzyTranspositions(bool)**  
    If true, edits for fuzzy matching include transpositions of adjacent characters.
    
* **WithLenient(bool)**  
    If true, format-based errors are ignored.
    
* **WithMinimumShouldMatch(string)**  
    Minimum number of clauses that must match for a document to be returned.
    
* **WithQuoteFieldSuffix(string)**  
    Suffix appended to quoted text in the query string.
    

### Additional Information

For more details on the simple query string query and its parameters, refer to the [official Elasticsearch documentation on simple query string queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html).

