package multiterms

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MultiTermsBody struct {
	Terms               any            `json:"terms,omitempty"`
	Size                *int           `json:"size,omitempty"`
	ShardSize           *int           `json:"shard_size,omitempty"`
	Order               map[string]any `json:"order,omitempty"`
	MinDocCount         *int           `json:"min_doc_count,omitempty"`
	ShardMinDocCount    *int           `json:"shard_min_doc_count,omitempty"`
	CollectMode         string         `json:"collect_mode,omitempty"`
	ShowTermDocCountErr *bool          `json:"show_term_doc_count_error,omitempty"`
}

type MultiTermsS struct {
	name string
	body MultiTermsBody
	aggregations.BaseAggregation
}

func (terms *MultiTermsS) AggregationName() string {
	return terms.name
}

func (terms *MultiTermsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("multi_terms", terms.body, terms.Extras())
}

type MultiTermsOption = aggregations.Option[*MultiTermsS]

func MultiTerms(name string, terms any, opts ...MultiTermsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &MultiTermsS{
		name: name,
		body: MultiTermsBody{
			Terms: terms,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.Size = &size
		return nil
	}
}

func WithShardSize(size int) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.ShardSize = &size
		return nil
	}
}

func WithOrder(order map[string]any) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.Order = order
		return nil
	}
}

func WithMinDocCount(count int) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.MinDocCount = &count
		return nil
	}
}

func WithShardMinDocCount(count int) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.ShardMinDocCount = &count
		return nil
	}
}

func WithCollectMode(mode string) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.CollectMode = mode
		return nil
	}
}

func WithShowTermDocCountError(enabled bool) MultiTermsOption {
	return func(m *MultiTermsS) error {
		m.body.ShowTermDocCountErr = &enabled
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MultiTermsOption {
	return aggregations.WithSubAggregation[*MultiTermsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MultiTermsOption {
	return aggregations.WithNamedSubAggregation[*MultiTermsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MultiTermsOption {
	return aggregations.WithSubAggregationsMap[*MultiTermsS](subsMap)
}

func WithMetaField(key string, value any) MultiTermsOption {
	return aggregations.WithMetaField[*MultiTermsS](key, value)
}

func WithMetaMap(meta map[string]any) MultiTermsOption {
	return aggregations.WithMetaMap[*MultiTermsS](meta)
}
