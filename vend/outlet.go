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

import "fmt"

type Outlet struct {
	Id                *string  `json:"id,omitempty"`
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

type OutletResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Outlets    *[]Outlet   `json:"outlets,omitempty"`
}

type OutletService struct {
	client *Client
}

func (s *OutletService) List() ([]Outlet, error) {

	outlets := make([]Outlet, 0)

	outl, pagination, _, err := s.getOutletPage(1, 50)

	if err != nil {
		return nil, err
	}

	outlets = append(outlets, *outl...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			outl, pag, _, err := s.getOutletPage(*pagination.Page+1, 50)
			if err != nil {
				return nil, err
			}
			pagination = pag
			outlets = append(outlets, *outl...)
		}
	}

	return outlets, err
}

func (s *OutletService) getOutletPage(p, ps int) (*[]Outlet, *Pagination, *Response, error) {
	u := fmt.Sprintf("outlets?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(OutletResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	outlets := response.Outlets
	return outlets, pagination, resp, err
}
