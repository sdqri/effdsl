package rareterms

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type RareTermsBody struct {
	Field       string `json:"field,omitempty"`
	MaxDocCount *int   `json:"max_doc_count,omitempty"`
	Include     any    `json:"include,omitempty"`
	Exclude     any    `json:"exclude,omitempty"`
	Missing     any    `json:"missing,omitempty"`
}

type RareTermsS struct {
	name string
	body RareTermsBody
	aggregations.BaseAggregation
}

func (rare *RareTermsS) AggregationName() string {
	return rare.name
}

func (rare *RareTermsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("rare_terms", rare.body, rare.Extras())
}

type RareTermsOption = aggregations.Option[*RareTermsS]

func RareTerms(name, field string, opts ...RareTermsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &RareTermsS{
		name: name,
		body: RareTermsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithMaxDocCount(count int) RareTermsOption {
	return func(r *RareTermsS) error {
		r.body.MaxDocCount = &count
		return nil
	}
}

func WithInclude(include any) RareTermsOption {
	return func(r *RareTermsS) error {
		r.body.Include = include
		return nil
	}
}

func WithExclude(exclude any) RareTermsOption {
	return func(r *RareTermsS) error {
		r.body.Exclude = exclude
		return nil
	}
}

func WithMissing(missing any) RareTermsOption {
	return func(r *RareTermsS) error {
		r.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) RareTermsOption {
	return aggregations.WithSubAggregation[*RareTermsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) RareTermsOption {
	return aggregations.WithNamedSubAggregation[*RareTermsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) RareTermsOption {
	return aggregations.WithSubAggregationsMap[*RareTermsS](subsMap)
}

func WithMetaField(key string, value any) RareTermsOption {
	return aggregations.WithMetaField[*RareTermsS](key, value)
}

func WithMetaMap(meta map[string]any) RareTermsOption {
	return aggregations.WithMetaMap[*RareTermsS](meta)
}
