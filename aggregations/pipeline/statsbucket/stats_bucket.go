package statsbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type StatsBucketBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type StatsBucketS struct {
	name string
	body StatsBucketBody
	aggregations.BaseAggregation
}

func (stats *StatsBucketS) AggregationName() string {
	return stats.name
}

func (stats *StatsBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("stats_bucket", stats.body, stats.Extras())
}

type StatsBucketOption = aggregations.Option[*StatsBucketS]

func StatsBucket(name, bucketsPath string, opts ...StatsBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &StatsBucketS{
		name: name,
		body: StatsBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) StatsBucketOption {
	return func(s *StatsBucketS) error {
		s.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) StatsBucketOption {
	return func(s *StatsBucketS) error {
		s.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) StatsBucketOption {
	return aggregations.WithSubAggregation[*StatsBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) StatsBucketOption {
	return aggregations.WithNamedSubAggregation[*StatsBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) StatsBucketOption {
	return aggregations.WithSubAggregationsMap[*StatsBucketS](subsMap)
}

func WithMetaField(key string, value any) StatsBucketOption {
	return aggregations.WithMetaField[*StatsBucketS](key, value)
}

func WithMetaMap(meta map[string]any) StatsBucketOption {
	return aggregations.WithMetaMap[*StatsBucketS](meta)
}
