package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	effdsl "github.com/sdqri/effdsl"
)

func TestWithQuery(t *testing.T) {
	expectedBody := `{"query":{"query_string":{"query":"query","fields":["field1","field2"]}}}`
	f := effdsl.WithQuery(
		effdsl.QueryString("query", effdsl.WithFields("field1", "field2")),
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

func TestWithSuggest(t *testing.T) {
	expectedBody := `{"suggest":{"text":"test","my-suggestion-1":{"text":"tring out Elasticsearch","term":{"field":"message"}}}}`
	f := effdsl.WithSuggest(
		effdsl.Suggest("test",
			effdsl.TermSuggester("my-suggestion-1", "tring out Elasticsearch", "message"),
		),
	)
	body := effdsl.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
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
			  ]
		   }
		},
		"sort":[
		   {
			  "sort_field1":"desc"
		   },
		   "_score"
		],
		"collapse":{
		   "field":"field_to_collapse_by"
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
			effdsl.BoolQuery(
				effdsl.Must(
					effdsl.QueryString("fake_value1", effdsl.WithFields("title", "description", "content")),
					effdsl.QueryString("fake_value2", effdsl.WithFields("title", "description", "content")),
				),
				effdsl.Filter(
					effdsl.RangeQuery("published_at", effdsl.WithGT("now-24h")),
					effdsl.TermQuery("field1.keyword", "fake_value"),
					effdsl.ExistsQuery("fake_field"),
				),
				effdsl.MustNot(
					effdsl.QueryString("fake_value3", effdsl.WithFields("title", "description", "content")),
					effdsl.QueryString("fake_value4", effdsl.WithFields("title", "description", "content")),
				),
				effdsl.Should(
					effdsl.QueryString("fake_value5", effdsl.WithFields("title", "description", "content")),
					effdsl.QueryString("fake_value6", effdsl.WithFields("title", "description", "content")),
				),
			),
		),
		effdsl.WithSort(
			effdsl.SortClause("sort_field1", effdsl.SORT_DESC),
			effdsl.SortClause("_score", effdsl.SORT_DEFAULT),
		),
		effdsl.WithCollpse("field_to_collapse_by"),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Equal(t, expectedBody, string(jsonBody))
}
