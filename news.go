package iex

// News ...
type News []struct {
	Datetime   uint64 `json:"datetime"`
	Headline   string `json:"headline"`
	Source     string `json:"source"`
	URL        string `json:"url"`
	Summary    string `json:"summary"`
	Related    string `json:"related"`
	Image      string `json:"image"`
	Lang       string `json:"lang"`
	HasPaywall bool   `json:"hasPaywall"`
}
