package histogram

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type HistogramBody struct {
	Field          string         `json:"field,omitempty"`
	Interval       float64        `json:"interval,omitempty"`
	Offset         *float64       `json:"offset,omitempty"`
	MinDocCount    *int           `json:"min_doc_count,omitempty"`
	Order          map[string]any `json:"order,omitempty"`
	Keyed          *bool          `json:"keyed,omitempty"`
	ExtendedBounds map[string]any `json:"extended_bounds,omitempty"`
	HardBounds     map[string]any `json:"hard_bounds,omitempty"`
	Missing        any            `json:"missing,omitempty"`
}

type HistogramS struct {
	name string
	body HistogramBody
	aggregations.BaseAggregation
}

func (histogram *HistogramS) AggregationName() string {
	return histogram.name
}

func (histogram *HistogramS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("histogram", histogram.body, histogram.Extras())
}

type HistogramOption = aggregations.Option[*HistogramS]

func Histogram(name, field string, interval float64, opts ...HistogramOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &HistogramS{
		name: name,
		body: HistogramBody{
			Field:    field,
			Interval: interval,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithOffset(offset float64) HistogramOption {
	return func(h *HistogramS) error {
		h.body.Offset = &offset
		return nil
	}
}

func WithMinDocCount(count int) HistogramOption {
	return func(h *HistogramS) error {
		h.body.MinDocCount = &count
		return nil
	}
}

func WithOrder(order map[string]any) HistogramOption {
	return func(h *HistogramS) error {
		h.body.Order = order
		return nil
	}
}

func WithKeyed(keyed bool) HistogramOption {
	return func(h *HistogramS) error {
		h.body.Keyed = &keyed
		return nil
	}
}

func WithExtendedBounds(bounds map[string]any) HistogramOption {
	return func(h *HistogramS) error {
		h.body.ExtendedBounds = bounds
		return nil
	}
}

func WithHardBounds(bounds map[string]any) HistogramOption {
	return func(h *HistogramS) error {
		h.body.HardBounds = bounds
		return nil
	}
}

func WithMissing(missing any) HistogramOption {
	return func(h *HistogramS) error {
		h.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) HistogramOption {
	return aggregations.WithSubAggregation[*HistogramS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) HistogramOption {
	return aggregations.WithNamedSubAggregation[*HistogramS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) HistogramOption {
	return aggregations.WithSubAggregationsMap[*HistogramS](subsMap)
}

func WithMetaField(key string, value any) HistogramOption {
	return aggregations.WithMetaField[*HistogramS](key, value)
}

func WithMetaMap(meta map[string]any) HistogramOption {
	return aggregations.WithMetaMap[*HistogramS](meta)
}
