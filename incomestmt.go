package iex

// IncomeStmt ...
type IncomeStmt struct {
	Symbol string `json:"symbol"`
	Income []struct {
		ReportDate             string `json:"reportDate"`
		TotalRevenue           uint64 `json:"totalRevenue"`
		CostOfRevenue          uint64 `json:"costOfRevenue"`
		GrossProfit            uint64 `json:"grossProfit"`
		ResearchAndDevelopment uint64 `json:"researchAndDevelopment"`
		SellingGeneralAndAdmin uint64 `json:"sellingGeneralAndAdmin"`
		OperatingExpense       uint64 `json:"operatingExpense"`
		OperatingIncome        uint64 `json:"operatingIncome"`
		OtherIncomeExpenseNet  uint64 `json:"otherIncomeExpenseNet"`
		EBIT                   uint64 `json:"ebit"`
		InterestIncome         uint64 `json:"interestIncome"`
		PretaxIncome           uint64 `json:"pretaxIncome"`
		IncomeTax              uint64 `json:"incomeTax"`
		MinorityInterest       uint64 `json:"minorityInterest"`
		NetIncome              uint64 `json:"netIncome"`
		NetIncomeBasic         uint64 `json:"netIncomeBasic"`
	} `json:"income"`
}
