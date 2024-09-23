# Disjunction Max Query

A disjunction max query (dis_max) is used to find documents that match multiple query clauses. The query returns documents that match any of the provided queries, and the relevance score is determined based on the best match. This query is useful for combining multiple queries into one and adjusting their scores with a tie-breaker.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	dmq "github.com/sdqri/effdsl/queries/dismaxquery"
	tq "github.com/sdqri/effdsl/queries/termquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        dmq.DisMaxQuery(
            []effdsl.QueryResult{
                tq.TermQuery("title", "Quick pets"),
                tq.TermQuery("body", "Quick pets"),
            },
            dmq.WithTieBreaker(0.7),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Queries ([]query)**  
    _(Required, positional)_ An array of query objects that documents must match. This is a required parameter.

*   **WithTieBreaker (float64)**  
    _(Optional, Functional option)_ A floating-point number used to adjust the relevance scores when multiple queries match. This is an optional parameter.

### Additional Information

For more details on the disjunction max query and its parameters, refer to the [official Elasticsearch documentation on dis_max queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html).

