package extendedstatsbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type ExtendedStatsBucketBody struct {
	BucketsPath string   `json:"buckets_path,omitempty"`
	GapPolicy   string   `json:"gap_policy,omitempty"`
	Format      string   `json:"format,omitempty"`
	Sigma       *float64 `json:"sigma,omitempty"`
}

type ExtendedStatsBucketS struct {
	name string
	body ExtendedStatsBucketBody
	aggregations.BaseAggregation
}

func (stats *ExtendedStatsBucketS) AggregationName() string {
	return stats.name
}

func (stats *ExtendedStatsBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("extended_stats_bucket", stats.body, stats.Extras())
}

type ExtendedStatsBucketOption = aggregations.Option[*ExtendedStatsBucketS]

func ExtendedStatsBucket(name, bucketsPath string, opts ...ExtendedStatsBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &ExtendedStatsBucketS{
		name: name,
		body: ExtendedStatsBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) ExtendedStatsBucketOption {
	return func(s *ExtendedStatsBucketS) error {
		s.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) ExtendedStatsBucketOption {
	return func(s *ExtendedStatsBucketS) error {
		s.body.Format = format
		return nil
	}
}

func WithSigma(sigma float64) ExtendedStatsBucketOption {
	return func(s *ExtendedStatsBucketS) error {
		s.body.Sigma = &sigma
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) ExtendedStatsBucketOption {
	return aggregations.WithSubAggregation[*ExtendedStatsBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) ExtendedStatsBucketOption {
	return aggregations.WithNamedSubAggregation[*ExtendedStatsBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) ExtendedStatsBucketOption {
	return aggregations.WithSubAggregationsMap[*ExtendedStatsBucketS](subsMap)
}

func WithMetaField(key string, value any) ExtendedStatsBucketOption {
	return aggregations.WithMetaField[*ExtendedStatsBucketS](key, value)
}

func WithMetaMap(meta map[string]any) ExtendedStatsBucketOption {
	return aggregations.WithMetaMap[*ExtendedStatsBucketS](meta)
}
