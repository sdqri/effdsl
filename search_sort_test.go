package effdsl_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
	rq "github.com/sdqri/effdsl/v2/queries/rangequery"
	tq "github.com/sdqri/effdsl/v2/queries/termquery"
)

func TestSortClauseWithDefaultOrder(t *testing.T) {
	expectedBody := `"fake_field"`
	actualSortClauseResult := effdsl.SortClause("fake_field", effdsl.SORT_DEFAULT)
	sortClause, err := actualSortClauseResult.Ok, actualSortClauseResult.Err
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestSortClauseWithOrder(t *testing.T) {
	expectedBody := `{"fake_field":{"order":"asc"}}`
	actualSortClauseResult := effdsl.SortClause("fake_field", effdsl.SORT_ASC)
	sortClause, err := actualSortClauseResult.Ok, actualSortClauseResult.Err
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(sortClause)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestSortClauseWithParams(t *testing.T) {
	testsCases := []struct {
		name     string
		field    string
		order    effdsl.SortOrder
		opts     []effdsl.SortClauseParameter
		expected string
	}{
		// Simple cases
		{
			name:     "no parameters",
			field:    "quantity",
			order:    effdsl.SORT_DEFAULT,
			opts:     nil,
			expected: `"quantity"`,
		},
		{
			name:     "only order asc",
			order:    effdsl.SORT_ASC,
			field:    "price",
			expected: `{"price":{"order":"asc"}}`,
		},
		{
			name:     "default order with missing first",
			field:    "category",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing(effdsl.SORT_MISSING_FIRST)},
			expected: `{"category":{"missing":"_first"}}`,
		},
		{
			name:     "order desc with missing last",
			order:    effdsl.SORT_DESC,
			field:    "date",
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing(effdsl.SORT_MISSING_LAST)},
			expected: `{"date":{"order":"desc","missing":"_last"}}`,
		},
		{
			name:     "custom missing value",
			field:    "rating",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing("0")},
			expected: `{"rating":{"missing":"0"}}`,
		},

		// Sorting with formats
		{
			name:     "only order asc with format",
			order:    effdsl.SORT_ASC,
			field:    "post_date",
			opts:     []effdsl.SortClauseParameter{effdsl.WithFormat(effdsl.FORMAT_STRICT_DATE_OPTIONAL_TIME_NANOS)},
			expected: `{"post_date":{"order":"asc","format":"strict_date_optional_time_nanos"}}`,
		},
		{
			name:     "only format",
			field:    "post_date",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithFormat(effdsl.FORMAT_STRICT_DATE_OPTIONAL_TIME_NANOS)},
			expected: `{"post_date":{"format":"strict_date_optional_time_nanos"}}`,
		},

		// Sorting with modes
		{
			name:     "order asc with mode avg",
			order:    effdsl.SORT_ASC,
			field:    "price",
			opts:     []effdsl.SortClauseParameter{effdsl.WithSortMode(effdsl.SORT_MODE_AVG)},
			expected: `{"price":{"order":"asc","mode":"avg"}}`,
		},
		{
			name:     "missing _last",
			field:    "price",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithMissing(effdsl.SORT_MISSING_LAST)},
			expected: `{"price":{"missing":"_last"}}`,
		},
		{
			name:     "unmapped_type long",
			field:    "price",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithUnmappedType("long")},
			expected: `{"price":{"unmapped_type":"long"}}`,
		},

		// Sorting with numeric types
		{
			name:     "numeric_type double",
			field:    "field",
			order:    effdsl.SORT_DEFAULT,
			opts:     []effdsl.SortClauseParameter{effdsl.WithNumericType(effdsl.SORT_NUMERIC_TYPE_DOUBLE)},
			expected: `{"field":{"numeric_type":"double"}}`,
		},
		{
			name:  "numeric_type date_nanos",
			field: "field",
			order: effdsl.SORT_DEFAULT,
			opts:  []effdsl.SortClauseParameter{effdsl.WithNumericType(effdsl.SORT_NUMERIC_TYPE_DATE_NANOS)}, expected: `{"field":{"numeric_type":"date_nanos"}}`,
		},

		// Sorting with nested filters
		{
			name:  "nested sort with mode avg and order asc",
			field: "offer.price",
			order: effdsl.SORT_ASC,
			opts: []effdsl.SortClauseParameter{
				effdsl.WithSortMode(effdsl.SORT_MODE_AVG),
				effdsl.WithNested(effdsl.NewNested(
					"offer",
					tq.TermQuery("offer.color", "blue"),
					nil,
					nil,
				)),
			},
			expected: `{"offer.price":{"order":"asc","mode":"avg","nested":{"path":"offer","filter":{"term":{"offer.color":{"value":"blue"}}}}}}`,
		},
		{
			name:  "nested sort with min mode, order asc, and filter",
			field: "parent.child.age",
			order: effdsl.SORT_ASC,
			opts: []effdsl.SortClauseParameter{
				effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
				effdsl.WithNested(effdsl.NewNested(
					"parent",
					rq.RangeQuery("parent.age", rq.WithGTE(21)),
					nil,
					effdsl.NewNested(
						"parent.child",
						mq.MatchQuery("parent.child.name", "matt"),
						nil,
						nil,
					),
				)),
			},
			expected: `{"parent.child.age":{"order":"asc","mode":"min","nested":{"path":"parent","filter":{"range":{"parent.age":{"gte":21}}},"nested":{"path":"parent.child","filter":{"match":{"parent.child.name":{"query":"matt"}}}}}}}`,
		},

		//TODO: Script-based sorting and Geo-distance sorting isn't implemented
		// Script-based sorting
		// {
		// 	name:  "_script sort with order asc",
		// 	field: "_script",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode("number"),
		// 		effdsl.WithFormat(utils.StrPtr("painless")),
		// 	},
		// 	expected: `{"_script":{"type":"number","script":{"lang":"painless","source":"doc['field_name'].value * params.factor","params":{"factor":1.1}},"order":"asc"}}`,
		// },

		// Geo-distance sorting
		// {
		// 	name:  "geo_distance with arc distance type",
		// 	field: "_geo_distance",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
		// 		effdsl.WithUnmappedType("arc"),
		// 	},
		// 	expected: `{"_geo_distance":{"pin.location":[-70,40],"order":"asc","unit":"km","mode":"min","distance_type":"arc","ignore_unmapped":true}}`,
		// },
		// {
		// 	name:  "geo_distance with point coordinates",
		// 	field: "_geo_distance",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
		// 	},
		// 	expected: `{"_geo_distance":{"pin.location":{"lat":40,"lon":-70},"order":"asc","unit":"km"}}`,
		// },
		// {
		// 	name:  "geo_distance with point string",
		// 	field: "_geo_distance",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
		// 	},
		// 	expected: `{"_geo_distance":{"pin.location":"POINT (-70 40)","order":"asc","unit":"km"}}`,
		// },
		// {
		// 	name:  "geo_distance with encoded location",
		// 	field: "_geo_distance",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
		// 	},
		// 	expected: `{"_geo_distance":{"pin.location":"drm3btev3e86","order":"asc","unit":"km"}}`,
		// },
		// {
		// 	name:  "geo_distance with multiple points",
		// 	field: "_geo_distance",
		// 	order: effdsl.SORT_ASC,
		// 	opts: []effdsl.SortClauseParameter{
		// 		effdsl.WithSortMode(effdsl.SORT_MODE_MIN),
		// 	},
		// 	expected: `{"_geo_distance":{"pin.location":[[-70,40],[-71,42]],"order":"asc","unit":"km"}}`,
		// },
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			sortClause := effdsl.SortClause(tt.field, tt.order, tt.opts...)
			assert.NoError(t, sortClause.Err)
			jsonBody, err := json.Marshal(sortClause.Ok)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(jsonBody))
		})
	}
}

func TestSortWithTrackScores(t *testing.T) {
	expectedBody := `{"sort":[{"post_date":{"order":"desc"}},{"name":{"order":"desc"}},{"age":{"order":"asc"}}],"track_scores":true}`

	expectedBody = strings.ReplaceAll(expectedBody, " ", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\t", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\n", "")

	body, err := effdsl.Define(
		effdsl.WithSort(
			effdsl.SortClause("post_date", effdsl.SORT_DESC),
			effdsl.SortClause("name", effdsl.SORT_DESC),
			effdsl.SortClause("age", effdsl.SORT_ASC),
		),
		effdsl.WithTrackScores(),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
