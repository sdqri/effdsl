package effdsl_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sdqri/effdsl"
)

func TestDefindeQ1(t *testing.T) {
	expectedBody := `
	{
		"_source":{
		   "includes":[
			  "field1",
			  "field2"
		   ],
		   "excludes":[
			  "field3",
			  "field4"
		   ]
		},
		"from":1,
		"size":100,
		"query":{
		   "bool":{
			  "must":[
				 {
					"query_string":{
					   "query":"value1",
					   "fields":[
						  "title",
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
					   "field5.keyword":{
						  "value":"value2"
					   }
					}
				 },
				 {
					"exists":{
					   "field":"field6"
					}
				 }
			  ],
			  "must_not":[
				 {
					"query_string":{
					   "query":"value3",
					   "fields":[
						  "title",
						  "content"
					   ]
					}
				 }
			  ],
			  "should":[
				 {
					"query_string":{
					   "query":"value4",
					   "fields":[
						  "title",
						  "content"
					   ]
					}
				 }
			  ]
		   }
		},
		"sort":[
		   {
			  "field1":"desc"
		   },
		   "_score"
		],
		"collapse":{
		   "field":"field7"
		}
	 }
	`
	expectedBody = strings.ReplaceAll(expectedBody, " ", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\t", "")
	expectedBody = strings.ReplaceAll(expectedBody, "\n", "")
	body, err := effdsl.Define(
		effdsl.WithSourceFilter(
			effdsl.WithIncludes("field1", "field2"),
			effdsl.WithExcludes("field3", "field4"),
		),
		effdsl.WithPaginate(1, 100),
		effdsl.WithQuery(
			effdsl.BoolQuery(
				effdsl.Must(
					effdsl.QueryString("value1", effdsl.WithFields("title", "content")),
				),
				effdsl.Filter(
					effdsl.RangeQuery("published_at", effdsl.WithGT("now-24h")),
					effdsl.TermQuery("field5.keyword", "value2"),
					effdsl.ExistsQuery("field6"),
				),
				effdsl.MustNot(
					effdsl.QueryString("value3", effdsl.WithFields("title", "content")),
				),
				effdsl.Should(
					effdsl.QueryString("value4", effdsl.WithFields("title", "content")),
				),
			),
		),
		effdsl.WithSort(
			effdsl.SortClause("field1", effdsl.SORT_DESC),
			effdsl.SortClause("_score", effdsl.SORT_DEFAULT),
		),
		effdsl.WithCollpse("field7"),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Equal(t, expectedBody, string(jsonBody))
}
