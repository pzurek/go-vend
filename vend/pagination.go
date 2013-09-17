package vend

type Pagination struct {
	Results  *int `json:"results"`
	Page     *int `json:"page"`
	PageSize *int `json:"page_size"`
	Pages    *int `json:"pages"`
}
