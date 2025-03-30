package effdsl

import (
	"encoding/json"
)

const (
	AggregationTypeTerms string = "terms"
	AggregationTypeStats string = "stats"
)

type aggregation struct {
	Field           string
	aggregationType string
}

func (a aggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(M{
		a.aggregationType: M{
			"field": a.GetAggregationField(),
		},
	})
}

func (a aggregation) GetAggregationField() string {
	return a.Field
}

type TermsAggregationS struct {
	aggregation
	Size *int
}

func (t TermsAggregationS) MarshalJSON() ([]byte, error) {

	params := M{
		"field": t.GetAggregationField(),
	}
	if t.Size != nil {
		params["size"] = t.Size
	}
	return json.Marshal(M{
		t.aggregationType: params,
	})
}

func TermAggregation(field string, size ...int) TermsAggregationS {
	termsAgg := TermsAggregationS{aggregation: aggregation{
		Field:           field,
		aggregationType: AggregationTypeTerms,
	}}
	if len(size) > 0 {
		termsAgg.Size = &size[0]
	}
	return termsAgg
}

type StatsAggregationS struct {
	aggregation
	Size *int
}

func StatsAggregation(field string) StatsAggregationS {
	return StatsAggregationS{aggregation: aggregation{
		Field:           field,
		aggregationType: AggregationTypeStats,
	}}
}
