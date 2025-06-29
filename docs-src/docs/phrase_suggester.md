# Phrase Suggester

The phrase suggester corrects entire phrases based on n-gram matching. It is useful for providing "did you mean" functionality.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    mq "github.com/sdqri/effdsl/queries/matchquery"
    ps "github.com/sdqri/effdsl/suggesters/phrasesuggester"
)

query, err := effdsl.Define(
    effdsl.WithSuggest(
        ps.PhraseSuggester(
            "simple-phrase",
            "noble prize",
            "title.trigram",
            ps.WithLaplaceSmoothing(0.7),
            ps.WithDirectGenerator("title.trigram", ps.WithSuggestMode(ps.Always)),
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
  _(Required, positional)_ Text to generate phrase suggestions for.
* **Field (string)**
  _(Required, positional)_ Field that contains the n‑grams used for suggestions.
* **WithGramSize(uint64)**
  _(Optional, Functional option)_ Maximum size of the n‑grams in the field.
* **WithRealWordErrorLikelihood(float64)**
  _(Optional, Functional option)_ Likelihood of a term being misspelled even if it exists in the dictionary.
* **WithConfidence(float64)**
  _(Optional, Functional option)_ Threshold factor applied to input phrase scores.
* **WithMaxErrors(float64)**
  _(Optional, Functional option)_ Maximum percentage or number of terms considered misspellings.
* **WithSeparator(string)**
  _(Optional, Functional option)_ Separator used to join tokens in the bigram field.
* **WithSize(uint64)**
  _(Optional, Functional option)_ Number of candidate phrases to return.
* **WithAnalyzer(string)**
  _(Optional, Functional option)_ Analyzer used to analyze the suggest text.
* **WithShardSize(uint64)**
  _(Optional, Functional option)_ Maximum number of suggestions retrieved from each shard.
* **WithHighlight(preTag, postTag string)**
  _(Optional, Functional option)_ Adds highlighting to changed tokens.
* **WithCollate(queryResult, opts ...WithCollateOption)**
  _(Optional, Functional option)_ Checks suggestions against a query. `WithParams` and `WithPrune` are available options.
* **WithStupidBackoffSmoothing(discount float64)**, **WithLaplaceSmoothing(alpha float64)**, **WithLinearInterpolationSmoothing(tri, bi, uni float64)**
  _(Optional, Functional option)_ Smoothing models used to balance frequencies.
* **WithDirectGenerator(field string, opts ...WithDirectGeneratorOption)**
  _(Optional, Functional option)_ Adds a direct generator. Options include `WithSuggestMode`, `WithDirectGeneratorSize`, `WithMaxEdits`, `WithPrefixLength`, `WithMinWordLength`, `WithMaxInspections`, `WithMinDocFreq`, `WithMaxTermFreq`, `WithPreFilter` and `WithPostFilter`.

### Additional Information

For more details on the phrase suggester and its parameters, refer to the [official Elasticsearch documentation on phrase suggesters](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters-phrase.html).
