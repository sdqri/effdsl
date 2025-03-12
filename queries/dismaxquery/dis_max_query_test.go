package dismaxquery_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
	dmq "github.com/sdqri/effdsl/v2/queries/dismaxquery"
	tq "github.com/sdqri/effdsl/v2/queries/termquery"
)

func TestDisMaxQuery(t *testing.T) {
	expectedBody := `{"dis_max":{"queries":[{"term":{"title":{"value":"Quick pets"}}},{"term":{"body":{"value":"Quick pets"}}}],"tie_breaker":0.7}}`

	disMaxQueryResult := dmq.DisMaxQuery(
		[]effdsl.QueryResult{
			tq.TermQuery("title", "Quick pets"),
			tq.TermQuery("body", "Quick pets"),
		},
		dmq.WithTieBreaker(0.7),
	)
	err := disMaxQueryResult.Err
	disMaxQuery := disMaxQueryResult.Ok
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jsonBody, err := json.Marshal(disMaxQuery)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}
