package movingpercentiles

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type MovingPercentilesBody struct {
	BucketsPath string `json:"buckets_path,omitempty"`
	Window      *int   `json:"window,omitempty"`
	Shift       *int   `json:"shift,omitempty"`
}

type MovingPercentilesS struct {
	name string
	body MovingPercentilesBody
	aggregations.BaseAggregation
}

func (mov *MovingPercentilesS) AggregationName() string {
	return mov.name
}

func (mov *MovingPercentilesS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("moving_percentiles", mov.body, mov.Extras())
}

type MovingPercentilesOption = aggregations.Option[*MovingPercentilesS]

func MovingPercentiles(name, bucketsPath string, window int, opts ...MovingPercentilesOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: aggregation name cannot be empty")}
	}

	if window <= 0 {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("pipeline: moving_percentiles requires window greater than zero")}
	}

	agg := &MovingPercentilesS{
		name: name,
		body: MovingPercentilesBody{
			BucketsPath: bucketsPath,
			Window:      &window,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithShift(shift int) MovingPercentilesOption {
	return func(m *MovingPercentilesS) error {
		m.body.Shift = &shift
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) MovingPercentilesOption {
	return aggregations.WithSubAggregation[*MovingPercentilesS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) MovingPercentilesOption {
	return aggregations.WithNamedSubAggregation[*MovingPercentilesS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) MovingPercentilesOption {
	return aggregations.WithSubAggregationsMap[*MovingPercentilesS](subsMap)
}

func WithMetaField(key string, value any) MovingPercentilesOption {
	return aggregations.WithMetaField[*MovingPercentilesS](key, value)
}

func WithMetaMap(meta map[string]any) MovingPercentilesOption {
	return aggregations.WithMetaMap[*MovingPercentilesS](meta)
}
