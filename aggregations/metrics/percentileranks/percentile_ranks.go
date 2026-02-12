package percentileranks

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type HDRHistogram struct {
	NumberOfSignificantValueDigits int `json:"number_of_significant_value_digits,omitempty"`
}

type PercentileRanksBody struct {
	Field   string               `json:"field,omitempty"`
	Values  []float64            `json:"values,omitempty"`
	Missing any                  `json:"missing,omitempty"`
	Script  *aggregations.Script `json:"script,omitempty"`
	Format  string               `json:"format,omitempty"`
	Keyed   *bool                `json:"keyed,omitempty"`
	HDR     *HDRHistogram        `json:"hdr,omitempty"`
	TDigest map[string]any       `json:"tdigest,omitempty"`
}

type PercentileRanksS struct {
	name string
	body PercentileRanksBody
	aggregations.BaseAggregation
}

func (rank *PercentileRanksS) AggregationName() string {
	return rank.name
}

func (rank *PercentileRanksS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("percentile_ranks", rank.body, rank.Extras())
}

type PercentileRanksOption = aggregations.Option[*PercentileRanksS]

func PercentileRanks(name, field string, values []float64, opts ...PercentileRanksOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &PercentileRanksS{
		name: name,
		body: PercentileRanksBody{
			Field:  field,
			Values: values,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: percentile_ranks aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.Script = &script
		return nil
	}
}

func WithFormat(format string) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.Format = format
		return nil
	}
}

func WithKeyed(keyed bool) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.Keyed = &keyed
		return nil
	}
}

func WithHDR(numberOfSignificantValueDigits int) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.HDR = &HDRHistogram{NumberOfSignificantValueDigits: numberOfSignificantValueDigits}
		return nil
	}
}

func WithTDigest(settings map[string]any) PercentileRanksOption {
	return func(p *PercentileRanksS) error {
		p.body.TDigest = settings
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) PercentileRanksOption {
	return aggregations.WithSubAggregation[*PercentileRanksS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) PercentileRanksOption {
	return aggregations.WithNamedSubAggregation[*PercentileRanksS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) PercentileRanksOption {
	return aggregations.WithSubAggregationsMap[*PercentileRanksS](subsMap)
}

func WithMetaField(key string, value any) PercentileRanksOption {
	return aggregations.WithMetaField[*PercentileRanksS](key, value)
}

func WithMetaMap(meta map[string]any) PercentileRanksOption {
	return aggregations.WithMetaMap[*PercentileRanksS](meta)
}
