# Match All Query

A match all query returns all documents in the index, assigning each a `_score` of `1.0`. This query is useful when retrieving all documents or when combined with other queries.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	maq "github.com/sdqri/effdsl/queries/matchallquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        maq.MatchAllQuery(
            maq.WithBoost(1.2),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)

```

### Parameters

*   **WithBoost (float64)**  
    _(Optional, Functional option)_ Floating-point number used to decrease or increase the relevance scores of the query.

---

# Match None Query

A match none query returns no documents. This query is useful when dynamically constructing queries and needing an explicit way to return no results.

### Example

```go
import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/sdqri/effdsl/v2"
	mnq "github.com/sdqri/effdsl/queries/matchnonequery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        mnq.MatchNoneQuery(),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Additional Information

For more details on the match all query and its parameters, refer to the [official Elasticsearch documentation on match all queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html).


