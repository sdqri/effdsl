package serialdiff

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SerialDiffBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Lag         *int   `json:"lag,omitempty"`
	GapPolicy   string `json:"gap_policy,omitempty"`
	Format      string `json:"format,omitempty"`
}

type SerialDiffS struct {
	name string
	body SerialDiffBody
	aggregations.BaseAggregation
}

func (diff *SerialDiffS) AggregationName() string {
	return diff.name
}

func (diff *SerialDiffS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("serial_diff", diff.body, diff.Extras())
}

type SerialDiffOption = aggregations.Option[*SerialDiffS]

func SerialDiff(name, bucketsPath string, opts ...SerialDiffOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	agg := &SerialDiffS{
		name: name,
		body: SerialDiffBody{BucketsPath: bucketsPath},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithLag(lag int) SerialDiffOption {
	return func(s *SerialDiffS) error {
		s.body.Lag = &lag
		return nil
	}
}

func WithGapPolicy(policy string) SerialDiffOption {
	return func(s *SerialDiffS) error {
		s.body.GapPolicy = policy
		return nil
	}
}

func WithFormat(format string) SerialDiffOption {
	return func(s *SerialDiffS) error {
		s.body.Format = format
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SerialDiffOption {
	return aggregations.WithSubAggregation[*SerialDiffS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SerialDiffOption {
	return aggregations.WithNamedSubAggregation[*SerialDiffS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SerialDiffOption {
	return aggregations.WithSubAggregationsMap[*SerialDiffS](subsMap)
}

func WithMetaField(key string, value any) SerialDiffOption {
	return aggregations.WithMetaField[*SerialDiffS](key, value)
}

func WithMetaMap(meta map[string]any) SerialDiffOption {
	return aggregations.WithMetaMap[*SerialDiffS](meta)
}
