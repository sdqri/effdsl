# Completion Suggester

The completion suggester provides auto-complete or search-as-you-type functionality. It can return suggestions based on a prefix or a regular expression and supports filtering or boosting with contexts.

### Example

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    "github.com/sdqri/effdsl/v2"
    cs "github.com/sdqri/effdsl/suggesters/completionsuggester"
)

query, err := effdsl.Define(
    effdsl.WithSuggest(
        cs.CompletionSuggester(
            "song-suggest",
            "nir",
            "suggest",
            cs.WithSize(10),
            cs.WithCompletionSuggesterSkipDuplicates(true),
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
* **Prefix (string)**
  _(Required, positional)_ Prefix text used to generate suggestions. Use `CompletionSuggesterRegex` for regex input.
* **Field (string)**
  _(Required, positional)_ Completion field from which to fetch suggestions.
* **WithSize(uint64)**
  _(Optional, Functional option)_ Number of suggestions to return. Defaults to 5.
* **WithCompletionSuggesterSkipDuplicates(bool)**
  _(Optional, Functional option)_ Whether duplicate suggestions should be filtered out.
* **WithCompletionSuggesterFuzzy(opt ...FuzzyOption)**
  _(Optional, Functional option)_ Enables fuzzy prefix matching. `FuzzyOption` functions allow setting fuzziness, prefix length and other parameters.
* **WithRegexFlags(RegexFlag)**
  _(Optional, Functional option)_ Regex flags used with `CompletionSuggesterRegex`.
* **WithMaxDeterminizedStates(int64)**
  _(Optional, Functional option)_ Maximum allowed states for regex completion.
* **WithMultipleCategoryContext(name, contexts ...string)**
  _(Optional, Functional option)_ Adds category contexts used for filtering or boosting suggestions.
* **WithCategoryContext(name, context string, opts ...CategoryContextClauseOption)**
  _(Optional, Functional option)_ Adds a single category context with optional boost or prefix.
* **WithGeoContext(name string, lat, lon float64, opts ...GeoContextClauseOption)**
  _(Optional, Functional option)_ Adds a geo context for location based suggestions.

### Additional Information

For more details on the completion suggester and its parameters, refer to the [official Elasticsearch documentation on completion suggesters](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester).
