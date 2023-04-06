package dto

// ResponseList ...
type ResponseList struct {
	Total    uint        `json:"total_data"`
	Limit    uint        `json:"limit"`
	Page     uint        `json:"page"`
	LastPage uint        `json:"last_page"`
	Data     interface{} `json:"data"`
}
