package wildcardquery_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	wcq "github.com/sdqri/effdsl/v2/queries/wildcardquery"
)

func Test_WildcardQueryS_MarshalJSON(t *testing.T) {
	q := wcq.WildcardQuery("field_name", "some match query",
		wcq.WithBoost(1.0),
		wcq.WithRewrite(wcq.ConstantScoreBlended),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"wildcard":{"field_name":{"boost":1,"rewrite":"constant_score_blended","value":"some match query"}}}`
	require.Equal(t, expected, string(body))
}
