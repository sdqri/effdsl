package matchnonequery_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	mnq "github.com/sdqri/effdsl/v2/queries/matchnonequery"
)

func Test_MatchNoneQuery_MarshalJSON(t *testing.T) {
	q := mnq.MatchNoneQuery()

	body, err := q.Ok.MarshalJSON()
	require.NoError(t, err)

	const expected = `{"match_none":{}}`
	require.Equal(t, expected, string(body))
}
