package topmetrics

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type TopMetricsBody struct {
	Metrics any  `json:"metrics,omitempty"`
	Sort    any  `json:"sort,omitempty"`
	Size    *int `json:"size,omitempty"`
}

type TopMetricsS struct {
	name string
	body TopMetricsBody
	aggregations.BaseAggregation
}

func (metrics *TopMetricsS) AggregationName() string {
	return metrics.name
}

func (metrics *TopMetricsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("top_metrics", metrics.body, metrics.Extras())
}

type TopMetricsOption = aggregations.Option[*TopMetricsS]

func TopMetrics(name string, metricsField any, sort any, opts ...TopMetricsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &TopMetricsS{
		name: name,
		body: TopMetricsBody{
			Metrics: metricsField,
			Sort:    sort,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) TopMetricsOption {
	return func(t *TopMetricsS) error {
		t.body.Size = &size
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) TopMetricsOption {
	return aggregations.WithSubAggregation[*TopMetricsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) TopMetricsOption {
	return aggregations.WithNamedSubAggregation[*TopMetricsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) TopMetricsOption {
	return aggregations.WithSubAggregationsMap[*TopMetricsS](subsMap)
}

func WithMetaField(key string, value any) TopMetricsOption {
	return aggregations.WithMetaField[*TopMetricsS](key, value)
}

func WithMetaMap(meta map[string]any) TopMetricsOption {
	return aggregations.WithMetaMap[*TopMetricsS](meta)
}
