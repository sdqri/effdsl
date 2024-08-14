package effdsl

import "encoding/json"

type SourceFilterS struct {
	Includes []string `json:"includes,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
}

func (sf SourceFilterS) MarshalJSON() ([]byte, error) {
	type SourceFilterBase SourceFilterS
	return json.Marshal((SourceFilterBase)(sf))
}

type SourceFitlerOption func(*SourceFilterS)

func WithIncludes(fields ...string) SourceFitlerOption {
	return func(sf *SourceFilterS) {
		if sf.Includes == nil {
			sf.Includes = fields
		} else {
			sf.Includes = append(sf.Includes, fields...)
		}
	}
}

func WithExcludes(fields ...string) SourceFitlerOption {
	return func(sf *SourceFilterS) {
		if sf.Excludes == nil {
			sf.Excludes = fields
		} else {
			sf.Excludes = append(sf.Excludes, fields...)
		}
	}
}

func SourceFilter(opts ...SourceFitlerOption) SourceFilterS {
	sourceFilter := SourceFilterS{}
	for _, opt := range opts {
		opt(&sourceFilter)
	}
	return sourceFilter
}
