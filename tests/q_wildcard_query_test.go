package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	effdsl "github.com/sdqri/effdsl"
)

func Test_WildcardQueryS_MarshalJSON(t *testing.T) {
	q := effdsl.WildcardQuery("field_name", "some match query",
		effdsl.WithBoost(1.0),
		effdsl.WithRewriteParameter(effdsl.RewriteParameterConstantScore),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"wildcard":{"field_name":{"value":"some match query","boost":1,"rewrite":"constant_score"}}}`
	require.Equal(t, expected, string(body))
}
