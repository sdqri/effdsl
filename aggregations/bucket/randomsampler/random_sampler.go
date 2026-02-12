package randomsampler

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type RandomSamplerBody struct {
	Probability float64 `json:"probability,omitempty"`
	Seed        *int    `json:"seed,omitempty"`
}

type RandomSamplerS struct {
	name string
	body RandomSamplerBody
	aggregations.BaseAggregation
}

func (sampler *RandomSamplerS) AggregationName() string {
	return sampler.name
}

func (sampler *RandomSamplerS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("random_sampler", sampler.body, sampler.Extras())
}

type RandomSamplerOption = aggregations.Option[*RandomSamplerS]

func RandomSampler(name string, probability float64, opts ...RandomSamplerOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &RandomSamplerS{
		name: name,
		body: RandomSamplerBody{Probability: probability},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSeed(seed int) RandomSamplerOption {
	return func(r *RandomSamplerS) error {
		r.body.Seed = &seed
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) RandomSamplerOption {
	return aggregations.WithSubAggregation[*RandomSamplerS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) RandomSamplerOption {
	return aggregations.WithNamedSubAggregation[*RandomSamplerS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) RandomSamplerOption {
	return aggregations.WithSubAggregationsMap[*RandomSamplerS](subsMap)
}

func WithMetaField(key string, value any) RandomSamplerOption {
	return aggregations.WithMetaField[*RandomSamplerS](key, value)
}

func WithMetaMap(meta map[string]any) RandomSamplerOption {
	return aggregations.WithMetaMap[*RandomSamplerS](meta)
}
