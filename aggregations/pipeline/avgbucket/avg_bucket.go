package avgbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type AvgBucketBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type AvgBucketS struct {
	name string
	body AvgBucketBody
	aggregations.BaseAggregation
}

func (avg *AvgBucketS) AggregationName() string {
	return avg.name
}

func (avg *AvgBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("avg_bucket", avg.body, avg.Extras())
}

type AvgBucketOption = aggregations.Option[*AvgBucketS]

func AvgBucket(name, bucketsPath string, opts ...AvgBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &AvgBucketS{
		name: name,
		body: AvgBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) AvgBucketOption {
	return func(a *AvgBucketS) error {
		a.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) AvgBucketOption {
	return func(a *AvgBucketS) error {
		a.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) AvgBucketOption {
	return aggregations.WithSubAggregation[*AvgBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) AvgBucketOption {
	return aggregations.WithNamedSubAggregation[*AvgBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) AvgBucketOption {
	return aggregations.WithSubAggregationsMap[*AvgBucketS](subsMap)
}

func WithMetaField(key string, value any) AvgBucketOption {
	return aggregations.WithMetaField[*AvgBucketS](key, value)
}

func WithMetaMap(meta map[string]any) AvgBucketOption {
	return aggregations.WithMetaMap[*AvgBucketS](meta)
}
