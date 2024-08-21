# Terms Query

A terms query returns documents that contain one or more exact terms in a provided field.

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
	tsq "github.com/sdqri/effdsl/v2/queries/termsquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        tsq.TermsQuery(
            "user.id",
            []string{"kimchy", "elkbee"},
            tsq.WithBoost(1.0),
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

*   **Values ([]string)**  
    _(Required, positional)_ The array of terms you wish to find in the provided field. This is a required parameter.

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating point number used to decrease or increase the relevance scores of a query. Defaults to 1.0.
   

### Additional Information

For more details on the terms query and its parameters, refer to the [official Elasticsearch documentation on terms queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html).

