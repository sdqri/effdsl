# Term Suggester

The term suggester suggests corrections for individual terms based on edit distance. Each term in the provided text is analyzed, and possible corrections are returned.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    ts "github.com/sdqri/effdsl/suggesters/termsuggester"
)

query, err := effdsl.Define(
    effdsl.WithSuggest(
        ts.TermSuggester(
            "my-suggestion",
            "tring out Elasticsearch",
            "message",
            ts.WithAnalyzer("test"),
            ts.WithSize(1),
            ts.WithSort(ts.ByScore),
            ts.WithMode(ts.Always),
        ),
    ),
)

res, err := es.Search(
    es.Search.WithBody(strings.NewReader(query)),
)
```

### Parameters

* **SuggestName (string)**
  _(Required, positional)_ Name used to identify the suggestion in the response.
* **Text (string)**
  _(Required, positional)_ Text to generate suggestions for.
* **Field (string)**
  _(Required, positional)_ Field to fetch candidate suggestions from.
* **WithAnalyzer(string)**
  _(Optional, Functional option)_ Analyzer used to analyze the suggest text.
* **WithSize(uint64)**
  _(Optional, Functional option)_ Maximum number of suggestions to return.
* **WithSort(TermSuggestSort)**
  _(Optional, Functional option)_ Defines how suggestions are sorted. Possible values: `ByScore`, `ByFrequency`.
* **WithMode(TermSuggesterMode)**
  _(Optional, Functional option)_ Controls which suggestions are included. Possible values: `Missing`, `Popular`, `Always`.

### Additional Information

For more details on the term suggester and its parameters, refer to the [official Elasticsearch documentation on term suggesters](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#term-suggester).
