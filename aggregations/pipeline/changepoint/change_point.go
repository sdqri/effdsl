package changepoint

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ChangePointBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
}

type ChangePointS struct {
	name string
	body ChangePointBody
	aggregations.BaseAggregation
}

func (cp *ChangePointS) AggregationName() string {
	return cp.name
}

func (cp *ChangePointS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("change_point", cp.body, cp.Extras())
}

type ChangePointOption = aggregations.Option[*ChangePointS]

func ChangePoint(name, bucketsPath string, opts ...ChangePointOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if bucketsPath == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	agg := &ChangePointS{
		name: name,
		body: ChangePointBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSubAggregation(sub aggregations.AggregationResult) ChangePointOption {
	return aggregations.WithSubAggregation[*ChangePointS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ChangePointOption {
	return aggregations.WithNamedSubAggregation[*ChangePointS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ChangePointOption {
	return aggregations.WithSubAggregationsMap[*ChangePointS](subsMap)
}

func WithMetaField(key string, value any) ChangePointOption {
	return aggregations.WithMetaField[*ChangePointS](key, value)
}

func WithMetaMap(meta map[string]any) ChangePointOption {
	return aggregations.WithMetaMap[*ChangePointS](meta)
}
