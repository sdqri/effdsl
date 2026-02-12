package extendedstats

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ExtendedStatsBody struct {
	Field   string               `json:"field,omitempty"`
	Missing any                  `json:"missing,omitempty"`
	Script  *aggregations.Script `json:"script,omitempty"`
	Format  string               `json:"format,omitempty"`
	Sigma   *float64             `json:"sigma,omitempty"`
}

type ExtendedStatsS struct {
	name string
	body ExtendedStatsBody
	aggregations.BaseAggregation
}

func (stats *ExtendedStatsS) AggregationName() string {
	return stats.name
}

func (stats *ExtendedStatsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("extended_stats", stats.body, stats.Extras())
}

type ExtendedStatsOption = aggregations.Option[*ExtendedStatsS]

func ExtendedStats(name, field string, opts ...ExtendedStatsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &ExtendedStatsS{
		name: name,
		body: ExtendedStatsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: extended_stats aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) ExtendedStatsOption {
	return func(s *ExtendedStatsS) error {
		s.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) ExtendedStatsOption {
	return func(s *ExtendedStatsS) error {
		s.body.Script = &script
		return nil
	}
}

func WithFormat(format string) ExtendedStatsOption {
	return func(s *ExtendedStatsS) error {
		s.body.Format = format
		return nil
	}
}

func WithSigma(sigma float64) ExtendedStatsOption {
	return func(s *ExtendedStatsS) error {
		s.body.Sigma = &sigma
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) ExtendedStatsOption {
	return aggregations.WithSubAggregation[*ExtendedStatsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ExtendedStatsOption {
	return aggregations.WithNamedSubAggregation[*ExtendedStatsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ExtendedStatsOption {
	return aggregations.WithSubAggregationsMap[*ExtendedStatsS](subsMap)
}

func WithMetaField(key string, value any) ExtendedStatsOption {
	return aggregations.WithMetaField[*ExtendedStatsS](key, value)
}

func WithMetaMap(meta map[string]any) ExtendedStatsOption {
	return aggregations.WithMetaMap[*ExtendedStatsS](meta)
}
