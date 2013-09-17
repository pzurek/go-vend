package vend

type User struct {
	Id          *string `json:"id,omitempty"`
	Username    *string `json:"username,omitempty"`
	Name        *string `json:"name,omitempty"`
	OutletId    *string `json:"outlet_id,omitempty"`
	OutletName  *string `json:"outlet_,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
}
