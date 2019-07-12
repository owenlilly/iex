package iex

// FundOwnership returns the top 10 fund holders, meaning any firm not defined
// as buy-side or sell-side such as mutual funds, pension funds, endowments,
// investment firms, and other large entities that manage funds on behalf of
// others.
type FundOwnership []struct {
	AdjHolding       uint64 `json:"adjHolding"`
	AdjMv            uint64 `json:"adjMv"`
	EntityProperName string `json:"entityProperName"`
	ReportDate       uint64 `json:"reportDate"`
	ReportedHolding  uint64 `json:"reportedHolding"`
	ReportedMv       uint64 `json:"reportedMv"`
}
