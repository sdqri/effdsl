package effdsl_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
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
	expectedBody := `{"fake_field":{"order":"asc"}}`
	sortClauseResult := effdsl.SortClause("fake_field", effdsl.SORT_ASC)
	err := sortClauseResult.Err
	sortClause := sortClauseResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNewSortClauseWithOrderAndMissing(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		order    effdsl.SortOrder
		opts     []effdsl.SortClauseParameter
		expected string
	}{
		{
			name:     "only order asc",
			order:    effdsl.SORT_ASC,
			field:    "price",
			expected: `{"price":{"order":"asc"}}`,
		},
		{
			name:     "order desc with missing last",
			order:    effdsl.SORT_DESC,
			field:    "date",
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing(effdsl.MISSING_LAST)},
			expected: `{"date":{"order":"desc","missing":"_last"}}`,
		},
		{
			name:     "custom missing value",
			field:    "rating",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing("0")},
			expected: `{"rating":{"missing":"0"}}`,
		},
		{
			name:     "default order with missing first",
			field:    "category",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing(effdsl.MISSING_FIRST)},
			expected: `{"category":{"missing":"_first"}}`,
		},
		{
			name:     "no parameters",
			field:    "quantity",
			order:    effdsl.SORT_DEFAULT,
			opts:     nil,
			expected: `"quantity"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortClause := effdsl.SortClause(tt.field, tt.order, tt.opts...)
			jsonBody, err := json.Marshal(sortClause.Ok)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(jsonBody))
		})
	}
}
