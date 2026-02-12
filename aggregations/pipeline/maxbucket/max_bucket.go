package maxbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MaxBucketBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type MaxBucketS struct {
	name string
	body MaxBucketBody
	aggregations.BaseAggregation
}

func (bucket *MaxBucketS) AggregationName() string {
	return bucket.name
}

func (bucket *MaxBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("max_bucket", bucket.body, bucket.Extras())
}

type MaxBucketOption = aggregations.Option[*MaxBucketS]

func MaxBucket(name, bucketsPath string, opts ...MaxBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &MaxBucketS{
		name: name,
		body: MaxBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) MaxBucketOption {
	return func(m *MaxBucketS) error {
		m.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) MaxBucketOption {
	return func(m *MaxBucketS) error {
		m.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MaxBucketOption {
	return aggregations.WithSubAggregation[*MaxBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MaxBucketOption {
	return aggregations.WithNamedSubAggregation[*MaxBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MaxBucketOption {
	return aggregations.WithSubAggregationsMap[*MaxBucketS](subsMap)
}

func WithMetaField(key string, value any) MaxBucketOption {
	return aggregations.WithMetaField[*MaxBucketS](key, value)
}

func WithMetaMap(meta map[string]any) MaxBucketOption {
	return aggregations.WithMetaMap[*MaxBucketS](meta)
}
