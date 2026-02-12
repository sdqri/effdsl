package variablewidthhistogram

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type VariableWidthHistogramBody struct {
	Field     string `json:"field,omitempty"`
	Buckets   int    `json:"buckets,omitempty"`
	ShardSize *int   `json:"shard_size,omitempty"`
}

type VariableWidthHistogramS struct {
	name string
	body VariableWidthHistogramBody
	aggregations.BaseAggregation
}

func (histogram *VariableWidthHistogramS) AggregationName() string {
	return histogram.name
}

func (histogram *VariableWidthHistogramS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("variable_width_histogram", histogram.body, histogram.Extras())
}

type VariableWidthHistogramOption = aggregations.Option[*VariableWidthHistogramS]

func VariableWidthHistogram(name, field string, buckets int, opts ...VariableWidthHistogramOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &VariableWidthHistogramS{
		name: name,
		body: VariableWidthHistogramBody{
			Field:   field,
			Buckets: buckets,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithShardSize(size int) VariableWidthHistogramOption {
	return func(v *VariableWidthHistogramS) error {
		v.body.ShardSize = &size
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) VariableWidthHistogramOption {
	return aggregations.WithSubAggregation[*VariableWidthHistogramS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) VariableWidthHistogramOption {
	return aggregations.WithNamedSubAggregation[*VariableWidthHistogramS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) VariableWidthHistogramOption {
	return aggregations.WithSubAggregationsMap[*VariableWidthHistogramS](subsMap)
}

func WithMetaField(key string, value any) VariableWidthHistogramOption {
	return aggregations.WithMetaField[*VariableWidthHistogramS](key, value)
}

func WithMetaMap(meta map[string]any) VariableWidthHistogramOption {
	return aggregations.WithMetaMap[*VariableWidthHistogramS](meta)
}
