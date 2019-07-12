package iex

// BalanceSheet pulls balance sheet data. Available quarterly or annually with the default being the last available quarter.
type BalanceSheet struct {
	Symbol string `json:"symbol"`

	BalanceSheet []struct {
		ReportDate              string `json:"reportDate"`
		CurrentCash             uint64 `json:"currentCash"`
		ShortTermInvestments    uint64 `json:"shortTermInvestments"`
		Receivables             uint64 `json:"receivables"`
		Inventory               uint64 `json:"inventory"`
		OtherCurrentAssets      uint64 `json:"otherCurrentAssets"`
		CurrentAssets           uint64 `json:"currentAssets"`
		LongTermInvestments     uint64 `json:"longTermInvestments"`
		PropertyPlantEquipment  uint64 `json:"propertyPlantEquipment"`
		Goodwill                uint64 `json:"goodwill"`
		IntangibleAssets        uint64 `json:"intangibleAssets"`
		OtherAssets             uint64 `json:"otherAssets"`
		TotalAssets             uint64 `json:"totalAssets"`
		AccountsPayable         uint64 `json:"accountsPayable"`
		CurrentLongTermDebt     uint64 `json:"currentLongTermDebt"`
		OtherCurrentLiabilities uint64 `json:"otherCurrentLiabilities"`
		TotalCurrentLiabilities uint64 `json:"totalCurrentLiabilities"`
		LongTermDebt            uint64 `json:"longTermDebt"`
		OtherLiabilities        uint64 `json:"otherLiabilities"`
		MinorityInterest        uint64 `json:"minorityInterest"`
		TotalLiabilities        uint64 `json:"totalLiabilities"`
		CommonStock             uint64 `json:"commonStock"`
		RetainedEarnings        uint64 `json:"retainedEarnings"`
		TreasuryStock           uint64 `json:"treasuryStock"`
		CapitalSurplus          uint64 `json:"capitalSurplus"`
		ShareholderEquity       uint64 `json:"shareholderEquity"`
		NetTangibleAssets       uint64 `json:"netTangibleAssets"`
	} `json:"balanceSheet"`
}
