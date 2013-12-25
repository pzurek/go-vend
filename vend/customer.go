/*
   "customer": {
     "id": "02fb2229-0adc-11e3-a415-bc764e10976c",
     "name": "",
     "customer_code": "WALKIN",
     "customer_group_id": "02fa2d73-0adc-11e3-a415-bc764e10976c",
     "customer_group_name": "All Customers",
     "updated_at": "2013-11-14 20:41:04",
     "deleted_at": "",
     "balance": "0.000",
     "year_to_date": "5884.37012",
     "date_of_birth": "",
     "sex": "",
     "custom_field_1": "",
     "custom_field_2": "",
     "custom_field_3": "",
     "custom_field_4": "",
     "contact": {}
   },

*/

package vend

import "fmt"

type Customer struct {
	Id                *string          `json:"id,omitempty"`
	Name              *string          `json:"name,omitempty"`
	CustomerCode      *string          `json:"customer_code,omitempty"`
	CustomerGroupId   *string          `json:"customer_group_id,omitempty"`
	CustomerGroupName *string          `json:"customer_group_name,omitempty"`
	FirstName         *string          `json:"first_name,omitempty"`
	LastName          *string          `json:"last_name,omitempty"`
	CompanyName       *string          `json:"company_name,omitempty"`
	Phone             *string          `json:"phone,omitempty"`
	Mobile            *string          `json:"mobile,omitempty"`
	Fax               *string          `json:"fax,omitempty"`
	Email             *string          `json:"email,omitempty"`
	Twitter           *string          `json:"twitter,omitempty"`
	Website           *string          `json:"website,omitempty"`
	PhysicalAddress1  *string          `json:"physical_address1,omitempty"`
	PhysicalAddress2  *string          `json:"physical_address2,omitempty"`
	PhysicalSuburb    *string          `json:"physical_suburb,omitempty"`
	PhysicalCity      *string          `json:"physical_city,omitempty"`
	PhysicalPostcode  *string          `json:"physical_postcode,omitempty"`
	PhysicalState     *string          `json:"physical_state,omitempty"`
	PhysicalCountryId *string          `json:"physical_country_id,omitempty"`
	PostalAddress1    *string          `json:"postal_address1,omitempty"`
	PostalAddress2    *string          `json:"postal_address2,omitempty"`
	PostalSuburb      *string          `json:"postal_suburb,omitempty"`
	PostalCity        *string          `json:"postal_city,omitempty"`
	PostalPostcode    *string          `json:"postal_postcode,omitempty"`
	PostalState       *string          `json:"postal_state,omitempty"`
	PostalCountryId   *string          `json:"postal_country_id,omitempty"`
	EnableLoyalty     *int             `json:"enable_loayaly,omitempty"`
	LoyaltyBalance    *string          `json:"loyalty_balance,omitempty"`
	UpdatedAt         *string          `json:"updated_at,omitempty"`
	DeletedAt         *string          `json:"deleted_at,omitempty"`
	Balance           *string          `json:"balance,omitempty"`
	YearToDate        *string          `json:"year_to_date,omitempty"`
	DateOfBirth       *string          `json:"date_of_birth,omitempty"`
	Sex               *string          `json:"sex,omitempty"`
	CustomField1      *string          `json:"custom_field_1,omitempty"`
	CustomField2      *string          `json:"custom_field_2,omitempty"`
	CustomField3      *string          `json:"custom_field_3,omitempty"`
	CustomField4      *string          `json:"custom_field_4,omitempty"`
	Contact           *CustomerContact `json:"contact,omitempty"`
}

// Contact type used by the Customer type
type CustomerContact struct {
	CompanyName *string `json:"company_name,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
}

type CustomerResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Customers  *[]Customer `json:"customers,omitempty"`
}

type CustomerService struct {
	client *Client
}

func (s *CustomerService) List() ([]Customer, error) {

	customers := make([]Customer, 0)

	cust, pagination, _, err := s.getCustomerPage(1, 200)

	if err != nil {
		return nil, err
	}

	customers = append(customers, *cust...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			cust, pag, _, err := s.getCustomerPage(*pagination.Page+1, 200)
			if err != nil {
				return nil, err
			}
			pagination = pag
			customers = append(customers, *cust...)
		}
	}

	return customers, err
}

func (s *CustomerService) getCustomerPage(p, ps int) (*[]Customer, *Pagination, *Response, error) {
	u := fmt.Sprintf("customers?page=%v&page_size=%v", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	custuctResp := new(CustomerResponse)
	resp, err := s.client.Do(req, custuctResp)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := custuctResp.Pagination
	customers := custuctResp.Customers
	return customers, pagination, resp, err
}
