package derivative

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type DerivativeBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

type DerivativeS struct {
	name string
	body DerivativeBody
	aggregations.BaseAggregation
}

func (derivative *DerivativeS) AggregationName() string {
	return derivative.name
}

func (derivative *DerivativeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("derivative", derivative.body, derivative.Extras())
}

type DerivativeOption = aggregations.Option[*DerivativeS]

func Derivative(name, bucketsPath string, opts ...DerivativeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &DerivativeS{
		name: name,
		body: DerivativeBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) DerivativeOption {
	return func(d *DerivativeS) error {
		d.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) DerivativeOption {
	return func(d *DerivativeS) error {
		d.body.Format = format
		return nil
	}
}

func WithUnit(unit string) DerivativeOption {
	return func(d *DerivativeS) error {
		d.body.Unit = unit
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) DerivativeOption {
	return aggregations.WithSubAggregation[*DerivativeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) DerivativeOption {
	return aggregations.WithNamedSubAggregation[*DerivativeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) DerivativeOption {
	return aggregations.WithSubAggregationsMap[*DerivativeS](subsMap)
}

func WithMetaField(key string, value any) DerivativeOption {
	return aggregations.WithMetaField[*DerivativeS](key, value)
}

func WithMetaMap(meta map[string]any) DerivativeOption {
	return aggregations.WithMetaMap[*DerivativeS](meta)
}
