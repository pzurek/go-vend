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
