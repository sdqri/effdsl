# Match Phrase Query

A match phrase query returns documents that match a given phrase, considering the position of the terms. The provided text is analyzed before matching.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	mpq "github.com/sdqri/effdsl/queries/matchphrasequery"
)

query, err := effdsl.Define(
    mpq.MatchPhraseQuery(
        "field_name",
        "some phrase query",
        mpq.WithAnalyzer("my_analyzer"),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field to search. This is a required parameter.

* **Query string**  
    The text to search for in the provided field. This is a required parameter.

* **WithAnalyzer(string)**  
    Analyzer used to convert the text in the query value into tokens. Defaults to the index-time analyzer mapped for the field. If no analyzer is mapped, the index’s default analyzer is used.

* **WithSlop(int)**  
    Maximum number of positions allowed between matching tokens for phrases. Defaults to 0.

* **WithZeroTermsquery(ZeroTerms)**  
    Indicates what to do when the analyzed text contains no terms. Valid values are:
    * none (Default): No documents are returned if the analyzer removes all tokens.
    * all: Returns all documents, similar to a match_all query.

### Additional Information

For more details on the match phrase query and its parameters, refer to the [official Elasticsearch documentation on match phrase queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase.html).

