package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	effdsl "github.com/sdqri/effdsl"
)

func Test_MatchQueryS_MarshalJSON(t *testing.T) {
	q := effdsl.MatchQuery("field_name", "some match query",
		effdsl.WithMatchOperator(effdsl.MatchOperatorAND),
		effdsl.WithFuzzinessParameter(effdsl.FuzzinessAUTO),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match":{"field_name":{"query":"some match query","operator":"AND","fuzziness":"AUTO"}}}`
	require.Equal(t, expected, string(body))
}
