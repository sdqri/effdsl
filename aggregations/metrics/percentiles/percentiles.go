package percentiles

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type HDRHistogram struct {
	NumberOfSignificantValueDigits int `json:"number_of_significant_value_digits,omitempty"`
}

type PercentilesBody struct {
	Field    string               `json:"field,omitempty"`
	Percents []float64            `json:"percents,omitempty"`
	Missing  any                  `json:"missing,omitempty"`
	Script   *aggregations.Script `json:"script,omitempty"`
	Format   string               `json:"format,omitempty"`
	Keyed    *bool                `json:"keyed,omitempty"`
	HDR      *HDRHistogram        `json:"hdr,omitempty"`
	TDigest  map[string]any       `json:"tdigest,omitempty"`
}

type PercentilesS struct {
	name string
	body PercentilesBody
	aggregations.BaseAggregation
}

func (percentiles *PercentilesS) AggregationName() string {
	return percentiles.name
}

func (percentiles *PercentilesS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("percentiles", percentiles.body, percentiles.Extras())
}

type PercentilesOption = aggregations.Option[*PercentilesS]

func Percentiles(name, field string, percents []float64, opts ...PercentilesOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &PercentilesS{
		name: name,
		body: PercentilesBody{
			Field:    field,
			Percents: percents,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: percentiles aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.Script = &script
		return nil
	}
}

func WithFormat(format string) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.Format = format
		return nil
	}
}

func WithKeyed(keyed bool) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.Keyed = &keyed
		return nil
	}
}

func WithHDR(numberOfSignificantValueDigits int) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.HDR = &HDRHistogram{NumberOfSignificantValueDigits: numberOfSignificantValueDigits}
		return nil
	}
}

func WithTDigest(settings map[string]any) PercentilesOption {
	return func(p *PercentilesS) error {
		p.body.TDigest = settings
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) PercentilesOption {
	return aggregations.WithSubAggregation[*PercentilesS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) PercentilesOption {
	return aggregations.WithNamedSubAggregation[*PercentilesS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) PercentilesOption {
	return aggregations.WithSubAggregationsMap[*PercentilesS](subsMap)
}

func WithMetaField(key string, value any) PercentilesOption {
	return aggregations.WithMetaField[*PercentilesS](key, value)
}

func WithMetaMap(meta map[string]any) PercentilesOption {
	return aggregations.WithMetaMap[*PercentilesS](meta)
}
