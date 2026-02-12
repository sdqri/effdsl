package movingfn

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MovingFnBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Window      *int   `json:"window,omitempty"`
	Script      string `json:"script,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Shift       *int   `json:"shift,omitempty"`
}

type MovingFnS struct {
	name string
	body MovingFnBody
	aggregations.BaseAggregation
}

func (fn *MovingFnS) AggregationName() string {
	return fn.name
}

func (fn *MovingFnS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("moving_fn", fn.body, fn.Extras())
}

type MovingFnOption = aggregations.Option[*MovingFnS]

func MovingFn(name, bucketsPath string, window int, script string, opts ...MovingFnOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if window <= 0 {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: moving_fn requires window greater than zero")}
	}

	if script == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: moving_fn requires script")}
	}

	agg := &MovingFnS{
		name: name,
		body: MovingFnBody{
			BucketsPath: bucketsPath,
			Window:      &window,
			Script:      script,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) MovingFnOption {
	return func(m *MovingFnS) error {
		m.body.GapPolicy = policy
		return nil
	}
}

func WithShift(shift int) MovingFnOption {
	return func(m *MovingFnS) error {
		m.body.Shift = &shift
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MovingFnOption {
	return aggregations.WithSubAggregation[*MovingFnS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MovingFnOption {
	return aggregations.WithNamedSubAggregation[*MovingFnS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MovingFnOption {
	return aggregations.WithSubAggregationsMap[*MovingFnS](subsMap)
}

func WithMetaField(key string, value any) MovingFnOption {
	return aggregations.WithMetaField[*MovingFnS](key, value)
}

func WithMetaMap(meta map[string]any) MovingFnOption {
	return aggregations.WithMetaMap[*MovingFnS](meta)
}
