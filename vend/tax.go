/*

{
  "id": "41ab729c-44d4-11e3-a29a-bc305bf5da20",
  "name": "NZ GST",
  "rate": 0.15,
  "default": true,
  "rates": [
    {
      "id": "622d3131-3562-2d34-3031-646231356665",
      "rate": 0.15,
      "name": "NZ GST"
    }
  ]
}

*/

package vend

import "fmt"

type Tax struct {
	Id      *string    `json:"id,omitempty"`
	Name    *string    `json:"tax,omitempty"`
	Rate    *float64   `json:"tax_rate,omitempty"`
	Default *bool      `json:"default,omitempty"`
	Rates   *[]TaxRate `json:"rates,omitempty"`
}

type TaxRate struct {
	Id   *string  `json:"id,omitempty"`
	Name *string  `json:"tax,omitempty"`
	Rate *float64 `json:"tax_rate,omitempty"`
}

type TaxResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Taxes      *[]Tax      `json:"taxes,omitempty"`
}

type TaxService struct {
	client *Client
}

func (s *TaxService) List() ([]Tax, error) {

	resource := make([]Tax, 0)

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

func (s *TaxService) getPage(p, ps int) (*[]Tax, *Pagination, *Response, error) {
	u := fmt.Sprintf("taxes?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(TaxResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Taxes
	return resource, pagination, resp, err
}
