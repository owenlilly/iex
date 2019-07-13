package iex

// InstlOwnership ...
type InstlOwnership []struct {
	AdjHolding       uint64 `json:"adjHolding"`
	AdjMv            uint64 `json:"adjMv"`
	EntityProperName string `json:"entityProperName"`
	ReportDate       uint64 `json:"reportDate"`
	ReportedHolding  uint64 `json:"reportedHolding"`
}
