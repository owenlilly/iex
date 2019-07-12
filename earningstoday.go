package iex

// EarningsToday returns earnings that will be reported today as two arrays:
// before the open bto and after market close amc. Each array contains an
// object with all keys from earnings, a quote object, and a headline key.
type EarningsToday struct {
	BTO []struct {
		ConsensusEPS      float64 `json:"consensusEPS"`
		AnnounceTime      string  `json:"announceTime"`
		NumberOfEstimates uint64  `json:"numberOfEstimates"`
		FiscalPeriod      string  `json:"fiscalPeriod"`
		FiscalEndDate     string  `json:"fiscalEndDate"`
		Symbol            string  `json:"symbol"`
		// Quote             *Quote  `json:"quote"`
	} `json:"bmc"`

	AMC []struct {
		ConsensusEPS      float64 `json:"consensusEPS"`
		AnnounceTime      string  `json:"announceTime"`
		NumberOfEstimates uint64  `json:"numberOfEstimates"`
		FiscalPeriod      string  `json:"fiscalPeriod"`
		FiscalEndDate     string  `json:"fiscalEndDate"`
		Symbol            string  `json:"symbol"`
		// Quote             *Quote  `json:"quote"`
	} `json:"amt"`

	DMT []struct {
		ConsensusEPS      float64 `json:"consensusEPS"`
		AnnounceTime      string  `json:"announceTime"`
		NumberOfEstimates uint64  `json:"numberOfEstimates"`
		FiscalPeriod      string  `json:"fiscalPeriod"`
		FiscalEndDate     string  `json:"fiscalEndDate"`
		Symbol            string  `json:"symbol"`
		// Quote             *Quote  `json:"quote"`
	} `json:"dmt"`
}
