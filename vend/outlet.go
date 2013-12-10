/*
{
  "id": "02f399e9-0adc-11e3-a415-bc764e10976c",
  "retailer_id": "02cc34c5-0adc-11e3-a415-bc764e10976c",
  "name": "Main Outlet",
  "time_zone": "America/New_York",
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
  },
  "email": null,
  "physical_address1": null,
  "physical_address2": null,
  "physical_suburb": null,
  "physical_city": null,
  "physical_postcode": null,
  "physical_state": null,
  "physical_country_id": null
}
*/

package vend

type Outlet struct {
	Id                *string  `json:"d,omitempty"`
	RetailerId        *string  `json:"retailer_id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	TimeZone          *string  `json:"timezone,omitempty"`
	Contact           *Contact `json:"contact,omitempty"`
	Email             *string  `json:"email,omitempty"`
	PhysicalAddress1  *string  `json:"physical_address1,omitempty"`
	PhysicalAddress2  *string  `json:"physical_address2,omitempty"`
	PhysicalSuburb    *string  `json:"physical_suburb,omitempty"`
	PhysicalCity      *string  `json:"physical_city,omitempty"`
	PhysicalPostcode  *string  `json:"physical_postcode,omitempty"`
	PhysicalState     *string  `json:"physical_state,omitempty"`
	PhysicalCountryId *string  `json:"physical_country_id,omitempty"`
}
