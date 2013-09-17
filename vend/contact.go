package vend

type Contact struct {
	FirstName         *string `json:"first_name,omitempty"`
	LastName          *string `json:"last_name,omitempty"`
	CompanyName       *string `json:"company_name,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	Mobile            *string `json:"mobile,omitempty"`
	Fax               *string `json:"fax,omitempty"`
	Email             *string `json:"email,omitempty"`
	Twitter           *string `json:"twitter,omitempty"`
	Website           *string `json:"website,omitempty"`
	PhysicalAddress1  *string `json:"physical_address1,omitempty"`
	PhysicalAddress2  *string `json:"physical_address2,omitempty"`
	PhysicalSuburb    *string `json:"physical_suburb,omitempty"`
	PhysicalCity      *string `json:"physical_city,omitempty"`
	PhysicalPostcode  *string `json:"physical_postcode,omitempty"`
	PhysicalState     *string `json:"physical_state,omitempty"`
	PhysicalCountryId *string `json:"physical_country_id,omitempty"`
	PostalAddress1    *string `json:"postal_address1,omitempty"`
	PostalAddress2    *string `json:"postal_address2,omitempty"`
	PostalSuburb      *string `json:"postal_suburb,omitempty"`
	PostalCity        *string `json:"postal_city,omitempty"`
	PostalPostcode    *string `json:"postal_postcode,omitempty"`
	PostalState       *string `json:"postal_state,omitempty"`
	PostalCountryId   *string `json:"postal_country_id,omitempty"`
}
