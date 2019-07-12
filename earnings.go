package iex

// Earnings returns the earnings data for a given company including the actual
// EPS, consensus, and fiscal period. Earnings are available quarterly (last
// 4 quarters).
type Earnings struct {
	Symbol string `json:"symbol"`

	Earnings []struct {
		ActualEPS            float64 `json:"actualEPS"`
		ConsensusEPS         float64 `json:"consensusEPS"`
		AnnounceTime         string  `json:"announceTime"`
		NumberOfEstimates    uint64  `json:"numberOfEstimates"`
		EPSSurpriseDollar    float64 `json:"EPSSurpriseDollar"`
		EPSReportDate        string  `json:"EPSReportDate"`
		FiscalPeriod         string  `json:"fiscalPeriod"`
		FiscalEndDate        string  `json:"fiscalEndDate"`
		YearAgo              float64 `json:"yearAgo"`
		YearAgoChangePercent float64 `json:"yearAgoChangePercent"`
	} `json:"earnings"`
}
