package sumbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SumBucketBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type SumBucketS struct {
	name string
	body SumBucketBody
	aggregations.BaseAggregation
}

func (sum *SumBucketS) AggregationName() string {
	return sum.name
}

func (sum *SumBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("sum_bucket", sum.body, sum.Extras())
}

type SumBucketOption = aggregations.Option[*SumBucketS]

func SumBucket(name, bucketsPath string, opts ...SumBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &SumBucketS{
		name: name,
		body: SumBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) SumBucketOption {
	return func(s *SumBucketS) error {
		s.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) SumBucketOption {
	return func(s *SumBucketS) error {
		s.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SumBucketOption {
	return aggregations.WithSubAggregation[*SumBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SumBucketOption {
	return aggregations.WithNamedSubAggregation[*SumBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SumBucketOption {
	return aggregations.WithSubAggregationsMap[*SumBucketS](subsMap)
}

func WithMetaField(key string, value any) SumBucketOption {
	return aggregations.WithMetaField[*SumBucketS](key, value)
}

func WithMetaMap(meta map[string]any) SumBucketOption {
	return aggregations.WithMetaMap[*SumBucketS](meta)
}
