package vce

type SalesPayload struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Sales      *[]Sale     `json:"register_sales,omitempty"`
}
