package vend

import "fmt"

// Customer struct
type Customer struct {
	ID                *string          `json:"id,omitempty"`
	Name              *string          `json:"name,omitempty"`
	CustomerCode      *string          `json:"customer_code,omitempty"`
	CustomerGroupID   *string          `json:"customer_group_id,omitempty"`
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
	PhysicalCountryID *string          `json:"physical_country_id,omitempty"`
	PostalAddress1    *string          `json:"postal_address1,omitempty"`
	PostalAddress2    *string          `json:"postal_address2,omitempty"`
	PostalSuburb      *string          `json:"postal_suburb,omitempty"`
	PostalCity        *string          `json:"postal_city,omitempty"`
	PostalPostcode    *string          `json:"postal_postcode,omitempty"`
	PostalState       *string          `json:"postal_state,omitempty"`
	PostalCountryID   *string          `json:"postal_country_id,omitempty"`
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

// CustomerContact type used by the Customer type
type CustomerContact struct {
	CompanyName *string `json:"company_name,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
}

// CustomerResponse is used to unmarshal the standard customer collection reposnse
type CustomerResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Customers  []Customer  `json:"customers,omitempty"`
}

// CustomerService gives access to all customer related actions
type CustomerService struct {
	client *Client
}

// List returns a slice of all customers
func (s *CustomerService) List() ([]Customer, error) {

	var resource []Customer

	rp, pagination, _, err := s.getPage(1, 200)

	if err != nil {
		return nil, err
	}

	resource = append(resource, rp...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			rp, pg, _, err := s.getPage(*pagination.Page+1, 200)
			if err != nil {
				return nil, err
			}
			pagination = pg
			resource = append(resource, rp...)
		}
	}

	return resource, err
}

func (s *CustomerService) Update(c Customer) (Customer, error) {
	body := ""

	url := fmt.Sprintf("customers")

	req, err := s.client.NewRequest("POST", url, body)
	if err != nil {
		return Customer{}, fmt.Errorf("creating new request failerd: %s\n", err)
	}

	customer := new(Customer)
	_, err = s.client.Do(req, customer)
	if err != nil {
		return *customer, err
	}

	return *customer, nil

}

func (s *CustomerService) getPage(p, ps int) ([]Customer, *Pagination, *Response, error) {
	url := fmt.Sprintf("customers?page=%v&page_size=%v", p, ps)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(CustomerResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Customers
	return resource, pagination, resp, err
}

// webhook (and API 1.0) specific structs

// WebhookCustomer is used to unmarshal payload of a webhook
type WebhookCustomer struct {
	ID               *string                 `json:"id,omitempty"`
	RetailerID       *string                 `json:"retailer_id,omitempty"`
	CustomerCode     *string                 `json:"customer_code,omitempty"`
	Balance          *string                 `json:"balance,omitempty"`
	Points           *int                    `json:"points,omitempty"`
	Note             *string                 `json:"note,omitempty"`
	YearToDate       *string                 `json:"year_to_date,omitempty"`
	Gender           *string                 `json:"sex,omitempty"`
	DateOfBirth      *string                 `json:"date_of_birth,omitempty"`
	CustomField1     *string                 `json:"custom_field_1,omitempty"`
	CustomField2     *string                 `json:"custom_field_2,omitempty"`
	CustomField3     *string                 `json:"custom_field_3,omitempty"`
	CustomField4     *string                 `json:"custom_field_4,omitempty"`
	UpdatedAt        *string                 `json:"updated_at,omitempty"`
	CreatedAt        *string                 `json:"created_at,omitempty"`
	DeletedAt        *string                 `json:"deleted_at,omitempty"`
	ContactFirstName *string                 `json:"contact_first_name,omitempty"`
	ContactLastName  *string                 `json:"contact_last_name,omitempty"`
	Contact          *WebhookCustomerContact `json:"contact,omitempty"`
}

// WebhookCustomerContact is used to unmarshal contact from the webhook customer payload
type WebhookCustomerContact struct {
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
	PhysicalCountryID *string `json:"physical_country_id,omitempty"`
	PostalAddress1    *string `json:"postal_address1,omitempty"`
	PostalAddress2    *string `json:"postal_address2,omitempty"`
	PostalSuburb      *string `json:"postal_suburb,omitempty"`
	PostalCity        *string `json:"postal_city,omitempty"`
	PostalPostcode    *string `json:"postal_postcode,omitempty"`
	PostalState       *string `json:"postal_state,omitempty"`
	PostalCountryID   *string `json:"postal_country_id,omitempty"`
}
