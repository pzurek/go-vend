/*

{
  "resource": [
    {
      "id": "02f9bd97-0adc-11e3-a415-bc764e10976c",
      "username": "p.zurek@gmail.com",
      "name": "A cashier",
      "outlet_id": "",
      "outlet_name": "",
      "account_type": "admin"
    }
  ]
}

*/

package vend

import "fmt"

type User struct {
	Id          *string `json:"id,omitempty"`
	Username    *string `json:"username,omitempty"`
	Name        *string `json:"name,omitempty"`
	OutletId    *string `json:"outlet_id,omitempty"`
	OutletName  *string `json:"outlet_,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
}

type UserResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Users      *[]User     `json:"users,omitempty"`
}

type UserService struct {
	client *Client
}

func (s *UserService) List() ([]User, error) {

	resource := make([]User, 0)

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

func (s *UserService) getPage(p, ps int) (*[]User, *Pagination, *Response, error) {
	u := fmt.Sprintf("users?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(UserResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Users
	return resource, pagination, resp, err
}
