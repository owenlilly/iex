package iex

// CashFlow pulls cash flow data. Available quarterly or annually, with the
// default being the last available quarter.
type CashFlow struct {
	Symbol string `json:"symbol"`

	Cashflow []struct {
		ReportDate              string `json:"reportDate"`
		NetIncome               int64  `json:"netIncome"`
		Depreciation            int64  `json:"depreciation"`
		ChangesInReceivables    int64  `json:"changesInReceivables"`
		ChangesInInventories    int64  `json:"changesInInventories"`
		CashChange              int64  `json:"cashChange"`
		CashFlow                int64  `json:"cashFlow"`
		CapitalExpenditures     int64  `json:"capitalExpenditures"`
		Investments             int64  `json:"investments"`
		InvestingActivityOther  int64  `json:"investingActivityOther"`
		TotalInvestingCashFlows int64  `json:"totalInvestingCashFlows"`
		DividendsPaid           int64  `json:"dividendsPaid"`
		NetBorrowings           int64  `json:"netBorrowings"`
		OtherFinancingCashFlows int64  `json:"otherFinancingCashFlows"`
		CashFlowFinancing       int64  `json:"cashFlowFinancing"`
		ExchangeRateEffect      int64  `json:"exchangeRateEffect"`
	}
}
