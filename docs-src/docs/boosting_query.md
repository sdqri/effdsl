# Boosting Query

A boosting query matches documents based on a positive query while reducing the relevance score of documents that also match a negative query. This type of query is useful for situations where you want to boost the relevance of documents that match a primary condition but penalize documents that match a secondary, less desired condition. The boosting query is constructed using a positive query, a negative query, and a negative boost factor.

### Example

```go
import (
	"github.com/sdqri/effdsl"
	bq "github.com/sdqri/effdsl/queries/boostingquery"
	tq "github.com/sdqri/effdsl/queries/termquery"
)

query, err := effdsl.Define(
    bq.BoostingQuery(
	    tq.TermQuery("text", "apple"),
        tq.TermQuery("text", "pie tart fruit crumble tree"),
        0.5, 
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

- **Positive (Query)**  
  The query that documents must match to be considered for inclusion in the results.

- **Negative (Query)**  
  The query object used to reduce the relevance score of documents matching this query.
  
- **NegativeBoost (float64)**  
  A floating-point number between 0 and 1.0 used to decrease the relevance scores of documents matching the negative query.

### Additional Information

For more details on the boosting query and its parameters, refer to the [official Elasticsearch documentation on boosting queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html).

