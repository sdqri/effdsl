package medianabsolutedeviation

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MedianAbsoluteDeviationBody struct {
	Field       string `json:"field,omitempty"`
	Missing     any    `json:"missing,omitempty"`
	Compression int    `json:"compression,omitempty"`
}

type MedianAbsoluteDeviationS struct {
	name string
	body MedianAbsoluteDeviationBody
	aggregations.BaseAggregation
}

func (mad *MedianAbsoluteDeviationS) AggregationName() string {
	return mad.name
}

func (mad *MedianAbsoluteDeviationS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("median_absolute_deviation", mad.body, mad.Extras())
}

type MedianAbsoluteDeviationOption = aggregations.Option[*MedianAbsoluteDeviationS]

func MedianAbsoluteDeviation(name, field string, opts ...MedianAbsoluteDeviationOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &MedianAbsoluteDeviationS{
		name: name,
		body: MedianAbsoluteDeviationBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) MedianAbsoluteDeviationOption {
	return func(m *MedianAbsoluteDeviationS) error {
		m.body.Missing = missing
		return nil
	}
}

func WithCompression(compression int) MedianAbsoluteDeviationOption {
	return func(m *MedianAbsoluteDeviationS) error {
		m.body.Compression = compression
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MedianAbsoluteDeviationOption {
	return aggregations.WithSubAggregation[*MedianAbsoluteDeviationS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MedianAbsoluteDeviationOption {
	return aggregations.WithNamedSubAggregation[*MedianAbsoluteDeviationS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MedianAbsoluteDeviationOption {
	return aggregations.WithSubAggregationsMap[*MedianAbsoluteDeviationS](subsMap)
}

func WithMetaField(key string, value any) MedianAbsoluteDeviationOption {
	return aggregations.WithMetaField[*MedianAbsoluteDeviationS](key, value)
}

func WithMetaMap(meta map[string]any) MedianAbsoluteDeviationOption {
	return aggregations.WithMetaMap[*MedianAbsoluteDeviationS](meta)
}
