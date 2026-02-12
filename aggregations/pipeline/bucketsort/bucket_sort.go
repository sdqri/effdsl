package bucketsort

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type BucketSortBody struct {
	Sort      []any  `json:"sort,omitempty"`
	From      *int   `json:"from,omitempty"`
	Size      *int   `json:"size,omitempty"`
	GapPolicy string `json:"gap_policy,omitempty"`
}

type BucketSortS struct {
	name string
	body BucketSortBody
	aggregations.BaseAggregation
}

func (bucket *BucketSortS) AggregationName() string {
	return bucket.name
}

func (bucket *BucketSortS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("bucket_sort", bucket.body, bucket.Extras())
}

type BucketSortOption = aggregations.Option[*BucketSortS]

func BucketSort(name string, opts ...BucketSortOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &BucketSortS{
		name: name,
		body: BucketSortBody{},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSort(sort []any) BucketSortOption {
	return func(b *BucketSortS) error {
		b.body.Sort = sort
		return nil
	}
}

func WithFrom(from int) BucketSortOption {
	return func(b *BucketSortS) error {
		b.body.From = &from
		return nil
	}
}

func WithSize(size int) BucketSortOption {
	return func(b *BucketSortS) error {
		b.body.Size = &size
		return nil
	}
}

func WithGapPolicy(policy string) BucketSortOption {
	return func(b *BucketSortS) error {
		b.body.GapPolicy = policy
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BucketSortOption {
	return aggregations.WithSubAggregation[*BucketSortS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BucketSortOption {
	return aggregations.WithNamedSubAggregation[*BucketSortS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BucketSortOption {
	return aggregations.WithSubAggregationsMap[*BucketSortS](subsMap)
}

func WithMetaField(key string, value any) BucketSortOption {
	return aggregations.WithMetaField[*BucketSortS](key, value)
}

func WithMetaMap(meta map[string]any) BucketSortOption {
	return aggregations.WithMetaMap[*BucketSortS](meta)
}
