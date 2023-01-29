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
		"from":0,
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
		effdsl.D.WithSourceFilter(
			effdsl.D.WithIncludes("field1", "field2"),
			effdsl.D.WithExcludes("field3", "field4"),
		),
		effdsl.D.WithPaginate(0, 100),
		effdsl.D.WithQuery(
			effdsl.D.BoolQuery(
				effdsl.D.Must(
					effdsl.D.QueryString("value1", effdsl.D.WithFields("title", "content")),
				),
				effdsl.D.Filter(
					effdsl.D.RangeQuery("published_at", effdsl.D.WithGT("now-24h")),
					effdsl.D.TermQuery("field5.keyword", "value2"),
					effdsl.D.ExistsQuery("field6"),
				),
				effdsl.D.MustNot(
					effdsl.D.QueryString("value3", effdsl.D.WithFields("title", "content")),
				),
				effdsl.D.Should(
					effdsl.D.QueryString("value4", effdsl.D.WithFields("title", "content")),
				),
			),
		),
		effdsl.D.WithSort(
			effdsl.D.SortClause("field1", effdsl.SORT_DESC),
			effdsl.D.SortClause("_score", effdsl.SORT_DEFAULT),
		),
		effdsl.D.WithCollpse("field7"),
	)
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(body)
	assert.Equal(t, expectedBody, string(jsonBody))
}
