package objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WildcardQueryS_MarshalJSON(t *testing.T) {
	q := WildcardQuery("field_name", "some match query",
		WithBoost(1.0),
		WithRewriteParameter(RewriteParameterConstantScore),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"wildcard":{"field_name":{"value":"some match query","boost":1,"rewrite":"constant_score"}}}`
	require.Equal(t, expected, string(body))
}
