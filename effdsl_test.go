package effdsl_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl/v2"
	bq "github.com/sdqri/effdsl/v2/queries/boolquery"
	eq "github.com/sdqri/effdsl/v2/queries/existsquery"
	qs "github.com/sdqri/effdsl/v2/queries/querystring"
	rq "github.com/sdqri/effdsl/v2/queries/rangequery"
	tq "github.com/sdqri/effdsl/v2/queries/termquery"
)

func TestWithQuery(t *testing.T) {
	expectedBody := `{"query":{"query_string":{"query":"query","fields":["field1","field2"]}}}`
	f := effdsl.WithQuery(
		qs.QueryString("query", qs.WithFields("field1", "field2")),
	)
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithSort(t *testing.T) {
	expectedBody := `{"sort":["sort_field"]}`
	f := effdsl.WithSort(effdsl.SortClause("sort_field", effdsl.SORT_DEFAULT))
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithCollapse(t *testing.T) {
	expectedBody := `{"collapse":{"field":"collapsed_field"}}`
	f := effdsl.WithCollpse("collapsed_field")
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithSearchAfter(t *testing.T) {
	expectedBody := `{"search_after":["2021-05-20T05:30:04.832Z",4294967298]}`
	f := effdsl.WithSearchAfter("2021-05-20T05:30:04.832Z", 4294967298)
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithPIT(t *testing.T) {
	expectedBody := `{"pit":{"id":"test_id","keep_alive":"1m"}}`
	f := effdsl.WithPIT("test_id", "1m")
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithAggregations(t *testing.T) {
	expectedBody := `
		{
			"aggs":{
				"terms_aggregation_field":{
					"terms":{"field":"terms_aggregation_field","size":10}
				},
				"stats_aggregation_field":{
					"stats":{"field":"stats_aggregation_field"}
				}
			}
		}`
	f := effdsl.WithAggregations(
		effdsl.TermAggregation("terms_aggregation_field", 10),
		effdsl.StatsAggregation("stats_aggregation_field"),
	)
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.JSONEq(t, expectedBody, string(jsonBody))
}

func TestDefindeQ1(t *testing.T) {
	expectedBody := `
	{
		"_source":{
		   "includes":[
			  "field1",
			  "field2"
		   ],
		   "excludes":[
			  "field2",
			  "field4"
		   ]
		},
		"size":10,
		"query":{
		   "bool":{
			  "must":[
				 {
					"query_string":{
					   "query":"fake_value1",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 },
				 {
					"query_string":{
					   "query":"fake_value2",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 }
			  ],
			  "filter":[
				 {
					"range":{
					   "published_at":{
						  "gt":"now-24h"
					   }
					}
				 },
				 {
					"term":{
					   "field1.keyword":{
						  "value":"fake_value"
					   }
					}
				 },
				 {
					"exists":{
					   "field":"fake_field"
					}
				 }
			  ],
			  "should":[
				 {
					"query_string":{
					   "query":"fake_value5",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 },
				 {
					"query_string":{
					   "query":"fake_value6",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 }
			  ],
			  "must_not":[
				 {
					"query_string":{
					   "query":"fake_value3",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 },
				 {
					"query_string":{
					   "query":"fake_value4",
					   "fields":[
						  "title",
						  "description",
						  "content"
					   ]
					}
				 }
			  ]
		   }
		},
		"sort":[
		   {
			  "sort_field1":{"order": "desc"}
		   },
		   "_score"
		],
		"collapse":{
		   "field":"field_to_collapse_by"
		},
		"aggs":{
			"terms_aggregation_field":{
				"terms":{"field":"terms_aggregation_field","size":10}
			}
		}
	 }
	`
	expectedBody = strings.ReplaceAll(expectedBody, " ", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\t", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\n", "")
	body, err := effdsl.Define(
		effdsl.WithSourceFilter(
			effdsl.WithIncludes("field1", "field2"),
			effdsl.WithExcludes("field2", "field4"),
		),
		effdsl.WithPaginate(0, 10),
		effdsl.WithQuery(
			bq.BoolQuery(
				bq.Must(
					qs.QueryString("fake_value1", qs.WithFields("title", "description", "content")),
					qs.QueryString("fake_value2", qs.WithFields("title", "description", "content")),
				),
				bq.Filter(
					rq.RangeQuery("published_at", rq.WithGT("now-24h")),
					tq.TermQuery("field1.keyword", "fake_value"),
					eq.ExistsQuery("fake_field"),
				),
				bq.MustNot(
					qs.QueryString("fake_value3", qs.WithFields("title", "description", "content")),
					qs.QueryString("fake_value4", qs.WithFields("title", "description", "content")),
				),
				bq.Should(
					qs.QueryString("fake_value5", qs.WithFields("title", "description", "content")),
					qs.QueryString("fake_value6", qs.WithFields("title", "description", "content")),
				),
			),
		),
		effdsl.WithSort(
			effdsl.SortClause("sort_field1", effdsl.SORT_DESC),
			effdsl.SortClause("_score", effdsl.SORT_DEFAULT),
		),
		effdsl.WithCollpse("field_to_collapse_by"),
		effdsl.WithAggregations(
			effdsl.TermAggregation("terms_aggregation_field", 10),
		),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Equal(t, expectedBody, string(jsonBody))
}
