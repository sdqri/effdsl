package aggregations

import (
	"encoding/json"
	"fmt"
)

//  Aggregation Interfaces & Types

type AggregationType interface {
	AggregationName() string
	Extras() *AggregationExtras
	json.Marshaler
}

type AggregationResult struct {
	Ok  AggregationType
	Err error
}

//  Options & ApplyOptions

type Option[T AggregationType] func(T) error

func ApplyOptions[T AggregationType](agg T, opts ...Option[T]) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(agg); err != nil {
			return err
		}
	}
	return nil
}

//  Sub-Aggregation Option Helpers

func withSubAggregation[T AggregationType](name string, sub AggregationResult, named bool) Option[T] {
	return func(agg T) error {
		if sub.Err != nil {
			return fmt.Errorf("failed to add sub-aggregation: %w", sub.Err)
		}

		subAggregation := sub.Ok
		if subAggregation == nil {
			return fmt.Errorf("sub-aggregation cannot be nil")
		}

		if named {
			if name == "" {
				return fmt.Errorf("sub-aggregation name cannot be empty")
			}
			agg.Extras().AddNamedSubAggregation(name, subAggregation)
			return nil
		}

		agg.Extras().AddSubAggregation(subAggregation)
		return nil
	}
}

func WithSubAggregation[T AggregationType](sub AggregationResult) Option[T] {
	return withSubAggregation[T]("", sub, false)
}

func WithNamedSubAggregation[T AggregationType](name string, sub AggregationResult) Option[T] {
	return withSubAggregation[T](name, sub, true)
}

func WithSubAggregationsMap[T AggregationType](subs map[string]AggregationResult) Option[T] {
	return func(agg T) error {
		if subs == nil {
			return nil
		}

		for name, res := range subs {
			if name == "" {
				continue
			}

			if res.Err != nil {
				return fmt.Errorf("failed to add sub-aggregation %q: %w", name, res.Err)
			}

			if res.Ok == nil {
				return fmt.Errorf("sub-aggregation %q cannot be nil", name)
			}

			agg.Extras().AddNamedSubAggregation(name, res.Ok)
		}

		return nil
	}
}

// Metadata Option Helpers
func WithMetaField[T AggregationType](key string, value any) Option[T] {
	return func(agg T) error {
		agg.Extras().AddMetaField(key, value)
		return nil
	}
}

func WithMetaMap[T AggregationType](meta map[string]any) Option[T] {
	return func(agg T) error {
		agg.Extras().SetMeta(meta)
		return nil
	}
}

//  AggregationExtras

type AggregationExtras struct {
	SubAggregations map[string]AggregationType `json:"aggs,omitempty"`
	Meta            map[string]any             `json:"meta,omitempty"`
}

func (extras *AggregationExtras) AddSubAggregation(agg AggregationType) {
	if agg == nil {
		return
	}
	extras.AddNamedSubAggregation(agg.AggregationName(), agg)
}

func (extras *AggregationExtras) AddNamedSubAggregation(name string, agg AggregationType) {
	if agg == nil || name == "" {
		return
	}

	if extras.SubAggregations == nil {
		extras.SubAggregations = make(map[string]AggregationType)
	}

	extras.SubAggregations[name] = agg
}

func (extras *AggregationExtras) SetSubAggregations(aggs map[string]AggregationType) {
	if aggs == nil {
		extras.SubAggregations = nil
		return
	}

	extras.SubAggregations = make(map[string]AggregationType, len(aggs))
	for name, agg := range aggs {
		if agg == nil || name == "" {
			continue
		}
		extras.SubAggregations[name] = agg
	}
}

func (extras *AggregationExtras) AddMetaField(key string, value any) {
	if key == "" {
		return
	}

	if extras.Meta == nil {
		extras.Meta = make(map[string]any)
	}

	extras.Meta[key] = value
}

func (extras *AggregationExtras) SetMeta(meta map[string]any) {
	if meta == nil {
		extras.Meta = nil
		return
	}

	if extras.Meta == nil {
		extras.Meta = make(map[string]any, len(meta))
	} else {
		for k := range extras.Meta {
			delete(extras.Meta, k)
		}
	}

	for k, v := range meta {
		extras.Meta[k] = v
	}
}

//  BaseAggregation

type BaseAggregation struct {
	extras AggregationExtras
}

func (b *BaseAggregation) Extras() *AggregationExtras {
	return &b.extras
}

func (b *BaseAggregation) AddSubAggregation(agg AggregationType) {
	b.extras.AddSubAggregation(agg)
}

func (b *BaseAggregation) AddNamedSubAggregation(name string, agg AggregationType) {
	b.extras.AddNamedSubAggregation(name, agg)
}

func (b *BaseAggregation) SetSubAggregations(aggs map[string]AggregationType) {
	b.extras.SetSubAggregations(aggs)
}

func (b *BaseAggregation) SetMeta(meta map[string]any) {
	b.extras.SetMeta(meta)
}

func (b *BaseAggregation) AddMetaField(key string, value any) {
	b.extras.AddMetaField(key, value)
}

func marshalAggregationBody(aggType string, body any, data *AggregationExtras) ([]byte, error) {
	payload := make(map[string]any)

	if aggType != "" {
		payload[aggType] = body
	}

	if data != nil {
		if len(data.SubAggregations) > 0 {
			payload["aggs"] = data.SubAggregations
		}
		if len(data.Meta) > 0 {
			payload["meta"] = data.Meta
		}
	}

	return json.Marshal(payload)
}

func MarshalAggregation(aggType string, body any, extras *AggregationExtras) ([]byte, error) {
	return marshalAggregationBody(aggType, body, extras)
}
