# Match Phrase Prefix Query

## Overview

The `MatchPhrasePrefixQuery` is used to create a match phrase prefix query for a specific field and query text. This query type is helpful when you need to search for phrases that start with the specified query text and allow for prefix matching. The query supports various options, including the analyzer, slop, max expansions, and behavior for zero terms.

## Example

```code
import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	mppq "github.com/sdqri/effdsl/queries/matchphraseprefix"
)

func TestNewMatchPhrasePrefixQueryWithNoOptions(t *testing.T) {
	matchQueryResult := mppq.MatchPhrasePrefixQuery("message", "quick brown f")
	err := matchQueryResult.Err
	matchQuery := matchQueryResult.Ok
	assert.Nil(t, err)
	expectedBody := `{"match_phrase_prefix":{"message":{"query":"quick brown f"}}}`
	jsonBody, err := json.Marshal(matchQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
```

## Parameters

*   `Field` (Required, string): The field to search.
*   `Query` (Required, string): The query text to search for.
*   `Analyzer` (Optional, string): The analyzer used to convert text in the query string into tokens.
*   `Slop` (Optional, integer): The maximum number of positions allowed between matching tokens for phrases.
*   `MaxExpansions` (Optional, integer): The maximum number of terms to which the last provided term will expand.
*   `ZeroTermsQuery` (Optional, string): Specifies what to do when the analyzed text contains no terms. Valid values are "none" (default) or "all".

## Additional Information

This query is used when you want to search for a prefix of a phrase in a specific field. It is particularly useful for autocomplete-like search experiences, where you want to suggest phrases based on partial input.

For more details, refer to the official [Elasticsearch documentation](https://elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html).
