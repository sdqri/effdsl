package objects_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestWithQuery(t *testing.T) {
	expectedBody := `{"query":{"query_string":{"query":"query","fields":["field1","field2"]}}}`
	f := objs.D.WithQuery(
		objs.D.QueryString("query", objs.D.WithFields("field1", "field2")),
	)
	body := objs.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithSort(t *testing.T) {
	expectedBody := `{"sort":["sort_field"]}`
	f := objs.D.WithSort(objs.D.SortClause("sort_field", objs.SORT_DEFAULT))
	body := objs.SearchBody{}
	f(&body)
	jsonBody, err := json.Marshal(body)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestWithCollapse(t *testing.T) {
	expectedBody := `{"collapse":{"field":"collapsed_field"}}`
	f := objs.D.WithCollpse("collapsed_field")
	body := objs.SearchBody{}
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
	body, err := objs.D(
		objs.D.WithSourceFilter(
			objs.D.WithIncludes("field1", "field2"),
			objs.D.WithExcludes("field2", "field4"),
		),
		objs.D.WithPaginate(0, 10),
		objs.D.WithQuery(
			objs.D.BoolQuery(
				objs.D.Must(
					objs.D.QueryString("fake_value1", objs.D.WithFields("title", "description", "content")),
					objs.D.QueryString("fake_value2", objs.D.WithFields("title", "description", "content")),
				),
				objs.D.Filter(
					objs.D.RangeQuery("published_at", objs.D.WithGT("now-24h")),
					objs.D.TermQuery("field1.keyword", "fake_value"),
					objs.D.ExistsQuery("fake_field"),
				),
				objs.D.MustNot(
					objs.D.QueryString("fake_value3", objs.D.WithFields("title", "description", "content")),
					objs.D.QueryString("fake_value4", objs.D.WithFields("title", "description", "content")),
				),
				objs.D.Should(
					objs.D.QueryString("fake_value5", objs.D.WithFields("title", "description", "content")),
					objs.D.QueryString("fake_value6", objs.D.WithFields("title", "description", "content")),
				),
			),
		),
		objs.D.WithSort(
			objs.D.SortClause("sort_field1", objs.SORT_DESC),
			objs.D.SortClause("_score", objs.SORT_DEFAULT),
		),
		objs.D.WithCollpse("field_to_collapse_by"),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Equal(t, expectedBody, string(jsonBody))
}
