package iex

// Company ...
type Company struct {
	Symbol       string   `json:"symbol"`
	CompanyName  string   `json:"companyName"`
	Exchange     string   `json:"exchange"`
	Industry     string   `json:"industry"`
	Website      string   `json:"website"`
	Description  string   `json:"description"`
	CEO          string   `json:"CEO"`
	SecurityName string   `json:"securityName"`
	IssueType    string   `json:"issueType"`
	Sector       string   `json:"sector"`
	Employees    uint64   `json:"employees"`
	Tags         []string `json:"tags"`
}
