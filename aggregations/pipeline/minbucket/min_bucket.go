package minbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MinBucketBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type MinBucketS struct {
	name string
	body MinBucketBody
	aggregations.BaseAggregation
}

func (bucket *MinBucketS) AggregationName() string {
	return bucket.name
}

func (bucket *MinBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("min_bucket", bucket.body, bucket.Extras())
}

type MinBucketOption = aggregations.Option[*MinBucketS]

func MinBucket(name, bucketsPath string, opts ...MinBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &MinBucketS{
		name: name,
		body: MinBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) MinBucketOption {
	return func(m *MinBucketS) error {
		m.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) MinBucketOption {
	return func(m *MinBucketS) error {
		m.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MinBucketOption {
	return aggregations.WithSubAggregation[*MinBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MinBucketOption {
	return aggregations.WithNamedSubAggregation[*MinBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MinBucketOption {
	return aggregations.WithSubAggregationsMap[*MinBucketS](subsMap)
}

func WithMetaField(key string, value any) MinBucketOption {
	return aggregations.WithMetaField[*MinBucketS](key, value)
}

func WithMetaMap(meta map[string]any) MinBucketOption {
	return aggregations.WithMetaMap[*MinBucketS](meta)
}
