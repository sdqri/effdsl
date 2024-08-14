package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl"
)

func TestNewSortClauseWithDefaultOrder(t *testing.T) {
	expectedBody := `"fake_field"`
	sortClauseResult := effdsl.SortClause("fake_field", effdsl.SORT_DEFAULT)
	err := sortClauseResult.Err
	sortClause := sortClauseResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewSortClauseWithOrder(t *testing.T) {
	expectedBody := `{"fake_field":"asc"}`
	sortClauseResult := effdsl.SortClause("fake_field", effdsl.SORT_ASC)
	err := sortClauseResult.Err
	sortClause := sortClauseResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
