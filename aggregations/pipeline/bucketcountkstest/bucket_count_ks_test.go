package bucketcountkstest

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type BucketCountKSTestBody struct {
	BucketsPath    string    `json:"buckets_path,omitempty"`
	Alternative    []string  `json:"alternative,omitempty"`
	Fractions      []float64 `json:"fractions,omitempty"`
	SamplingMethod string    `json:"sampling_method,omitempty"`
}

type BucketCountKSTestS struct {
	name string
	body BucketCountKSTestBody
	aggregations.BaseAggregation
}

func (ks *BucketCountKSTestS) AggregationName() string {
	return ks.name
}

func (ks *BucketCountKSTestS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("bucket_count_ks_test", ks.body, ks.Extras())
}

type BucketCountKSTestOption = aggregations.Option[*BucketCountKSTestS]

func BucketCountKSTest(name, bucketsPath string, opts ...BucketCountKSTestOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if bucketsPath == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	agg := &BucketCountKSTestS{
		name: name,
		body: BucketCountKSTestBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithAlternatives(alternatives []string) BucketCountKSTestOption {
	return func(k *BucketCountKSTestS) error {
		k.body.Alternative = alternatives
		return nil
	}
}

func WithFractions(fractions []float64) BucketCountKSTestOption {
	return func(k *BucketCountKSTestS) error {
		k.body.Fractions = fractions
		return nil
	}
}

func WithSamplingMethod(method string) BucketCountKSTestOption {
	return func(k *BucketCountKSTestS) error {
		k.body.SamplingMethod = method
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BucketCountKSTestOption {
	return aggregations.WithSubAggregation[*BucketCountKSTestS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BucketCountKSTestOption {
	return aggregations.WithNamedSubAggregation[*BucketCountKSTestS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BucketCountKSTestOption {
	return aggregations.WithSubAggregationsMap[*BucketCountKSTestS](subsMap)
}

func WithMetaField(key string, value any) BucketCountKSTestOption {
	return aggregations.WithMetaField[*BucketCountKSTestS](key, value)
}

func WithMetaMap(meta map[string]any) BucketCountKSTestOption {
	return aggregations.WithMetaMap[*BucketCountKSTestS](meta)
}
