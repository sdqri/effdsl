package cumulativesum

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CumulativeSumBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Format      string `json:"format,omitempty"`
}

type CumulativeSumS struct {
	name string
	body CumulativeSumBody
	aggregations.BaseAggregation
}

func (sum *CumulativeSumS) AggregationName() string {
	return sum.name
}

func (sum *CumulativeSumS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("cumulative_sum", sum.body, sum.Extras())
}

type CumulativeSumOption = aggregations.Option[*CumulativeSumS]

func CumulativeSum(name, bucketsPath string, opts ...CumulativeSumOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &CumulativeSumS{
		name: name,
		body: CumulativeSumBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) CumulativeSumOption {
	return func(c *CumulativeSumS) error {
		c.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) CumulativeSumOption {
	return aggregations.WithSubAggregation[*CumulativeSumS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) CumulativeSumOption {
	return aggregations.WithNamedSubAggregation[*CumulativeSumS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) CumulativeSumOption {
	return aggregations.WithSubAggregationsMap[*CumulativeSumS](subsMap)
}

func WithMetaField(key string, value any) CumulativeSumOption {
	return aggregations.WithMetaField[*CumulativeSumS](key, value)
}

func WithMetaMap(meta map[string]any) CumulativeSumOption {
	return aggregations.WithMetaMap[*CumulativeSumS](meta)
}
