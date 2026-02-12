package daterange

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type DateRangeItem struct {
	Key  string `json:"key,omitempty"`
	From any    `json:"from,omitempty"`
	To   any    `json:"to,omitempty"`
}

type DateRangeBody struct {
	Field    string          `json:"field,omitempty"`
	Ranges   []DateRangeItem `json:"ranges,omitempty"`
	Format   string          `json:"format,omitempty"`
	TimeZone string          `json:"time_zone,omitempty"`
	Missing  any             `json:"missing,omitempty"`
}

type DateRangeS struct {
	name string
	body DateRangeBody
	aggregations.BaseAggregation
}

func (rangeAgg *DateRangeS) AggregationName() string {
	return rangeAgg.name
}

func (rangeAgg *DateRangeS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("date_range", rangeAgg.body, rangeAgg.Extras())
}

type DateRangeOption = aggregations.Option[*DateRangeS]

func DateRange(name, field string, ranges []DateRangeItem, opts ...DateRangeOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &DateRangeS{
		name: name,
		body: DateRangeBody{
			Field:  field,
			Ranges: ranges,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) DateRangeOption {
	return func(d *DateRangeS) error {
		d.body.Format = format
		return nil
	}
}

func WithTimeZone(timeZone string) DateRangeOption {
	return func(d *DateRangeS) error {
		d.body.TimeZone = timeZone
		return nil
	}
}

func WithMissing(missing any) DateRangeOption {
	return func(d *DateRangeS) error {
		d.body.Missing = missing
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) DateRangeOption {
	return aggregations.WithSubAggregation[*DateRangeS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) DateRangeOption {
	return aggregations.WithNamedSubAggregation[*DateRangeS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) DateRangeOption {
	return aggregations.WithSubAggregationsMap[*DateRangeS](subsMap)
}

func WithMetaField(key string, value any) DateRangeOption {
	return aggregations.WithMetaField[*DateRangeS](key, value)
}

func WithMetaMap(meta map[string]any) DateRangeOption {
	return aggregations.WithMetaMap[*DateRangeS](meta)
}
