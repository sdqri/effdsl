# Nested Query

A nested query searches nested field objects as if they were separate documents. It returns the root document when a nested document matches the provided query.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    bq "github.com/sdqri/effdsl/queries/boolquery"
    mq "github.com/sdqri/effdsl/queries/matchquery"
    nq "github.com/sdqri/effdsl/queries/nestedquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        nq.WithNested(
            "path",
            bq.BoolQuery(
                bq.Should(
                    mq.MatchQuery("field1", "val1"),
                ),
            ),
            nq.WithScoreMode("avg"),
            nq.WithIgnoreUnmapped(true),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **Path (string)**
  _(Required, positional)_ Path to the nested object you wish to search.
* **Query (effdsl.Query)**
  _(Required, positional)_ Query to run on the nested objects.
* **WithScoreMode(string)**
  _(Optional, Functional option)_ Indicates how scores for matching child objects affect the parent document's score.
* **WithIgnoreUnmapped(bool)**
  _(Optional, Functional option)_ If true, ignore unmapped paths and return no documents instead of an error.

### Additional Information

For more details on the nested query and its parameters, refer to the [official Elasticsearch documentation on nested queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html).
