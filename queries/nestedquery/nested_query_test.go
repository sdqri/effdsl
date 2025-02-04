package nestedquery_test

import (
	"encoding/json"
	"github.com/sdqri/effdsl/v2"
	bq "github.com/sdqri/effdsl/v2/queries/boolquery"
	mq "github.com/sdqri/effdsl/v2/queries/matchquery"
	nq "github.com/sdqri/effdsl/v2/queries/nestedquery"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedQuery_InQuery(t *testing.T) {
	expectedBody := `{
  "query": {
    "nested": {
      "path": "path",
      "query": {
        "bool": {
          "should": [
            {
              "match": {
                "field1": {
                  "query": "val1"
                }
              }
            }
          ]
        }
      },
      "score_mode": "avg",
      "ignore_unmapped": true
    }
  }
}`
	query, err := effdsl.Define(
		effdsl.WithQuery(
			nq.WithNested(
				"path",
				bq.BoolQuery(
					bq.Should(
						mq.MatchQuery("field1", "val1"),
					),
				),
				nq.WithScoreMode("avg"),
				nq.WithIgnoreUnmapped(true),
			),
		),
	)

	jsonBody, err := json.MarshalIndent(query, "", "  ")

	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}

func TestNestedQuery_InBool(t *testing.T) {
	expectedBody := `{
  "query": {
    "bool": {
      "must": [
        {
          "nested": {
            "path": "path",
            "query": {
              "bool": {
                "should": [
                  {
                    "match": {
                      "field1": {
                        "query": "val1"
                      }
                    }
                  }
                ]
              }
            }
          }
        }
      ]
    }
  }
}`
	query, err := effdsl.Define(
		effdsl.WithQuery(
			bq.BoolQuery(
				bq.Must(
					nq.WithNested(
						"path",
						bq.BoolQuery(
							bq.Should(
								mq.MatchQuery("field1", "val1"),
							),
						),
					),
				),
			),
		),
	)

	jsonBody, err := json.MarshalIndent(query, "", "  ")

	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
