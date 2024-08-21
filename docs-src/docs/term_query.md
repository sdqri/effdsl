# Term Query

A term query returns documents that contain an exact term in a provided field. The term must exactly match the field value, including whitespace and capitalization.

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
	tq "github.com/sdqri/effdsl/v2/queries/termquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        tq.TermQuery(
            "user.id",
            "kimchy",
            tq.WithBoost(1.5),
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
    _(Required, positional)_ The term you wish to find in the provided field. This is a required parameter. The term must exactly match the field value, including whitespace and capitalization.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to decrease or increase the relevance scores of the query. Defaults to 1.0.

*   **WithCaseInsensitive (bool)**  
    _(Optional, Functional option)_ Allows ASCII case-insensitive matching of the value with the indexed field values when set to true. Defaults to false.
   

### Additional Information

For more details on the term query and its parameters, refer to the [official Elasticsearch documentation on term queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html).

