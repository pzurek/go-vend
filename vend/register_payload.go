package vce

type RegisterPayload struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Registers  *[]Register `json:"registers,omitempty"`
}
