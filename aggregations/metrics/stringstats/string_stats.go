package stringstats

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type StringStatsBody struct {
	Field            string               `json:"field,omitempty"`
	Missing          any                  `json:"missing,omitempty"`
	Script           *aggregations.Script `json:"script,omitempty"`
	ShowDistribution *bool                `json:"show_distribution,omitempty"`
}

type StringStatsS struct {
	name string
	body StringStatsBody
	aggregations.BaseAggregation
}

func (stats *StringStatsS) AggregationName() string {
	return stats.name
}

func (stats *StringStatsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("string_stats", stats.body, stats.Extras())
}

type StringStatsOption = aggregations.Option[*StringStatsS]

func StringStats(name, field string, opts ...StringStatsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &StringStatsS{
		name: name,
		body: StringStatsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: string_stats aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) StringStatsOption {
	return func(s *StringStatsS) error {
		s.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) StringStatsOption {
	return func(s *StringStatsS) error {
		s.body.Script = &script
		return nil
	}
}

func WithShowDistribution(show bool) StringStatsOption {
	return func(s *StringStatsS) error {
		s.body.ShowDistribution = &show
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) StringStatsOption {
	return aggregations.WithSubAggregation[*StringStatsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) StringStatsOption {
	return aggregations.WithNamedSubAggregation[*StringStatsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) StringStatsOption {
	return aggregations.WithSubAggregationsMap[*StringStatsS](subsMap)
}

func WithMetaField(key string, value any) StringStatsOption {
	return aggregations.WithMetaField[*StringStatsS](key, value)
}

func WithMetaMap(meta map[string]any) StringStatsOption {
	return aggregations.WithMetaMap[*StringStatsS](meta)
}
