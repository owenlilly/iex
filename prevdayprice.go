package iex

// PrevDayPrice ...
type PrevDayPrice struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Volume         uint64  `json:"volume"`
	UOpen          float64 `json:"uOpen"`
	UClose         float64 `json:"uClose"`
	UHigh          float64 `json:"uHigh"`
	ULow           float64 `json:"uLow"`
	UVolume        uint64  `json:"uVolume"`
	Change         float64 `json:"change"`
	ChangePercent  float64 `json:"changePercent"`
	ChangeOverTime float64 `json:"changeOverTime"`
	Symbol         string  `json:"symbol"`
}
