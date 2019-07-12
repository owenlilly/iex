package iex

// Estimates provides the latest consensus estimate for the next fiscal period.
type Estimates struct {
	Symbol string `json:"symbol"`

	Estimates []struct {
		ConsensusEPS      float64 `json:"consensusEPS"`
		NumberOfEstimates uint64  `json:"numberOfEstimates"`
		FiscalPeriod      string  `json:"fiscalPeriod"`
		FiscalEndDate     string  `json:"fiscalEndDate"`
		ReportDate        string  `json:"reportDate"`
	} `json:"estimates"`
}
