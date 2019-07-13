package iex

// IPOCal returns a list of upcoming or today IPOs scheduled for the
// current and next month. The response is split into two structures:
// rawData and viewData. rawData represents all available data for an IPO.
// viewData represents data structured for display to a user.
type IPOCal struct {
	RawData []struct {
		Symbol                 string   `json:"symbol"`
		CompanyName            string   `json:"companyName"`
		ExpectedDate           string   `json:"expectedDate"`
		LeadUnderwriters       []string `json:"leadUnderwriters"`
		Underwriters           []string `json:"underwriters"`
		CompanyCounsel         []string `json:"companyCounsel"`
		UnderwriterCounsel     []string `json:"underwriterCounsel"`
		Auditor                string   `json:"auditor"`
		Market                 string   `json:"market"`
		CIK                    string   `json:"cik"`
		Address                string   `json:"address"`
		City                   string   `json:"city"`
		State                  string   `json:"state"`
		Zip                    string   `json:"zip"`
		Phone                  string   `json:"phone"`
		CEO                    string   `json:"ceo"`
		Employees              uint64   `json:"employees"`
		URL                    string   `json:"url"`
		Status                 string   `json:"status"`
		SharesOffered          uint64   `json:"sharesOffered"`
		PriceLow               uint64   `json:"priceLow"`
		PriceHigh              uint64   `json:"priceHigh"`
		OfferAmount            uint64   `json:"offerAmount"`
		TotalExpenses          uint64   `json:"totalExpenses"`
		SharesOverAlloted      uint64   `json:"sharesOverAlloted"`
		ShareholderShares      uint64   `json:"shareholderShares"`
		SharesOutstanding      uint64   `json:"sharesOutstanding"`
		LockupPeriodExpiration string   `json:"lockupPeriodExpiration"`
		QuietPeriodExpiration  string   `json:"quietPeriodExpiration"`
		Revenue                uint64   `json:"revenue"`
		NetIncome              int64    `json:"netIncome"`
		TotalAssets            uint64   `json:"totalAssets"`
		TotalLiabilities       uint64   `json:"totalLiabilities"`
		StockholderEquity      uint64   `json:"stockholderEquity"`
		CompanyDescription     string   `json:"companyDescription"`
		BusinessDescription    string   `json:"businessDescription"`
		UseOfProceeds          string   `json:"useOfProceeds"`
		Competition            string   `json:"competition"`
		Amount                 uint64   `json:"amount"`
		PercentOffered         string   `json:"percentOffered"`
	} `json:"rawData"`

	ViewData []struct {
		Company  string `json:"Company"`
		Symbol   string `json:"Symbol"`
		Price    string `json:"Price"`
		Shares   string `json:"Shares"`
		Amount   string `json:"Amount"`
		Float    string `json:"Float"`
		Percent  string `json:"Percent"`
		Market   string `json:"Market"`
		Expected string `json:"Expected"`
	} `json:"viewData"`
}
