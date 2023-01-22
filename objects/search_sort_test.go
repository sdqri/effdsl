package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestNewSortClauseWithDefaultOrder(t *testing.T) {
	expectedBody := `"fake_field"`
	sortClauseResult := objs.D.SortClause("fake_field", objs.SORT_DEFAULT)
	err := sortClauseResult.Err
	sortClause := sortClauseResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewSortClauseWithOrder(t *testing.T) {
	expectedBody := `{"fake_field":"asc"}`
	sortClauseResult := objs.D.SortClause("fake_field", objs.SORT_ASC)
	err := sortClauseResult.Err
	sortClause := sortClauseResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
