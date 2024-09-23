# Terms Set Query

A term set query returns documents that contain at least one of the specified terms in a provided field. To return a document, at least one of the terms must exactly match the field value, including whitespace and capitalization.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	tsq "github.com/sdqri/effdsl/queries/termssetquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        tsq.TermsSetQuery(
            "programming_languages",
            []string{"c++", "java", "php"},
            tsq.WithMinimumShouldMatchField("required_matches"),
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

*   **Terms ([]string)**  
    _(Required, positional)_ An array of terms you wish to find in the provided field. To return a document, at least one of the terms must exactly match the field value, including whitespace and capitalization. This is a required parameter.

*   **WithMinimumShouldMatchField (string)**  
    _(Optional, Functional option)_ The field that holds the minimum number of terms that should match. Only used when `minimum_should_match_script` is not set.

*   **WithMinimumShouldMatchScript (string)**  
    _(Optional, Functional option)_ Script that returns the minimum number of terms that should match.

### Additional Information

For more details on the term set query and its parameters, refer to the [official Elasticsearch documentation on term set queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-set-query.html).

