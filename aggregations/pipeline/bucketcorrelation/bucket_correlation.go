package bucketcorrelation

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type CountCorrelationIndicator struct {
	Expectations []float64 `json:"expectations,omitempty"`
	DocCount     *int      `json:"doc_count,omitempty"`
	Fractions    []float64 `json:"fractions,omitempty"`
}

type CountCorrelationFunction struct {
	Indicator *CountCorrelationIndicator `json:"indicator,omitempty"`
}

type BucketCorrelationFunction struct {
	CountCorrelation *CountCorrelationFunction `json:"count_correlation,omitempty"`
}

type BucketCorrelationBody struct {
	BucketsPath string                    `json:"buckets_path,omitempty"`
	Function    BucketCorrelationFunction `json:"function,omitempty"`
}

type BucketCorrelationS struct {
	name string
	body BucketCorrelationBody
	aggregations.BaseAggregation
}

func (corr *BucketCorrelationS) AggregationName() string {
	return corr.name
}

func (corr *BucketCorrelationS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("bucket_correlation", corr.body, corr.Extras())
}

type BucketCorrelationOption = aggregations.Option[*BucketCorrelationS]

func BucketCorrelation(name, bucketsPath string, function BucketCorrelationFunction, opts ...BucketCorrelationOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if bucketsPath == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	if function.CountCorrelation == nil || function.CountCorrelation.Indicator == nil {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: bucket_correlation requires count_correlation indicator")}
	}

	agg := &BucketCorrelationS{
		name: name,
		body: BucketCorrelationBody{
			BucketsPath: bucketsPath,
			Function:    function,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func CountCorrelation(expectations []float64, docCount int, fractions []float64) BucketCorrelationFunction {
	indicator := &CountCorrelationIndicator{
		Expectations: expectations,
		DocCount:     &docCount,
		Fractions:    fractions,
	}

	return BucketCorrelationFunction{
		CountCorrelation: &CountCorrelationFunction{
			Indicator: indicator,
		},
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) BucketCorrelationOption {
	return aggregations.WithSubAggregation[*BucketCorrelationS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) BucketCorrelationOption {
	return aggregations.WithNamedSubAggregation[*BucketCorrelationS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) BucketCorrelationOption {
	return aggregations.WithSubAggregationsMap[*BucketCorrelationS](subsMap)
}

func WithMetaField(key string, value any) BucketCorrelationOption {
	return aggregations.WithMetaField[*BucketCorrelationS](key, value)
}

func WithMetaMap(meta map[string]any) BucketCorrelationOption {
	return aggregations.WithMetaMap[*BucketCorrelationS](meta)
}
