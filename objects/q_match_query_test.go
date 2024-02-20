package objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MatchQueryS_MarshalJSON(t *testing.T) {
	q := MatchQuery("field_name", "some match query",
		WithMatchOperator(MatchOperatorAND),
		WithFuzzinessParameter(FuzzinessAUTO),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match":{"field_name":{"query":"some match query","operator":"AND","fuzziness":"AUTO"}}}`
	require.Equal(t, expected, string(body))
}
