package diversifiedsampler

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type DiversifiedSamplerBody struct {
	ShardSize       *int   `json:"shard_size,omitempty"`
	Field           string `json:"field,omitempty"`
	MaxDocsPerValue *int   `json:"max_docs_per_value,omitempty"`
	ExecutionHint   string `json:"execution_hint,omitempty"`
}

type DiversifiedSamplerS struct {
	name string
	body DiversifiedSamplerBody
	aggregations.BaseAggregation
}

func (sampler *DiversifiedSamplerS) AggregationName() string {
	return sampler.name
}

func (sampler *DiversifiedSamplerS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("diversified_sampler", sampler.body, sampler.Extras())
}

type DiversifiedSamplerOption = aggregations.Option[*DiversifiedSamplerS]

func DiversifiedSampler(name string, opts ...DiversifiedSamplerOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &DiversifiedSamplerS{name: name, body: DiversifiedSamplerBody{}}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithShardSize(size int) DiversifiedSamplerOption {
	return func(s *DiversifiedSamplerS) error {
		s.body.ShardSize = &size
		return nil
	}
}

func WithField(field string) DiversifiedSamplerOption {
	return func(s *DiversifiedSamplerS) error {
		s.body.Field = field
		return nil
	}
}

func WithMaxDocsPerValue(maxDocs int) DiversifiedSamplerOption {
	return func(s *DiversifiedSamplerS) error {
		s.body.MaxDocsPerValue = &maxDocs
		return nil
	}
}

func WithExecutionHint(hint string) DiversifiedSamplerOption {
	return func(s *DiversifiedSamplerS) error {
		s.body.ExecutionHint = hint
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) DiversifiedSamplerOption {
	return aggregations.WithSubAggregation[*DiversifiedSamplerS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) DiversifiedSamplerOption {
	return aggregations.WithNamedSubAggregation[*DiversifiedSamplerS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) DiversifiedSamplerOption {
	return aggregations.WithSubAggregationsMap[*DiversifiedSamplerS](subsMap)
}

func WithMetaField(key string, value any) DiversifiedSamplerOption {
	return aggregations.WithMetaField[*DiversifiedSamplerS](key, value)
}

func WithMetaMap(meta map[string]any) DiversifiedSamplerOption {
	return aggregations.WithMetaMap[*DiversifiedSamplerS](meta)
}
