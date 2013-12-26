/*

{
  "id": "ade99704-587e-11e3-a29a-bc305bf5da20",
  "retailer_id": "02cc34c5-0adc-11e3-a415-bc764e10976c",
  "name": "Shopify",
  "description": null,
  "source": "USER",
  "contact": {
    "first_name": "Created for Supplier Shopify",
    "last_name": "Created for Supplier Shopify",
    "company_name": null,
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
}

*/

package vend

import "fmt"

type Supplier struct {
	Id          *string  `json:"id,omitempty"`
	RetailerId  *string  `json:"retailer_id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Contact     *Contact `json:"contact,omitempty"`
}

type SupplierResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Suppliers  *[]Supplier `json:"suppliers,omitempty"`
}

type SupplierService struct {
	client *Client
}

func (s *SupplierService) List() ([]Supplier, error) {

	resource := make([]Supplier, 0)

	rp, pagination, _, err := s.getPage(1, 200)

	if err != nil {
		return nil, err
	}

	resource = append(resource, *rp...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			rp, pg, _, err := s.getPage(*pagination.Page+1, 200)
			if err != nil {
				return nil, err
			}
			pagination = pg
			resource = append(resource, *rp...)
		}
	}

	return resource, err
}

func (s *SupplierService) getPage(p, ps int) (*[]Supplier, *Pagination, *Response, error) {
	u := fmt.Sprintf("supplier?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(SupplierResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Suppliers
	return resource, pagination, resp, err
}
