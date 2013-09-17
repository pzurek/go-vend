package vend

type CustomerContact struct {
	CompanyName *string `json:"company_name,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
}
