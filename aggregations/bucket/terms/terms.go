package terms

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type TermsBody struct {
	Field               string               `json:"field,omitempty"`
	Script              *aggregations.Script `json:"script,omitempty"`
	Size                *int                 `json:"size,omitempty"`
	ShardSize           *int                 `json:"shard_size,omitempty"`
	Order               map[string]any       `json:"order,omitempty"`
	MinDocCount         *int                 `json:"min_doc_count,omitempty"`
	ShardMinDocCount    *int                 `json:"shard_min_doc_count,omitempty"`
	Include             any                  `json:"include,omitempty"`
	Exclude             any                  `json:"exclude,omitempty"`
	Missing             any                  `json:"missing,omitempty"`
	ExecutionHint       string               `json:"execution_hint,omitempty"`
	CollectMode         string               `json:"collect_mode,omitempty"`
	ShowTermDocCountErr *bool                `json:"show_term_doc_count_error,omitempty"`
	ValueType           string               `json:"value_type,omitempty"`
}

type TermsS struct {
	name string
	body TermsBody
	aggregations.BaseAggregation
}

func (terms *TermsS) AggregationName() string {
	return terms.name
}

func (terms *TermsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("terms", terms.body, terms.Extras())
}

type TermsOption = aggregations.Option[*TermsS]

func Terms(name, field string, opts ...TermsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &TermsS{
		name: name,
		body: TermsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithScript(script aggregations.Script) TermsOption {
	return func(t *TermsS) error {
		t.body.Script = &script
		return nil
	}
}

func WithSize(size int) TermsOption {
	return func(t *TermsS) error {
		t.body.Size = &size
		return nil
	}
}

func WithShardSize(size int) TermsOption {
	return func(t *TermsS) error {
		t.body.ShardSize = &size
		return nil
	}
}

func WithOrder(order map[string]any) TermsOption {
	return func(t *TermsS) error {
		t.body.Order = order
		return nil
	}
}

func WithMinDocCount(count int) TermsOption {
	return func(t *TermsS) error {
		t.body.MinDocCount = &count
		return nil
	}
}

func WithShardMinDocCount(count int) TermsOption {
	return func(t *TermsS) error {
		t.body.ShardMinDocCount = &count
		return nil
	}
}

func WithInclude(include any) TermsOption {
	return func(t *TermsS) error {
		t.body.Include = include
		return nil
	}
}

func WithExclude(exclude any) TermsOption {
	return func(t *TermsS) error {
		t.body.Exclude = exclude
		return nil
	}
}

func WithMissing(missing any) TermsOption {
	return func(t *TermsS) error {
		t.body.Missing = missing
		return nil
	}
}

func WithExecutionHint(hint string) TermsOption {
	return func(t *TermsS) error {
		t.body.ExecutionHint = hint
		return nil
	}
}

func WithCollectMode(mode string) TermsOption {
	return func(t *TermsS) error {
		t.body.CollectMode = mode
		return nil
	}
}

func WithShowTermDocCountError(enabled bool) TermsOption {
	return func(t *TermsS) error {
		t.body.ShowTermDocCountErr = &enabled
		return nil
	}
}

func WithValueType(valueType string) TermsOption {
	return func(t *TermsS) error {
		t.body.ValueType = valueType
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) TermsOption {
	return aggregations.WithSubAggregation[*TermsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) TermsOption {
	return aggregations.WithNamedSubAggregation[*TermsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) TermsOption {
	return aggregations.WithSubAggregationsMap[*TermsS](subsMap)
}

func WithMetaField(key string, value any) TermsOption {
	return aggregations.WithMetaField[*TermsS](key, value)
}

func WithMetaMap(meta map[string]any) TermsOption {
	return aggregations.WithMetaMap[*TermsS](meta)
}
