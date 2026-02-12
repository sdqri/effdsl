package autodatehistogram

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type AutoDateHistogramBody struct {
	Field           string `json:"field,omitempty"`
	Buckets         int    `json:"buckets,omitempty"`
	Format          string `json:"format,omitempty"`
	Missing         any    `json:"missing,omitempty"`
	TimeZone        string `json:"time_zone,omitempty"`
	MinimumInterval string `json:"minimum_interval,omitempty"`
}

type AutoDateHistogramS struct {
	name string
	body AutoDateHistogramBody
	aggregations.BaseAggregation
}

func (histogram *AutoDateHistogramS) AggregationName() string {
	return histogram.name
}

func (histogram *AutoDateHistogramS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("auto_date_histogram", histogram.body, histogram.Extras())
}

type AutoDateHistogramOption = aggregations.Option[*AutoDateHistogramS]

func AutoDateHistogram(name, field string, buckets int, opts ...AutoDateHistogramOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &AutoDateHistogramS{
		name: name,
		body: AutoDateHistogramBody{
			Field:   field,
			Buckets: buckets,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithFormat(format string) AutoDateHistogramOption {
	return func(a *AutoDateHistogramS) error {
		a.body.Format = format
		return nil
	}
}

func WithMissing(missing any) AutoDateHistogramOption {
	return func(a *AutoDateHistogramS) error {
		a.body.Missing = missing
		return nil
	}
}

func WithTimeZone(timeZone string) AutoDateHistogramOption {
	return func(a *AutoDateHistogramS) error {
		a.body.TimeZone = timeZone
		return nil
	}
}

func WithMinimumInterval(interval string) AutoDateHistogramOption {
	return func(a *AutoDateHistogramS) error {
		a.body.MinimumInterval = interval
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) AutoDateHistogramOption {
	return aggregations.WithSubAggregation[*AutoDateHistogramS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) AutoDateHistogramOption {
	return aggregations.WithNamedSubAggregation[*AutoDateHistogramS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) AutoDateHistogramOption {
	return aggregations.WithSubAggregationsMap[*AutoDateHistogramS](subsMap)
}

func WithMetaField(key string, value any) AutoDateHistogramOption {
	return aggregations.WithMetaField[*AutoDateHistogramS](key, value)
}

func WithMetaMap(meta map[string]any) AutoDateHistogramOption {
	return aggregations.WithMetaMap[*AutoDateHistogramS](meta)
}
