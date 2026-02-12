package sampler

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SamplerBody struct {
	ShardSize *int `json:"shard_size,omitempty"`
}

type SamplerS struct {
	name string
	body SamplerBody
	aggregations.BaseAggregation
}

func (sampler *SamplerS) AggregationName() string {
	return sampler.name
}

func (sampler *SamplerS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("sampler", sampler.body, sampler.Extras())
}

type SamplerOption = aggregations.Option[*SamplerS]

func Sampler(name string, opts ...SamplerOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &SamplerS{name: name, body: SamplerBody{}}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithShardSize(size int) SamplerOption {
	return func(s *SamplerS) error {
		s.body.ShardSize = &size
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SamplerOption {
	return aggregations.WithSubAggregation[*SamplerS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SamplerOption {
	return aggregations.WithNamedSubAggregation[*SamplerS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SamplerOption {
	return aggregations.WithSubAggregationsMap[*SamplerS](subsMap)
}

func WithMetaField(key string, value any) SamplerOption {
	return aggregations.WithMetaField[*SamplerS](key, value)
}

func WithMetaMap(meta map[string]any) SamplerOption {
	return aggregations.WithMetaMap[*SamplerS](meta)
}
