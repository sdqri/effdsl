package ttest

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type TTestPopulation struct {
	Field  string         `json:"field,omitempty"`
	Filter map[string]any `json:"filter,omitempty"`
}

type TTestBody struct {
	A    TTestPopulation `json:"a,omitempty"`
	B    TTestPopulation `json:"b,omitempty"`
	Type string          `json:"type,omitempty"`
}

type TTestS struct {
	name string
	body TTestBody
	aggregations.BaseAggregation
}

func (test *TTestS) AggregationName() string {
	return test.name
}

func (test *TTestS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("t_test", test.body, test.Extras())
}

type TTestOption = aggregations.Option[*TTestS]

func TTest(name, aField, bField string, opts ...TTestOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &TTestS{
		name: name,
		body: TTestBody{
			A: TTestPopulation{Field: aField},
			B: TTestPopulation{Field: bField},
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithAFilter(filter map[string]any) TTestOption {
	return func(t *TTestS) error {
		t.body.A.Filter = filter
		return nil
	}
}

func WithBFilter(filter map[string]any) TTestOption {
	return func(t *TTestS) error {
		t.body.B.Filter = filter
		return nil
	}
}

func WithType(testType string) TTestOption {
	return func(t *TTestS) error {
		t.body.Type = testType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) TTestOption {
	return aggregations.WithSubAggregation[*TTestS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) TTestOption {
	return aggregations.WithNamedSubAggregation[*TTestS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) TTestOption {
	return aggregations.WithSubAggregationsMap[*TTestS](subsMap)
}

func WithMetaField(key string, value any) TTestOption {
	return aggregations.WithMetaField[*TTestS](key, value)
}

func WithMetaMap(meta map[string]any) TTestOption {
	return aggregations.WithMetaMap[*TTestS](meta)
}
