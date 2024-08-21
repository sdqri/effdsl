# Terms Set Query

A term set query returns documents that contain at least one of the specified terms in a provided field. To return a document, at least one of the terms must exactly match the field value, including whitespace and capitalization.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	tsq "github.com/sdqri/effdsl/queries/termssetquery"
)

query, err := effdsl.Define(
    tsq.TermsSetQuery(
        "field_name",
        []string{"term1", "term2"},
        tsq.WithMinimumShouldMatchField("min_should_match_field"),
        tsq.WithMinimumShouldMatchScript("min_should_match_script"),
    )
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Field string**  
    The field you wish to search. This is a required parameter.

* **Terms []string**  
    An array of terms you wish to find in the provided field. To return a document, at least one of the terms must exactly match the field value, including whitespace and capitalization. This is a required parameter.

* **WithMinimumShouldMatchField(string)**  
    The field which holds the minimum number of terms that should match. Only used when `minimum_should_match_script` is not set.

* **WithMinimumShouldMatchScript(string)**  
    Script which returns the minimum number of terms that should match.

### Additional Information

For more details on the term set query and its parameters, refer to the [official Elasticsearch documentation on term set queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-set-query.html).

