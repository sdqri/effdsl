package matchallquery_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	maq "github.com/sdqri/effdsl/v2/queries/matchallquery"
)

func Test_MatchAllQueryWithBoost_MarshalJSON(t *testing.T) {
	q := maq.MatchAllQuery(maq.WithBoost(3.4))

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match_all":{"boost":3.4}}`
	require.Equal(t, expected, string(body))
}

func Test_MatchAllQueryWithoutOptions_MarshalJSON(t *testing.T) {
	q := maq.MatchAllQuery()

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match_all":{}}`
	require.Equal(t, expected, string(body))
}
