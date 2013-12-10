package vce

type OutletPayload struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Outlets    *[]Outlet   `json:"outlets,omitempty"`
}
