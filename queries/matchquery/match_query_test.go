package matchquery_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
)

func Test_MatchQueryS_MarshalJSON(t *testing.T) {
	q := mq.MatchQuery("field_name", "some match query",
		mq.WithOperator(mq.AND),
		mq.WithFuzzinessParameter(mq.FuzzinessAUTO),
	)

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match":{"field_name":{"query":"some match query","fuzziness":"AUTO","operator":"AND"}}}`
	require.Equal(t, expected, string(body))
}
