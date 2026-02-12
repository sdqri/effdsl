package inferencebucket

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type InferenceBucketBody struct {
	ModelID         string            `json:"model_id,omitempty"`
	InferenceConfig any               `json:"inference_config,omitempty"`
	BucketsPath     map[string]string `json:"buckets_path,omitempty"`
}

type InferenceBucketS struct {
	name string
	body InferenceBucketBody
	aggregations.BaseAggregation
}

func (inf *InferenceBucketS) AggregationName() string {
	return inf.name
}

func (inf *InferenceBucketS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("inference", inf.body, inf.Extras())
}

type InferenceBucketOption = aggregations.Option[*InferenceBucketS]

func InferenceBucket(name, modelID string, bucketsPath map[string]string, opts ...InferenceBucketOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if modelID == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: inference aggregation requires model_id")}
	}

	if len(bucketsPath) == 0 {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: buckets path cannot be empty")}
	}

	agg := &InferenceBucketS{
		name: name,
		body: InferenceBucketBody{
			ModelID:     modelID,
			BucketsPath: bucketsPath,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithInferenceConfig(config any) InferenceBucketOption {
	return func(i *InferenceBucketS) error {
		i.body.InferenceConfig = config
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) InferenceBucketOption {
	return aggregations.WithSubAggregation[*InferenceBucketS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) InferenceBucketOption {
	return aggregations.WithNamedSubAggregation[*InferenceBucketS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) InferenceBucketOption {
	return aggregations.WithSubAggregationsMap[*InferenceBucketS](subsMap)
}

func WithMetaField(key string, value any) InferenceBucketOption {
	return aggregations.WithMetaField[*InferenceBucketS](key, value)
}

func WithMetaMap(meta map[string]any) InferenceBucketOption {
	return aggregations.WithMetaMap[*InferenceBucketS](meta)
}
