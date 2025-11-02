package aggregations

import "encoding/json"

type AggregationType interface {
	AggregationName() string
	json.Marshaler
	AggregationData() *AggregationData
}

type AggregationResult struct {
	Ok  AggregationType
	Err error
}

type Script struct {
	Lang   string         `json:"lang,omitempty"`
	Source string         `json:"source,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}

type AggregationData struct {
	SubAggregations map[string]AggregationType `json:"aggs,omitempty"`
	Meta            map[string]any             `json:"meta,omitempty"`
}

func (d *AggregationData) AddSubAggregation(agg AggregationType) {
	if agg == nil {
		return
	}
	d.AddNamedSubAggregation(agg.AggregationName(), agg)
}

func (d *AggregationData) AddNamedSubAggregation(name string, agg AggregationType) {
	if agg == nil || name == "" {
		return
	}
	if d.SubAggregations == nil {
		d.SubAggregations = make(map[string]AggregationType)
	}
	d.SubAggregations[name] = agg
}

func (d *AggregationData) SetSubAggregations(aggs map[string]AggregationType) {
	if aggs == nil {
		d.SubAggregations = nil
		return
	}
	d.SubAggregations = make(map[string]AggregationType, len(aggs))
	for name, agg := range aggs {
		if agg == nil || name == "" {
			continue
		}
		d.SubAggregations[name] = agg
	}
}

func (d *AggregationData) SetMeta(meta map[string]any) {
	if meta == nil {
		d.Meta = nil
		return
	}
	if d.Meta == nil {
		d.Meta = make(map[string]any, len(meta))
	} else {
		for k := range d.Meta {
			delete(d.Meta, k)
		}
	}
	for k, v := range meta {
		d.Meta[k] = v
	}
}

func (d *AggregationData) AddMetaField(key string, value any) {
	if key == "" {
		return
	}
	if d.Meta == nil {
		d.Meta = make(map[string]any)
	}
	d.Meta[key] = value
}

type BaseAggregation struct {
	data AggregationData
}

func (b *BaseAggregation) AggregationData() *AggregationData {
	return &b.data
}

func (b *BaseAggregation) AddSubAggregation(agg AggregationType) {
	b.data.AddSubAggregation(agg)
}

func (b *BaseAggregation) AddNamedSubAggregation(name string, agg AggregationType) {
	b.data.AddNamedSubAggregation(name, agg)
}

func (b *BaseAggregation) SetSubAggregations(aggs map[string]AggregationType) {
	b.data.SetSubAggregations(aggs)
}

func (b *BaseAggregation) SetMeta(meta map[string]any) {
	b.data.SetMeta(meta)
}

func (b *BaseAggregation) AddMetaField(key string, value any) {
	b.data.AddMetaField(key, value)
}
