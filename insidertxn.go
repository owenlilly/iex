package iex

// InsiderTxn ...
type InsiderTxn []struct {
	EffectiveDate uint64  `json:"effectiveDate"`
	FullName      string  `json:"fullName"`
	ReportedTitle string  `json:"reportedTitle"`
	TranPrice     float64 `json:"tranPrice"`
	TranShares    int64   `json:"tranShares"`
	TranValue     float64 `json:"tranValue"`
}
