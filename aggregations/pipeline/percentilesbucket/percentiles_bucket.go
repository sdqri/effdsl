package percentilesbucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type PercentilesBucketBody struct {
	BucketsPath string    `json:"buckets_path,omitempty"`
	GapPolicy   string    `json:"gap_policy,omitempty"`
	Format      string    `json:"format,omitempty"`
	Percents    []float64 `json:"percents,omitempty"`
	Keyed       *bool     `json:"keyed,omitempty"`
}

type PercentilesBucketS struct {
	name string
	body PercentilesBucketBody
	aggregations.BaseAggregation
}

func (percentiles *PercentilesBucketS) AggregationName() string {
	return percentiles.name
}

func (percentiles *PercentilesBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("percentiles_bucket", percentiles.body, percentiles.Extras())
}

type PercentilesBucketOption = aggregations.Option[*PercentilesBucketS]

func PercentilesBucket(name, bucketsPath string, opts ...PercentilesBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &PercentilesBucketS{
		name: name,
		body: PercentilesBucketBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) PercentilesBucketOption {
	return func(p *PercentilesBucketS) error {
		p.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) PercentilesBucketOption {
	return func(p *PercentilesBucketS) error {
		p.body.Format = format
		return nil
	}
}

func WithPercents(percents []float64) PercentilesBucketOption {
	return func(p *PercentilesBucketS) error {
		p.body.Percents = percents
		return nil
	}
}

func WithKeyed(keyed bool) PercentilesBucketOption {
	return func(p *PercentilesBucketS) error {
		p.body.Keyed = &keyed
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) PercentilesBucketOption {
	return aggregations.WithSubAggregation[*PercentilesBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) PercentilesBucketOption {
	return aggregations.WithNamedSubAggregation[*PercentilesBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) PercentilesBucketOption {
	return aggregations.WithSubAggregationsMap[*PercentilesBucketS](subsMap)
}

func WithMetaField(key string, value any) PercentilesBucketOption {
	return aggregations.WithMetaField[*PercentilesBucketS](key, value)
}

func WithMetaMap(meta map[string]any) PercentilesBucketOption {
	return aggregations.WithMetaMap[*PercentilesBucketS](meta)
}
