# Boolean Query

A query that matches documents based on boolean combinations of other queries. The bool query maps to Lucene BooleanQuery. It is constructed using one or more boolean clauses, each with a specific occurrence type. The occurrence types are:

### Occur Types

- **must**  
  The clause (query) must appear in matching documents and will contribute to the score.

- **filter**  
  The clause (query) must appear in matching documents. Unlike must, the score of the query will be ignored. Filter clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching.

- **should**  
  The clause (query) should appear in the matching document.

- **must_not**  
  The clause (query) must not appear in the matching documents. Clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching. Because scoring is ignored, a score of 0 for all documents is returned.

The bool query adopts a more-matches-is-better approach, so the score from each matching must or should clause will be added together to provide the final _score for each document.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    mq "github.com/sdqri/effdsl/queries/matchquery"
    bq "github.com/sdqri/effdsl/queries/boolquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        bq.BoolQuery(
            bq.Must(mq.MatchQuery(effdsl.M{"user.name": "john_doe"})),
            bq.Must(mq.MatchQuery("post.status": "published")),
            bq.Filter(mq.MatchQuery("category": "technology")),
            bq.Filter(mq.MatchQuery("tags": "go")),
            bq.Should(mq.MatchQuery("title": "elasticsearch")),
            bq.Should(mq.MatchQuery("content": "search optimization")),
            bq.MustNot(mq.MatchQuery("user.role": "banned")),
            bq.MustNot(mq.MatchQuery("status": "draft")),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

*   **Must(query)**  
    _(Optional, Functional option)_ The clause must appear in matching documents and will contribute to the score.
    
*   **Filter(query)**  
    _(Optional, Functional option)_ The clause must appear in matching documents. Unlike must, the score of the query will be ignored. Filter clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching.

*   **Should(query)**  
    _(Optional, Functional option)_ The clause should appear in the matching document.

*   **MustNot(query)**  
    _(Optional, Functional option)_ The clause must not appear in the matching documents. Clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching. Because scoring is ignored, a score of 0 for all documents is returned.

*   **WithMinimumShouldMatch(string)**  
    _(Optional, Functional option)_ Minimum number of clauses that must match for a document to be returned.

### Additional Information
For more details on the boolean query and its parameters, refer to the [official Elasticsearch documentation on bool queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html).
