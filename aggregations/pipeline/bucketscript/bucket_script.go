package bucketscript

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type BucketScriptBody struct {
	BucketsPath map[string]string   `json:"buckets_path,omitempty"`
	Script      aggregations.Script `json:"script,omitempty"`
	GapPolicy   string              `json:"gap_policy,omitempty"`
	Format      string              `json:"format,omitempty"`
}

type BucketScriptS struct {
	name string
	body BucketScriptBody
	aggregations.BaseAggregation
}

func (bucket *BucketScriptS) AggregationName() string {
	return bucket.name
}

func (bucket *BucketScriptS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("bucket_script", bucket.body, bucket.Extras())
}

type BucketScriptOption = aggregations.Option[*BucketScriptS]

func BucketScript(name string, bucketsPath map[string]string, script aggregations.Script, opts ...BucketScriptOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if len(bucketsPath) == 0 {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	agg := &BucketScriptS{
		name: name,
		body: BucketScriptBody{
			BucketsPath: bucketsPath,
			Script:      script,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithGapPolicy(policy string) BucketScriptOption {
	return func(b *BucketScriptS) error {
		b.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) BucketScriptOption {
	return func(b *BucketScriptS) error {
		b.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BucketScriptOption {
	return aggregations.WithSubAggregation[*BucketScriptS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BucketScriptOption {
	return aggregations.WithNamedSubAggregation[*BucketScriptS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BucketScriptOption {
	return aggregations.WithSubAggregationsMap[*BucketScriptS](subsMap)
}

func WithMetaField(key string, value any) BucketScriptOption {
	return aggregations.WithMetaField[*BucketScriptS](key, value)
}

func WithMetaMap(meta map[string]any) BucketScriptOption {
	return aggregations.WithMetaMap[*BucketScriptS](meta)
}
