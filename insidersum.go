package iex

// InsiderSum ...
type InsiderSum []struct {
	FullName      string `json:"fullName"`
	NetTransacted int64  `json:"netTransacted"`
	ReportedTitle string `json:"reportedTitle"`
	TotalBought   int64  `json:"totalBought"`
	TotalSold     int64  `json:"totalSold"`
}
