package stats

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type StatsBody struct {
	Field   string               `json:"field,omitempty"`
	Missing any                  `json:"missing,omitempty"`
	Script  *aggregations.Script `json:"script,omitempty"`
	Format  string               `json:"format,omitempty"`
}

type StatsS struct {
	name string
	body StatsBody
	aggregations.BaseAggregation
}

func (stats *StatsS) AggregationName() string {
	return stats.name
}

func (stats *StatsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("stats", stats.body, stats.Extras())
}

type StatsOption = aggregations.Option[*StatsS]

func Stats(name, field string, opts ...StatsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &StatsS{
		name: name,
		body: StatsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	if agg.body.Script != nil && agg.body.Field != "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: stats aggregation cannot have both field and script set")}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) StatsOption {
	return func(s *StatsS) error {
		s.body.Missing = missing
		return nil
	}
}

func WithScript(script aggregations.Script) StatsOption {
	return func(s *StatsS) error {
		s.body.Script = &script
		return nil
	}
}

func WithFormat(format string) StatsOption {
	return func(s *StatsS) error {
		s.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) StatsOption {
	return aggregations.WithSubAggregation[*StatsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) StatsOption {
	return aggregations.WithNamedSubAggregation[*StatsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) StatsOption {
	return aggregations.WithSubAggregationsMap[*StatsS](subsMap)
}

func WithMetaField(key string, value any) StatsOption {
	return aggregations.WithMetaField[*StatsS](key, value)
}

func WithMetaMap(meta map[string]any) StatsOption {
	return aggregations.WithMetaMap[*StatsS](meta)
}
