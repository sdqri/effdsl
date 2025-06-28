package completionsuggester

import (
	"encoding/json"
	"errors"

	"github.com/sdqri/effdsl/v2"
)

// CompletionSuggesterS - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type CompletionSuggesterS struct {
	Name                  string          `json:"-"`                                 // Required
	Prefix                string          `json:"-"`                                 // Required
	Regex                 string          `json:"-"`                                 // Required
	Field                 string          `json:"field"`                             // Required
	Size                  *uint64         `json:"size,omitempty"`                    // The number of suggestions to return (defaults to 5).
	SkipDuplicates        bool            `json:"skip_duplicates,omitempty"`         // Whether duplicate suggestions should be filtered out (defaults to false).
	Flags                 RegexFlag       `json:"flags,omitempty"`                   // (regex specific) Possible flags are ALL (default), ANYSTRING, COMPLEMENT, EMPTY, INTERSECTION, INTERVAL, or NONE. See regexp-syntax for their meaning
	MaxDeterminizedStates *int64          `json:"max_determinized_states,omitempty"` // (regex specific) Max states for regex execution (default 10000)
	Fuzzy                 *FuzzyS         `json:"fuzzy,omitempty"`
	ContextClauses        []ContextClause `json:"-"`
	Contexts              map[string]any  `json:"contexts,omitempty"` // To achieve suggestion filtering and/or boosting, you can add context mappings while configuring a completion field.
}

func (c CompletionSuggesterS) SuggestName() string {
	return c.Name
}

func (cs CompletionSuggesterS) MarshalJSON() ([]byte, error) {
	type CompletionBase CompletionSuggesterS

	cs.Contexts = map[string]any{}

	for _, ctx := range cs.ContextClauses {
		cs.Contexts[ctx.ContextName()] = ctx
	}

	suggesterBody := effdsl.M{
		"completion": (CompletionBase)(cs),
	}

	if cs.Prefix != "" {
		suggesterBody["prefix"] = cs.Prefix
	} else if cs.Regex != "" {
		suggesterBody["regex"] = cs.Regex
	}

	return json.Marshal(suggesterBody)
}

type CompletionSuggesterOption func(*CompletionSuggesterS) error

func WithSize(size uint64) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) error {
		c.Size = &size
		return nil
	}
}

func WithCompletionSuggesterSkipDuplicates(skipDuplicates bool) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) error {
		c.SkipDuplicates = skipDuplicates
		return nil
	}
}

type FuzzyS struct {
	Fuzziness      string  `json:"fuzziness,omitempty"`      // The fuzziness factor, defaults to AUTO.
	Transpositions *bool   `json:"transpositions,omitempty"` // If set to true, transpositions are counted as one change instead of two, defaults to true.
	MinLength      *uint64 `json:"min_length,omitempty"`     // Minimum length of the input before fuzzy suggestions are returned, defaults 3.
	PrefixLength   *uint64 `json:"prefix_length,omitempty"`  // Minimum length of the input, which is not checked for fuzzy alternatives, defaults to 1.
	UnicodeAware   bool    `json:"unicode_aware,omitempty"`  // If true, all measurements (like fuzzy edit distance, transpositions, and lengths) are measured in Unicode code points instead of in bytes. This is slightly slower than raw bytes, so it is set to false by default.
}

type RegexFlag string

const (
	RegexFlagAll          RegexFlag = "ALL"
	RegexFlagAnyString    RegexFlag = "ANYSTRING"
	RegexFlagComplement   RegexFlag = "COMPLEMENT"
	RegexFlagEmpty        RegexFlag = "EMPTY"
	RegexFlagIntersection RegexFlag = "INTERSECTION"
	RegexFlagInterval     RegexFlag = "INTERVAL"
	RegexFlagNone         RegexFlag = "NONE"
)

func WithRegexFlags(flag RegexFlag) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) error {
		if c.Regex == "" {
			return errors.New("WithRegexFlags can only be used with CompletionSuggesterRegex")
		}
		c.Flags = flag
		return nil
	}
}

func WithMaxDeterminizedStates(max int64) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) error {
		if c.Regex == "" {
			return errors.New("WithMaxDeterminizedStates can only be used with CompletionSuggesterRegex")
		}
		c.MaxDeterminizedStates = &max
		return nil
	}
}

type FuzzyOption func(*FuzzyS)

func WithFuzzyFuzziness(fuzziness string) FuzzyOption {
	return func(f *FuzzyS) {
		f.Fuzziness = fuzziness
	}
}

func WithFuzzyTranspositions(transpositions bool) FuzzyOption {
	return func(f *FuzzyS) {
		f.Transpositions = &transpositions
	}
}

func WithFuzzyMinLength(minLength uint64) FuzzyOption {
	return func(f *FuzzyS) {
		f.MinLength = &minLength
	}
}

func WithFuzzyPrefixLength(prefixLength uint64) FuzzyOption {
	return func(f *FuzzyS) {
		f.PrefixLength = &prefixLength
	}
}

func WithFuzzyUnicodeAware(unicodeAware bool) FuzzyOption {
	return func(f *FuzzyS) {
		f.UnicodeAware = unicodeAware
	}
}

func completionSuggesterFuzzy(opt ...FuzzyOption) FuzzyS {
	fuzzy := FuzzyS{}
	for _, o := range opt {
		o(&fuzzy)
	}
	return fuzzy
}

func WithCompletionSuggesterFuzzy(opt ...FuzzyOption) CompletionSuggesterOption {
	fuzzy := completionSuggesterFuzzy(opt...)
	return func(c *CompletionSuggesterS) error {
		c.Fuzzy = &fuzzy
		return nil
	}
}

type ContextClause interface {
	ContextName() string
}

type MultipleCategoryClause struct {
	Name     string   `json:"-"`
	Contexts []string `json:"-"`
}

func (m MultipleCategoryClause) ContextName() string { return m.Name }
func (m MultipleCategoryClause) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Contexts)
}

func WithMultipleCategoryContext(name string, contexts ...string) CompletionSuggesterOption {
	return func(c *CompletionSuggesterS) error {
		if len(contexts) == 0 {
			return errors.New("contexts cannot be empty for MultipleCategoryClause")
		}

		mcc := MultipleCategoryClause{
			Name:     name,
			Contexts: contexts,
		}

		c.ContextClauses = append(c.ContextClauses, mcc)

		return nil
	}
}

type CategoryContextClause struct {
	Name    string   `json:"-"`
	Context string   `json:"context"`
	Boost   *float64 `json:"boost,omitempty"`
	Prefix  bool     `json:"prefix,omitempty"`
}

func (c CategoryContextClause) ContextName() string {
	return c.Name
}

type CategoryContextClauseOption func(*CategoryContextClause)

func WithCategoryContextBoost(boost float64) CategoryContextClauseOption {
	return func(ccc *CategoryContextClause) {
		ccc.Boost = &boost
		return
	}
}

func WithCategoryContextPrefix() CategoryContextClauseOption {
	return func(ccc *CategoryContextClause) {
		ccc.Prefix = true
		return
	}
}

func NewCategoryContextClause(name string, context string, opts ...CategoryContextClauseOption) CategoryContextClause {
	ccc := CategoryContextClause{
		Name:    name,
		Context: context,
	}

	for _, opt := range opts {
		opt(&ccc)
	}

	return ccc
}

func WithCategoryContext(name, context string, opts ...CategoryContextClauseOption) CompletionSuggesterOption {
	return func(css *CompletionSuggesterS) error {
		ccc := NewCategoryContextClause(name, context, opts...)
		css.ContextClauses = append(css.ContextClauses, ccc)
		return nil
	}
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GeoContextClause struct {
	Name       string   `json:"-"`
	Context    GeoPoint `json:"context"`
	Boost      *float64 `json:"boost,omitempty"`
	Precision  string   `json:"precision,omitempty"`
	Neighbours []string `json:"neighbours,omitempty"`
}

func (g GeoContextClause) ContextName() string { return g.Name }

type GeoContextClauseOption func(*GeoContextClause)

func WithGeoContextBoost(boost float64) GeoContextClauseOption {
	return func(gcc *GeoContextClause) {
		gcc.Boost = &boost
		return
	}
}

func WithGeoContextPrecision(precision string) GeoContextClauseOption {
	return func(gcc *GeoContextClause) {
		gcc.Precision = precision
		return
	}
}

func WithGeoContextNeighbours(neighbours ...string) GeoContextClauseOption {
	return func(gcc *GeoContextClause) {
		gcc.Neighbours = neighbours
		return
	}
}
func NewGeoContextClause(name string, lat, lon float64, opts ...GeoContextClauseOption) GeoContextClause {
	gcc := GeoContextClause{
		Name:    name,
		Context: GeoPoint{lat, lon},
	}

	for _, opt := range opts {
		opt(&gcc)
	}

	return gcc
}

func WithGeoContext(name string, lat, lon float64, opts ...GeoContextClauseOption) CompletionSuggesterOption {
	return func(css *CompletionSuggesterS) error {
		gcc := NewGeoContextClause(name, lat, lon, opts...)
		css.ContextClauses = append(css.ContextClauses, gcc)
		return nil
	}
}

// The completion suggester provides auto-complete/search-as-you-type functionality. This is a navigational feature to guide users to relevant results as they are typing, improving search precision. It is not meant for spell correction or did-you-mean functionality like the term or phrase suggesters.
// [CompletionSuggester]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
func CompletionSuggester(name, prefix, field string, opts ...CompletionSuggesterOption) effdsl.SuggestResult {
	cs := CompletionSuggesterS{
		Name:   name,
		Prefix: prefix,
		Field:  field,
	}

	var err error
	for _, opt := range opts {
		err = opt(&cs)
		if err != nil {
			break
		}
	}

	return effdsl.SuggestResult{
		Ok:  cs,
		Err: err,
	}
}

// The completion suggester also supports regex queries meaning you can express a prefix as a regular expression.
// [CompletionSuggesterRegex]: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#regex
func CompletionSuggesterRegex(name, regex, field string, opts ...CompletionSuggesterOption) effdsl.SuggestResult {
	cs := CompletionSuggesterS{
		Name:  name,
		Regex: regex,
		Field: field,
	}

	var err error
	for _, opt := range opts {
		err = opt(&cs)
		if err != nil {
			break
		}
	}

	return effdsl.SuggestResult{
		Ok:  cs,
		Err: err,
	}
}
