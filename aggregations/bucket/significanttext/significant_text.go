package significanttext

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SignificantTextBody struct {
	Field               string   `json:"field,omitempty"`
	FilterDuplicateText *bool    `json:"filter_duplicate_text,omitempty"`
	MinDocCount         *int     `json:"min_doc_count,omitempty"`
	ShardMinDocCount    *int     `json:"shard_min_doc_count,omitempty"`
	SourceFields        []string `json:"source_fields,omitempty"`
	BackgroundFilter    any      `json:"background_filter,omitempty"`
	Include             any      `json:"include,omitempty"`
	Exclude             any      `json:"exclude,omitempty"`
}

type SignificantTextS struct {
	name string
	body SignificantTextBody
	aggregations.BaseAggregation
}

func (text *SignificantTextS) AggregationName() string {
	return text.name
}

func (text *SignificantTextS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("significant_text", text.body, text.Extras())
}

type SignificantTextOption = aggregations.Option[*SignificantTextS]

func SignificantText(name, field string, opts ...SignificantTextOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &SignificantTextS{
		name: name,
		body: SignificantTextBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFilterDuplicateText(enabled bool) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.FilterDuplicateText = &enabled
		return nil
	}
}

func WithMinDocCount(count int) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.MinDocCount = &count
		return nil
	}
}

func WithShardMinDocCount(count int) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.ShardMinDocCount = &count
		return nil
	}
}

func WithSourceFields(fields []string) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.SourceFields = fields
		return nil
	}
}

func WithBackgroundFilter(filter any) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.BackgroundFilter = filter
		return nil
	}
}

func WithInclude(include any) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.Include = include
		return nil
	}
}

func WithExclude(exclude any) SignificantTextOption {
	return func(s *SignificantTextS) error {
		s.body.Exclude = exclude
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SignificantTextOption {
	return aggregations.WithSubAggregation[*SignificantTextS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SignificantTextOption {
	return aggregations.WithNamedSubAggregation[*SignificantTextS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SignificantTextOption {
	return aggregations.WithSubAggregationsMap[*SignificantTextS](subsMap)
}

func WithMetaField(key string, value any) SignificantTextOption {
	return aggregations.WithMetaField[*SignificantTextS](key, value)
}

func WithMetaMap(meta map[string]any) SignificantTextOption {
	return aggregations.WithMetaMap[*SignificantTextS](meta)
}
