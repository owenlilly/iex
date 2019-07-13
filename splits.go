package iex

// Splits ...
type Splits []struct {
	ExDate       string
	DeclaredDate string
	Ratio        float64
	ToFactor     uint64
	FromFactor   uint64
	Description  string
}
