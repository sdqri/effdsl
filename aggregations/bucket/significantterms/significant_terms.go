package significantterms

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type SignificantTermsBody struct {
	Field            string
	Size             *int
	MinDocCount      *int
	ShardMinDocCount *int
	Include          any
	Exclude          any
	ExecutionHint    string
	BackgroundFilter any
}

type SignificantTermsS struct {
	name            string
	body            SignificantTermsBody
	heuristicName   string
	heuristicParams map[string]any
	aggregations.BaseAggregation
}

func (terms *SignificantTermsS) AggregationName() string {
	return terms.name
}

func (terms *SignificantTermsS) MarshalJSON() ([]byte, error) {
	body := map[string]any{}
	if terms.body.Field != "" {
		body["field"] = terms.body.Field
	}
	if terms.body.Size != nil {
		body["size"] = *terms.body.Size
	}
	if terms.body.MinDocCount != nil {
		body["min_doc_count"] = *terms.body.MinDocCount
	}
	if terms.body.ShardMinDocCount != nil {
		body["shard_min_doc_count"] = *terms.body.ShardMinDocCount
	}
	if terms.body.Include != nil {
		body["include"] = terms.body.Include
	}
	if terms.body.Exclude != nil {
		body["exclude"] = terms.body.Exclude
	}
	if terms.body.ExecutionHint != "" {
		body["execution_hint"] = terms.body.ExecutionHint
	}
	if terms.body.BackgroundFilter != nil {
		body["background_filter"] = terms.body.BackgroundFilter
	}
	if terms.heuristicName != "" {
		body[terms.heuristicName] = terms.heuristicParams
	}

	return aggregations.MarshalAggregation("significant_terms", body, terms.Extras())
}

type SignificantTermsOption = aggregations.Option[*SignificantTermsS]

func SignificantTerms(name, field string, opts ...SignificantTermsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &SignificantTermsS{
		name: name,
		body: SignificantTermsBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSize(size int) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.Size = &size
		return nil
	}
}

func WithMinDocCount(count int) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.MinDocCount = &count
		return nil
	}
}

func WithShardMinDocCount(count int) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.ShardMinDocCount = &count
		return nil
	}
}

func WithInclude(include any) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.Include = include
		return nil
	}
}

func WithExclude(exclude any) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.Exclude = exclude
		return nil
	}
}

func WithExecutionHint(hint string) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.ExecutionHint = hint
		return nil
	}
}

func WithBackgroundFilter(filter any) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.body.BackgroundFilter = filter
		return nil
	}
}

func WithHeuristic(name string, params map[string]any) SignificantTermsOption {
	return func(s *SignificantTermsS) error {
		s.heuristicName = name
		s.heuristicParams = params
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) SignificantTermsOption {
	return aggregations.WithSubAggregation[*SignificantTermsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) SignificantTermsOption {
	return aggregations.WithNamedSubAggregation[*SignificantTermsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) SignificantTermsOption {
	return aggregations.WithSubAggregationsMap[*SignificantTermsS](subsMap)
}

func WithMetaField(key string, value any) SignificantTermsOption {
	return aggregations.WithMetaField[*SignificantTermsS](key, value)
}

func WithMetaMap(meta map[string]any) SignificantTermsOption {
	return aggregations.WithMetaMap[*SignificantTermsS](meta)
}
