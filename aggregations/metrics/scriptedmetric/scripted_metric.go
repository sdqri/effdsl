package scriptedmetric

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ScriptedMetricBody struct {
	InitScript    *aggregations.Script `json:"init_script,omitempty"`
	MapScript     *aggregations.Script `json:"map_script,omitempty"`
	CombineScript *aggregations.Script `json:"combine_script,omitempty"`
	ReduceScript  *aggregations.Script `json:"reduce_script,omitempty"`
	Params        map[string]any       `json:"params,omitempty"`
}

type ScriptedMetricS struct {
	name string
	body ScriptedMetricBody
	aggregations.BaseAggregation
}

func (metric *ScriptedMetricS) AggregationName() string {
	return metric.name
}

func (metric *ScriptedMetricS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("scripted_metric", metric.body, metric.Extras())
}

type ScriptedMetricOption = aggregations.Option[*ScriptedMetricS]

func ScriptedMetric(name string, opts ...ScriptedMetricOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &ScriptedMetricS{name: name, body: ScriptedMetricBody{}}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithInitScript(script aggregations.Script) ScriptedMetricOption {
	return func(s *ScriptedMetricS) error {
		s.body.InitScript = &script
		return nil
	}
}

func WithMapScript(script aggregations.Script) ScriptedMetricOption {
	return func(s *ScriptedMetricS) error {
		s.body.MapScript = &script
		return nil
	}
}

func WithCombineScript(script aggregations.Script) ScriptedMetricOption {
	return func(s *ScriptedMetricS) error {
		s.body.CombineScript = &script
		return nil
	}
}

func WithReduceScript(script aggregations.Script) ScriptedMetricOption {
	return func(s *ScriptedMetricS) error {
		s.body.ReduceScript = &script
		return nil
	}
}

func WithParams(params map[string]any) ScriptedMetricOption {
	return func(s *ScriptedMetricS) error {
		s.body.Params = params
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) ScriptedMetricOption {
	return aggregations.WithSubAggregation[*ScriptedMetricS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ScriptedMetricOption {
	return aggregations.WithNamedSubAggregation[*ScriptedMetricS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ScriptedMetricOption {
	return aggregations.WithSubAggregationsMap[*ScriptedMetricS](subsMap)
}

func WithMetaField(key string, value any) ScriptedMetricOption {
	return aggregations.WithMetaField[*ScriptedMetricS](key, value)
}

func WithMetaMap(meta map[string]any) ScriptedMetricOption {
	return aggregations.WithMetaMap[*ScriptedMetricS](meta)
}
