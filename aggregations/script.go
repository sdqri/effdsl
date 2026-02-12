package aggregations

type Script struct {
	Lang   string         `json:"lang,omitempty"`
	Source string         `json:"source,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}
