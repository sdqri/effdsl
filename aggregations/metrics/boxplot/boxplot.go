package boxplot

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type BoxplotBody struct {
	Field         string `json:"field,omitempty"`
	Missing       any    `json:"missing,omitempty"`
	Compression   int    `json:"compression,omitempty"`
	ExecutionHint string `json:"execution_hint,omitempty"`
}

type BoxplotS struct {
	name string
	body BoxplotBody
	aggregations.BaseAggregation
}

func (boxplotS *BoxplotS) AggregationName() string {
	return boxplotS.name
}

func (boxplotS *BoxplotS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("boxplot", boxplotS.body, boxplotS.Extras())
}

type BoxplotOption = aggregations.Option[*BoxplotS]

func Boxplot(name, field string, opts ...BoxplotOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &BoxplotS{
		name: name,
		body: BoxplotBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) BoxplotOption {
	return func(b *BoxplotS) error {
		b.body.Missing = missing
		return nil
	}
}

func WithCompression(compression int) BoxplotOption {
	return func(b *BoxplotS) error {
		b.body.Compression = compression
		return nil
	}
}

func WithExecutionHint(hint string) BoxplotOption {
	return func(b *BoxplotS) error {
		b.body.ExecutionHint = hint
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BoxplotOption {
	return aggregations.WithSubAggregation[*BoxplotS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BoxplotOption {
	return aggregations.WithNamedSubAggregation[*BoxplotS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BoxplotOption {
	return aggregations.WithSubAggregationsMap[*BoxplotS](subsMap)
}

func WithMetaField(key string, value any) BoxplotOption {
	return aggregations.WithMetaField[*BoxplotS](key, value)
}

func WithMetaMap(meta map[string]any) BoxplotOption {
	return aggregations.WithMetaMap[*BoxplotS](meta)
}
