package dto

// ResponseList ...
type ResponseList struct {
	Total    uint64      `json:"total_data"`
	Limit    uint64      `json:"limit"`
	Page     uint64      `json:"page"`
	LastPage uint64      `json:"last_page"`
	Data     interface{} `json:"data"`
}
