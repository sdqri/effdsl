package adjacencymatrix

import (
	"fmt"

	"github.com/sdqri/effdsl/v2/aggregations"
)

type AdjacencyMatrixBody struct {
	Filters   map[string]any `json:"filters,omitempty"`
	Separator string         `json:"separator,omitempty"`
}

type AdjacencyMatrixS struct {
	name string
	body AdjacencyMatrixBody
	aggregations.BaseAggregation
}

func (matrix *AdjacencyMatrixS) AggregationName() string {
	return matrix.name
}

func (matrix *AdjacencyMatrixS) MarshalJSON() ([]byte, error) {
	return aggregations.MarshalAggregation("adjacency_matrix", matrix.body, matrix.Extras())
}

type AdjacencyMatrixOption = aggregations.Option[*AdjacencyMatrixS]

func AdjacencyMatrix(name string, filters map[string]any, opts ...AdjacencyMatrixOption) aggregations.AggregationResult {
	if name == "" {
		return aggregations.AggregationResult{Ok: nil, Err: fmt.Errorf("bucket: aggregation name cannot be empty")}
	}

	agg := &AdjacencyMatrixS{
		name: name,
		body: AdjacencyMatrixBody{
			Filters: filters,
		},
	}

	if err := aggregations.ApplyOptions(agg, opts...); err != nil {
		return aggregations.AggregationResult{Ok: nil, Err: err}
	}

	return aggregations.AggregationResult{Ok: agg, Err: nil}
}

func WithSeparator(separator string) AdjacencyMatrixOption {
	return func(a *AdjacencyMatrixS) error {
		a.body.Separator = separator
		return nil
	}
}

func WithSubAggregation(sub aggregations.AggregationResult) AdjacencyMatrixOption {
	return aggregations.WithSubAggregation[*AdjacencyMatrixS](sub)
}

func WithNamedSubAggregation(name string, sub aggregations.AggregationResult) AdjacencyMatrixOption {
	return aggregations.WithNamedSubAggregation[*AdjacencyMatrixS](name, sub)
}

func WithSubAggregationsMap(subsMap map[string]aggregations.AggregationResult) AdjacencyMatrixOption {
	return aggregations.WithSubAggregationsMap[*AdjacencyMatrixS](subsMap)
}

func WithMetaField(key string, value any) AdjacencyMatrixOption {
	return aggregations.WithMetaField[*AdjacencyMatrixS](key, value)
}

func WithMetaMap(meta map[string]any) AdjacencyMatrixOption {
	return aggregations.WithMetaMap[*AdjacencyMatrixS](meta)
}
