package matrixstats

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MatrixStatsBody struct {
	Fields  []string       `json:"fields,omitempty"`
	Missing map[string]any `json:"missing,omitempty"`
	Mode    string         `json:"mode,omitempty"`
}

type MatrixStatsS struct {
	name string
	body MatrixStatsBody
	aggregations.BaseAggregation
}

func (stats *MatrixStatsS) AggregationName() string {
	return stats.name
}

func (stats *MatrixStatsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("matrix_stats", stats.body, stats.Extras())
}

type MatrixStatsOption = aggregations.Option[*MatrixStatsS]

func MatrixStats(name string, fields []string, opts ...MatrixStatsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &MatrixStatsS{
		name: name,
		body: MatrixStatsBody{
			Fields: fields,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing map[string]any) MatrixStatsOption {
	return func(m *MatrixStatsS) error {
		m.body.Missing = missing
		return nil
	}
}

func WithMode(mode string) MatrixStatsOption {
	return func(m *MatrixStatsS) error {
		m.body.Mode = mode
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MatrixStatsOption {
	return aggregations.WithSubAggregation[*MatrixStatsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MatrixStatsOption {
	return aggregations.WithNamedSubAggregation[*MatrixStatsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MatrixStatsOption {
	return aggregations.WithSubAggregationsMap[*MatrixStatsS](subsMap)
}

func WithMetaField(key string, value any) MatrixStatsOption {
	return aggregations.WithMetaField[*MatrixStatsS](key, value)
}

func WithMetaMap(meta map[string]any) MatrixStatsOption {
	return aggregations.WithMetaMap[*MatrixStatsS](meta)
}
