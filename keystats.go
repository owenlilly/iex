package iex

// KeyStats ...
type KeyStats struct {
	CompanyName         string  `json:"companyName"`
	Marketcap           uint64  `json:"marketcap"`
	Week52high          float64 `json:"week52high"`
	Week52low           float64 `json:"week52low"`
	Week52change        float64 `json:"week52change"`
	SharesOutstanding   uint64  `json:"sharesOutstanding"`
	Float               uint64  `json:"float"`
	Avg10Volume         float64 `json:"avg10Volume"`
	Avg30Volume         float64 `json:"avg30Volume"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	Employees           uint64  `json:"employees"`
	TTMEPS              float64 `json:"ttmEPS"`
	TTMDividendRate     float64 `json:"ttmDividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	NextDividendDate    string  `json:"nextDividendDate"`
	ExDividendDate      string  `json:"exDividendDate"`
	NextEarningsDate    string  `json:"nextEarningsDate"`
	PERatio             float64 `json:"peRatio"`
	Beta                float64 `json:"beta"`
	MaxChangePercent    float64 `json:"maxChangePercent"`
	Year5ChangePercent  float64 `json:"year5ChangePercent"`
	Year2ChangePercent  float64 `json:"year2ChangePercent"`
	Year1ChangePercent  float64 `json:"year1ChangePercent"`
	YTDChangePercent    float64 `json:"ytdChangePercent"`
	Month6ChangePercent float64 `json:"month6ChangePercent"`
	Month3ChangePercent float64 `json:"month3ChangePercent"`
	Month1ChangePercent float64 `json:"month1ChangePercent"`
	Day30ChangePercent  float64 `json:"day30ChangePercent"`
	Day5ChangePercent   float64 `json:"day5ChangePercent"`
}
