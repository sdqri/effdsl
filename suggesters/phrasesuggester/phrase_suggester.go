package phrasesuggester

import (
	"encoding/json"

	"github.com/sdqri/effdsl/v2"
)

// PhraseSuggesterS - https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters-phrase.html
type PhraseSuggesterS struct {
	Name                    string            `json:"-"`                                    // Required
	Text                    string            `json:"-"`                                    // Required
	Field                   string            `json:"field"`                                // Required
	GramSize                *uint64           `json:"gram_size,omitempty"`                  // Sets max size of the n-grams (shingles) in the field. If the field doesn't contain n-grams (shingles), this should be omitted or set to 1. Note that Elasticsearch tries to detect the gram size based on the specified field. If the field uses a shingle filter, the gram_size is set to the max_shingle_size if not explicitly set.
	RealWordErrorLikelihood *float64          `json:"real_word_error_likelihood,omitempty"` // The likelihood of a term being misspelled even if the term exists in the dictionary. The default is 0.95, meaning 5% of the real words are misspelled.
	Confidence              *float64          `json:"confidence,omitempty"`                 // The confidence level defines a factor applied to the input phrases score which is used as a threshold for other suggest candidates. Only candidates that score higher than the threshold will be included in the result. For instance a confidence level of 1.0 will only return suggestions that score higher than the input phrase. If set to 0.0 the top N candidates are returned. The default is 1.0.
	MaxErrors               *float64          `json:"max_errors,omitempty"`                 // The maximum percentage of the terms considered to be misspellings in order to form a correction. This method accepts a float value in the range [0..1) as a fraction of the actual query terms or a number >=1 as an absolute number of query terms. The default is set to 1.0, meaning only corrections with at most one misspelled term are returned. Note that setting this too high can negatively impact performance. Low values like 1 or 2 are recommended; otherwise the time spend in suggest calls might exceed the time spend in query execution.
	Separator               *string           `json:"separator,omitempty"`                  // The separator that is used to separate terms in the bigram field. If not set the whitespace character is used as a separator.
	Size                    *uint64           `json:"size,omitempty"`                       // The number of candidates that are generated for each individual query term. Low numbers like 3 or 5 typically produce good results. Raising this can bring up terms with higher edit distances. The default is 5.
	Analyzer                string            `json:"analyzer,omitempty"`                   // Sets the analyzer to analyze to suggest text with. Defaults to the search analyzer of the suggest field passed via field.
	ShardSize               *uint64           `json:"shard_size,omitempty"`                 // Sets the maximum number of suggested terms to be retrieved from each individual shard. During the reduce phase, only the top N suggestions are returned based on the size option. Defaults to 5.
	Highlight               *Highlight        `json:"highlight,omitempty"`                  // Sets up suggestion highlighting. If not provided then no highlighted field is returned. If provided must contain exactly pre_tag and post_tag, which are wrapped around the changed tokens. If multiple tokens in a row are changed the entire phrase of changed tokens is wrapped rather than each token.
	Collate                 *Collate          `json:"collate,omitempty"`                    // Checks each suggestion against the specified query to prune suggestions for which no matching docs exist in the index. The collate query for a suggestion is run only on the local shard from which the suggestion has been generated from. The query must be specified and it can be templated. Refer to Search templates. The current suggestion is automatically made available as the {{suggestion}} variable, which should be used in your query. You can still specify your own template params — the suggestion value will be added to the variables you specify. Additionally, you can specify a prune to control if all phrase suggestions will be returned; when set to true the suggestions will have an additional option collate_match, which will be true if matching documents for the phrase was found, false otherwise. The default value for prune is false.
	Smoothing               SmoothingModel    `json:"smoothing,omitempty"`                  // The phrase suggester supports multiple smoothing models to balance weight between infrequent grams (grams (shingles) are not existing in the index) and frequent grams (appear at least once in the index).
	DirectGenerator         []DirectGenerator `json:"direct_generator,omitempty"`
}

func (ps PhraseSuggesterS) SuggestName() string {
	return ps.Name
}

func (ps PhraseSuggesterS) MarshalJSON() ([]byte, error) {
	type PhraseSuggesterBase PhraseSuggesterS
	return json.Marshal(effdsl.M{
		"text":   ps.Text,
		"phrase": (PhraseSuggesterBase)(ps),
	})
}

type PhraseSuggesterOption func(*PhraseSuggesterS) error

func WithGramSize(size uint64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.GramSize = &size
		return nil
	}
}

func WithRealWordErrorLikelihood(val float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.RealWordErrorLikelihood = &val
		return nil
	}
}

func WithConfidence(conf float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Confidence = &conf
		return nil
	}
}

func WithMaxErrors(max float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.MaxErrors = &max
		return nil
	}
}

func WithSeparator(sep string) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Separator = &sep
		return nil
	}
}

func WithSize(size uint64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Size = &size
		return nil
	}
}

func WithAnalyzer(analyzer string) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Analyzer = analyzer
		return nil
	}
}

func WithShardSize(size uint64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.ShardSize = &size
		return nil
	}
}

type Highlight struct {
	PreTag  string `json:"pre_tag"`
	PostTag string `json:"post_tag"`
}

func WithHighlight(preTag, postTag string) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Highlight = &Highlight{
			PreTag:  preTag,
			PostTag: postTag,
		}
		return nil
	}
}

type SmoothingModel interface {
	json.Marshaler
}

type StupidBackoffSmoothing struct {
	Discount float64 `json:"discount"`
}

func (s StupidBackoffSmoothing) MarshalJSON() ([]byte, error) {
	type StupidBackoffSmoothingBase StupidBackoffSmoothing
	return json.Marshal(map[string]any{
		"stupid_backoff": (StupidBackoffSmoothingBase)(s),
	})
}

type LaplaceSmoothing struct {
	Alpha float64 `json:"alpha"`
}

func (l LaplaceSmoothing) MarshalJSON() ([]byte, error) {
	type LaplaceSmoothingBase LaplaceSmoothing
	return json.Marshal(map[string]any{
		"laplace": (LaplaceSmoothingBase)(l),
	})
}

type LinearInterpolationSmoothing struct {
	TrigramLambda float64 `json:"trigram_lambda"`
	BigramLambda  float64 `json:"bigram_lambda"`
	UnigramLambda float64 `json:"unigram_lambda"`
}

func (l LinearInterpolationSmoothing) MarshalJSON() ([]byte, error) {
	type LinearInterpolationSmoothingBase LinearInterpolationSmoothing
	return json.Marshal(map[string]any{
		"linear_interpolation": (LinearInterpolationSmoothingBase)(l),
	})
}

func WithStupidBackoffSmoothing(discount float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Smoothing = StupidBackoffSmoothing{Discount: discount}
		return nil
	}
}

func WithLaplaceSmoothing(alpha float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Smoothing = LaplaceSmoothing{Alpha: alpha}
		return nil
	}
}

func WithLinearInterpolationSmoothing(tri, bi, uni float64) PhraseSuggesterOption {
	return func(ps *PhraseSuggesterS) error {
		ps.Smoothing = LinearInterpolationSmoothing{
			TrigramLambda: tri,
			BigramLambda:  bi,
			UnigramLambda: uni,
		}
		return nil
	}
}

const CURRENT_SUGGESTION = "{{suggestion}}"

type Collate struct {
	Query  CollateQuery      `json:"query"`            // Required
	Params map[string]string `json:"params,omitempty"` // Optional template parameters
	Prune  *bool             `json:"prune,omitempty"`  // Optional
}

type CollateQuery struct {
	Source effdsl.Query `json:"source"` // Source is templated JSON with placeholders like {{suggestion}}
}

type WithCollateOption func(*Collate) error

func WithCollate(queryResult effdsl.QueryResult, opts ...WithCollateOption) PhraseSuggesterOption {
	collate := new(Collate)

	query := queryResult.Ok
	err := queryResult.Err

	if err != nil {
		return func(ps *PhraseSuggesterS) error {
			return err
		}
	}

	collate.Query = CollateQuery{Source: query}

	for _, opt := range opts {
		err = opt(collate)
		if err != nil {
			return func(ps *PhraseSuggesterS) error {
				return err
			}
		}
	}

	return func(ps *PhraseSuggesterS) error {
		ps.Collate = collate
		return nil
	}
}

func WithParams(params map[string]string) WithCollateOption {
	return func(c *Collate) error {
		c.Params = params
		return nil
	}
}

func WithPrune() WithCollateOption {
	return func(c *Collate) error {
		prune := true
		c.Prune = &prune
		return nil
	}
}

type SuggestMode string

const (
	// Only generate suggestions for terms that are not in the shard. This is the default.
	Missing SuggestMode = "missing"
	// Only suggest terms that occur in more docs on the shard than the original term.
	Popular SuggestMode = "popular"
	// Suggest any matching suggestions based on terms in the suggest text.
	Always SuggestMode = "always"
)

type DirectGenerator struct {
	Field          string       `json:"field"`                     // required: The field to fetch the candidate suggestions from. This is a required option that either needs to be set globally or per suggestion.
	Size           *uint64      `json:"size,omitempty"`            // The maximum corrections to be returned per suggest text token.
	SuggestMode    *SuggestMode `json:"suggest_mode,omitempty"`    // The suggest mode controls what suggestions are included on the suggestions generated on each shard. All values other than always can be thought of as an optimization to generate fewer suggestions to test on each shard and are not rechecked when combining the suggestions generated on each shard. Thus missing will generate suggestions for terms on shards that do not contain them even if other shards do contain them. Those should be filtered out using confidence. Three possible values can be specified:missing, popular, always
	MaxEdits       *int         `json:"max_edits,omitempty"`       // The maximum edit distance candidate suggestions can have in order to be considered as a suggestion. Can only be a value between 1 and 2. Any other value results in a bad request error being thrown. Defaults to 2.
	PrefixLength   *int         `json:"prefix_length,omitempty"`   // The number of minimal prefix characters that must match in order be a candidate suggestions. Defaults to 1. Increasing this number improves spellcheck performance. Usually misspellings don't occur in the beginning of terms.
	MinWordLength  *int         `json:"min_word_length,omitempty"` // The minimum length a suggest text term must have in order to be included. Defaults to 4.
	MaxInspections *int         `json:"max_inspections,omitempty"` // A factor that is used to multiply with the shard_size in order to inspect more candidate spelling corrections on the shard level. Can improve accuracy at the cost of performance. Defaults to 5.
	MinDocFreq     *float64     `json:"min_doc_freq,omitempty"`    // The minimal threshold in number of documents a suggestion should appear in. This can be specified as an absolute number or as a relative percentage of number of documents. This can improve quality by only suggesting high frequency terms. Defaults to 0f and is not enabled. If a value higher than 1 is specified, then the number cannot be fractional. The shard level document frequencies are used for this option.
	MaxTermFreq    *float64     `json:"max_term_freq,omitempty"`   // The maximum threshold in number of documents in which a suggest text token can exist in order to be included. Can be a relative percentage number (e.g., 0.4) or an absolute number to represent document frequencies. If a value higher than 1 is specified, then fractional can not be specified. Defaults to 0.01f. This can be used to exclude high frequency terms — which are usually spelled correctly — from being spellchecked. This also improves the spellcheck performance. The shard level document frequencies are used for this option.
	PreFilter      *string      `json:"pre_filter,omitempty"`      // A filter (analyzer) that is applied to each of the tokens passed to this candidate generator. This filter is applied to the original token before candidates are generated.
	PostFilter     *string      `json:"post_filter,omitempty"`     // A filter (analyzer) that is applied to each of the generated tokens before they are passed to the actual phrase scorer.
}

func WithDirectGenerator(field string, opts ...WithDirectGeneratorOption) PhraseSuggesterOption {
	gen := DirectGenerator{Field: field}

	for _, opt := range opts {
		opt(&gen)
	}

	return func(ps *PhraseSuggesterS) error {
		ps.DirectGenerator = append(ps.DirectGenerator, gen)
		return nil
	}
}

type WithDirectGeneratorOption func(*DirectGenerator)

func WithSuggestMode(mode SuggestMode) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.SuggestMode = &mode
	}
}

func WithDirectGeneratorSize(size uint64) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.Size = &size
	}
}

func WithMaxEdits(edits int) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.MaxEdits = &edits
	}
}

func WithPrefixLength(length int) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.PrefixLength = &length
	}
}

func WithMinWordLength(length int) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.MinWordLength = &length
	}
}

func WithMaxInspections(inspections int) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.MaxInspections = &inspections
	}
}

func WithMinDocFreq(freq float64) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.MinDocFreq = &freq
	}
}

func WithMaxTermFreq(freq float64) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.MaxTermFreq = &freq
	}
}

func WithPreFilter(analyzer string) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.PreFilter = &analyzer
	}
}

func WithPostFilter(analyzer string) WithDirectGeneratorOption {
	return func(dg *DirectGenerator) {
		dg.PostFilter = &analyzer
	}
}

// PhraseSuggester builds a phrase suggester with provided options
func PhraseSuggester(suggestName, text, field string, opts ...PhraseSuggesterOption) effdsl.SuggestResult {
	ps := PhraseSuggesterS{
		Name:  suggestName,
		Text:  text,
		Field: field,
	}

	var err error
	for _, opt := range opts {
		err = opt(&ps)
		if err != nil {
			break
		}

	}

	return effdsl.SuggestResult{
		Ok:  ps,
		Err: err,
	}
}
