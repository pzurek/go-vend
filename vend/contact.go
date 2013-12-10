/*
"contact": {
    "first_name": null,
    "last_name": null,
    "company_name": "piotr",
    "phone": null,
    "mobile": null,
    "fax": null,
    "email": null,
    "twitter": null,
    "website": null,
    "physical_address1": null,
    "physical_address2": null,
    "physical_suburb": null,
    "physical_city": null,
    "physical_postcode": null,
    "physical_state": null,
    "physical_country_id": null,
    "postal_address1": null,
    "postal_address2": null,
    "postal_suburb": null,
    "postal_city": null,
    "postal_postcode": null,
    "postal_state": null,
    "postal_country_id": null
  }
*/

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
