package datehistogram

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type DateHistogramBody struct {
	Field            string         `json:"field,omitempty"`
	CalendarInterval string         `json:"calendar_interval,omitempty"`
	FixedInterval    string         `json:"fixed_interval,omitempty"`
	Format           string         `json:"format,omitempty"`
	TimeZone         string         `json:"time_zone,omitempty"`
	Offset           string         `json:"offset,omitempty"`
	Order            map[string]any `json:"order,omitempty"`
	Keyed            *bool          `json:"keyed,omitempty"`
	MinDocCount      *int           `json:"min_doc_count,omitempty"`
	ExtendedBounds   map[string]any `json:"extended_bounds,omitempty"`
	HardBounds       map[string]any `json:"hard_bounds,omitempty"`
	Missing          any            `json:"missing,omitempty"`
}

type DateHistogramS struct {
	name string
	body DateHistogramBody
	aggregations.BaseAggregation
}

func (histogram *DateHistogramS) AggregationName() string {
	return histogram.name
}

func (histogram *DateHistogramS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("date_histogram", histogram.body, histogram.Extras())
}

type DateHistogramOption = aggregations.Option[*DateHistogramS]

func DateHistogram(name, field string, opts ...DateHistogramOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &DateHistogramS{
		name: name,
		body: DateHistogramBody{
			Field: field,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithCalendarInterval(interval string) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.CalendarInterval = interval
		return nil
	}
}

func WithFixedInterval(interval string) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.FixedInterval = interval
		return nil
	}
}

func WithFormat(format string) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.Format = format
		return nil
	}
}

func WithTimeZone(timeZone string) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.TimeZone = timeZone
		return nil
	}
}

func WithOffset(offset string) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.Offset = offset
		return nil
	}
}

func WithOrder(order map[string]any) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.Order = order
		return nil
	}
}

func WithKeyed(keyed bool) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.Keyed = &keyed
		return nil
	}
}

func WithMinDocCount(count int) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.MinDocCount = &count
		return nil
	}
}

func WithExtendedBounds(bounds map[string]any) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.ExtendedBounds = bounds
		return nil
	}
}

func WithHardBounds(bounds map[string]any) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.HardBounds = bounds
		return nil
	}
}

func WithMissing(missing any) DateHistogramOption {
	return func(d *DateHistogramS) error {
		d.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) DateHistogramOption {
	return aggregations.WithSubAggregation[*DateHistogramS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) DateHistogramOption {
	return aggregations.WithNamedSubAggregation[*DateHistogramS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) DateHistogramOption {
	return aggregations.WithSubAggregationsMap[*DateHistogramS](subsMap)
}

func WithMetaField(key string, value any) DateHistogramOption {
	return aggregations.WithMetaField[*DateHistogramS](key, value)
}

func WithMetaMap(meta map[string]any) DateHistogramOption {
	return aggregations.WithMetaMap[*DateHistogramS](meta)
}
