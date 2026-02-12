package normalize

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type NormalizeBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Method      string `json:"method,omitempty"`
	Format      string `json:"format,omitempty"`
}

type NormalizeS struct {
	name string
	body NormalizeBody
	aggregations.BaseAggregation
}

func (norm *NormalizeS) AggregationName() string {
	return norm.name
}

func (norm *NormalizeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("normalize", norm.body, norm.Extras())
}

type NormalizeOption = aggregations.Option[*NormalizeS]

func Normalize(name, bucketsPath, method string, opts ...NormalizeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if method == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: normalize aggregation requires method")}
	}

	agg := &NormalizeS{
		name: name,
		body: NormalizeBody{
			BucketsPath: bucketsPath,
			Method:      method,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) NormalizeOption {
	return func(n *NormalizeS) error {
		n.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) NormalizeOption {
	return aggregations.WithSubAggregation[*NormalizeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) NormalizeOption {
	return aggregations.WithNamedSubAggregation[*NormalizeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) NormalizeOption {
	return aggregations.WithSubAggregationsMap[*NormalizeS](subsMap)
}

func WithMetaField(key string, value any) NormalizeOption {
	return aggregations.WithMetaField[*NormalizeS](key, value)
}

func WithMetaMap(meta map[string]any) NormalizeOption {
	return aggregations.WithMetaMap[*NormalizeS](meta)
}
