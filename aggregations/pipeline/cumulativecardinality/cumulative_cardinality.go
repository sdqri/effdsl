package cumulativecardinality

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CumulativeCardinalityBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Format      string `json:"format,omitempty"`
}

type CumulativeCardinalityS struct {
	name string
	body CumulativeCardinalityBody
	aggregations.BaseAggregation
}

func (cardinality *CumulativeCardinalityS) AggregationName() string {
	return cardinality.name
}

func (cardinality *CumulativeCardinalityS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("cumulative_cardinality", cardinality.body, cardinality.Extras())
}

type CumulativeCardinalityOption = aggregations.Option[*CumulativeCardinalityS]

func CumulativeCardinality(name, bucketsPath string, opts ...CumulativeCardinalityOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &CumulativeCardinalityS{
		name: name,
		body: CumulativeCardinalityBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) CumulativeCardinalityOption {
	return func(c *CumulativeCardinalityS) error {
		c.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) CumulativeCardinalityOption {
	return aggregations.WithSubAggregation[*CumulativeCardinalityS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CumulativeCardinalityOption {
	return aggregations.WithNamedSubAggregation[*CumulativeCardinalityS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CumulativeCardinalityOption {
	return aggregations.WithSubAggregationsMap[*CumulativeCardinalityS](subsMap)
}

func WithMetaField(key string, value any) CumulativeCardinalityOption {
	return aggregations.WithMetaField[*CumulativeCardinalityS](key, value)
}

func WithMetaMap(meta map[string]any) CumulativeCardinalityOption {
	return aggregations.WithMetaMap[*CumulativeCardinalityS](meta)
}
