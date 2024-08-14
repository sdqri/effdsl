package idsquery

import (
	"encoding/json"

	"github.com/sdqri/effdsl"
)

// IDsQueryS represents the structure of the IDs query.
type IDsQueryS struct {
	Values []string `json:"values"` // (Required, array of strings) An array of document IDs.
}

func (iq IDsQueryS) QueryInfo() string {
	return "IDs query"
}

func (iq IDsQueryS) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		effdsl.M{
			"ids": effdsl.M{
				"values": iq.Values,
			},
		},
	)
}

// IDsQuery Returns documents based on their IDs.
//
// For more details, see the official Elasticsearch documentation:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html

func IDsQuery(values ...string) effdsl.QueryResult {
	idsQuery := IDsQueryS{
		Values: values,
	}
	return effdsl.QueryResult{
		Ok:  idsQuery,
		Err: nil,
	}
}
