package iex

// SecPerf ...
type SecPerf []struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Performance float64 `json:"performance"`
	LastUpdated uint64  `json:"lastUpdated"`
}
