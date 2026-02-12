package cardinality

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CardinalityBody struct {
	Field              string `json:"field,omitempty"`
	Missing            any    `json:"missing,omitempty"`
	PrecisionThreshold int    `json:"precision_threshold,omitempty"`
	ExecutionHint      string `json:"execution_hint,omitempty"`
}

type CardinalityS struct {
	name string
	body CardinalityBody
	aggregations.BaseAggregation
}

func (cardinalityS *CardinalityS) AggregationName() string {
	return cardinalityS.name
}

func (cardinalityS *CardinalityS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("cardinality", cardinalityS.body, cardinalityS.Extras())
}

type CardinalityOption = aggregations.Option[*CardinalityS]

func Cardinality(name, field string, opts ...CardinalityOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &CardinalityS{
		name: name,
		body: CardinalityBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMissing(missing any) CardinalityOption {
	return func(c *CardinalityS) error {
		c.body.Missing = missing
		return nil
	}
}

func WithPrecisionThreshold(threshold int) CardinalityOption {
	return func(c *CardinalityS) error {
		c.body.PrecisionThreshold = threshold
		return nil
	}
}

func WithExecutionHint(hint string) CardinalityOption {
	return func(c *CardinalityS) error {
		c.body.ExecutionHint = hint
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) CardinalityOption {
	return aggregations.WithSubAggregation[*CardinalityS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CardinalityOption {
	return aggregations.WithNamedSubAggregation[*CardinalityS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CardinalityOption {
	return aggregations.WithSubAggregationsMap[*CardinalityS](subsMap)
}

func WithMetaField(key string, value any) CardinalityOption {
	return aggregations.WithMetaField[*CardinalityS](key, value)
}

func WithMetaMap(meta map[string]any) CardinalityOption {
	return aggregations.WithMetaMap[*CardinalityS](meta)
}
