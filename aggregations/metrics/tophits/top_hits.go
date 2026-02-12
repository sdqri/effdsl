package tophits

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type TopHitsBody struct {
	From             *int           `json:"from,omitempty"`
	Size             *int           `json:"size,omitempty"`
	Sort             []any          `json:"sort,omitempty"`
	Source           any            `json:"_source,omitempty"`
	StoredFields     []string       `json:"stored_fields,omitempty"`
	DocvalueFields   []any          `json:"docvalue_fields,omitempty"`
	ScriptFields     map[string]any `json:"script_fields,omitempty"`
	Highlight        map[string]any `json:"highlight,omitempty"`
	Explain          *bool          `json:"explain,omitempty"`
	TrackScores      *bool          `json:"track_scores,omitempty"`
	Version          *bool          `json:"version,omitempty"`
	SeqNoPrimaryTerm *bool          `json:"seq_no_primary_term,omitempty"`
	Fields           []any          `json:"fields,omitempty"`
}

type TopHitsS struct {
	name string
	body TopHitsBody
	aggregations.BaseAggregation
}

func (hits *TopHitsS) AggregationName() string {
	return hits.name
}

func (hits *TopHitsS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("top_hits", hits.body, hits.Extras())
}

type TopHitsOption = aggregations.Option[*TopHitsS]

func TopHits(name string, opts ...TopHitsOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("metric: aggregation name cannot be empty")}
	}

	agg := &TopHitsS{name: name, body: TopHitsBody{}}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFrom(from int) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.From = &from
		return nil
	}
}

func WithSize(size int) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Size = &size
		return nil
	}
}

func WithSort(sort []any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Sort = sort
		return nil
	}
}

func WithSource(source any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Source = source
		return nil
	}
}

func WithStoredFields(fields []string) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.StoredFields = fields
		return nil
	}
}

func WithDocvalueFields(fields []any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.DocvalueFields = fields
		return nil
	}
}

func WithScriptFields(fields map[string]any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.ScriptFields = fields
		return nil
	}
}

func WithHighlight(highlight map[string]any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Highlight = highlight
		return nil
	}
}

func WithExplain(explain bool) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Explain = &explain
		return nil
	}
}

func WithTrackScores(track bool) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.TrackScores = &track
		return nil
	}
}

func WithVersion(version bool) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Version = &version
		return nil
	}
}

func WithSeqNoPrimaryTerm(enabled bool) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.SeqNoPrimaryTerm = &enabled
		return nil
	}
}

func WithFields(fields []any) TopHitsOption {
	return func(t *TopHitsS) error {
		t.body.Fields = fields
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) TopHitsOption {
	return aggregations.WithSubAggregation[*TopHitsS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) TopHitsOption {
	return aggregations.WithNamedSubAggregation[*TopHitsS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) TopHitsOption {
	return aggregations.WithSubAggregationsMap[*TopHitsS](subsMap)
}

func WithMetaField(key string, value any) TopHitsOption {
	return aggregations.WithMetaField[*TopHitsS](key, value)
}

func WithMetaMap(meta map[string]any) TopHitsOption {
	return aggregations.WithMetaMap[*TopHitsS](meta)
}
