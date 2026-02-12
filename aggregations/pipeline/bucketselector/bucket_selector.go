package bucketselector

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type BucketSelectorBody struct {
	BucketsPath map[string]string   `json:"buckets_path,omitempty"`
	Script      aggregations.Script `json:"script,omitempty"`
	GapPolicy   string              `json:"gap_policy,omitempty"`
}

type BucketSelectorS struct {
	name string
	body BucketSelectorBody
	aggregations.BaseAggregation
}

func (bucket *BucketSelectorS) AggregationName() string {
	return bucket.name
}

func (bucket *BucketSelectorS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("bucket_selector", bucket.body, bucket.Extras())
}

type BucketSelectorOption = aggregations.Option[*BucketSelectorS]

func BucketSelector(name string, bucketsPath map[string]string, script aggregations.Script, opts ...BucketSelectorOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if len(bucketsPath) == 0 {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	agg := &BucketSelectorS{
		name: name,
		body: BucketSelectorBody{
			BucketsPath: bucketsPath,
			Script:      script,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) BucketSelectorOption {
	return func(b *BucketSelectorS) error {
		b.body.GapPolicy = policy
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BucketSelectorOption {
	return aggregations.WithSubAggregation[*BucketSelectorS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BucketSelectorOption {
	return aggregations.WithNamedSubAggregation[*BucketSelectorS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BucketSelectorOption {
	return aggregations.WithSubAggregationsMap[*BucketSelectorS](subsMap)
}

func WithMetaField(key string, value any) BucketSelectorOption {
	return aggregations.WithMetaField[*BucketSelectorS](key, value)
}

func WithMetaMap(meta map[string]any) BucketSelectorOption {
	return aggregations.WithMetaMap[*BucketSelectorS](meta)
}
