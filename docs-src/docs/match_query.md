# Match Query

A match query returns documents that match a provided text, number, date, or boolean value. The provided text is analyzed before matching. 

### Example

```go
import (
	"github.com/sdqri/effdsl"
	mq "github.com/sdqri/effdsl/queries/matchquery"
)

query, err := effdsl.Define(
    mq.MatchQuery(
        "field_name",
        "some match query",
        mq.WithOperator(mq.AND),
        mq.WithFuzzinessParameter(mq.FuzzinessAUTO),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

* **Query string**  
    The text, number, boolean value, or date you wish to find in the provided field. This is a required parameter.

* **WithAnalyzer(string)**  
    Analyzer used to convert the text in the query value into tokens. Defaults to the index-time analyzer mapped for the field. If no analyzer is mapped, the index’s default analyzer is used.

* **WithAutoGenerateSynonymsPhrase(bool)**
    If true, match phrase queries are automatically created for multi-term synonyms.

* **WithBoost(float64)** 
    Floating point number used to decrease or increase the relevance scores of the query.

* **WithFuzzinessParameter(int)**
    Maximum number of terms to which the query will expand.

* **WithMaxExpansions(int)**
    Maximum number of terms to which the query will expand. Defaults to 50.

* **WithPrefixLength(int)**
    Number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
  
* **WithFuzzyTranspositions(bool)**
    If true, edits for fuzzy matching include transpositions of two adjacent characters.

* **WithFuzzyRewrite(FuzzyRewrite)**
    Method used to rewrite the query. See the rewrite parameter for valid values and more information.

* **WithOperator(Operator)**
    Boolean logic used to interpret text in the query value. Valid values are:
      * `OR` (Default): For example, a query value of "capital of Hungary" is interpreted as "capital OR of OR Hungary".
      * `AND`: For example, a query value of "capital of Hungary" is interpreted as "capital AND of AND Hungary".

* **WithMinimumShouldMatch(string)** 
    Minimum number of clauses that must match for a document to be returned.

### Additional Information

For more details on the match query and its parameters, refer to the [official Elasticsearch documentation on match queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html).

