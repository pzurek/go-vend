/*
{
  "id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
  "name": "Main Register",
  "regset_id": "6c8f04b3-3110-11e3-a29a-bc305bf5da20",
  "print_receipt": "0",
  "email_receipt": "0",
  "ask_for_note_on_save": "1",
  "print_note_on_receipt": "0",
  "ask_for_user_on_sale": "0",
  "show_discounts_on_receipt": "1",
  "receipt_header": "<h1>piotr<h1>",
  "receipt_barcoded": "1",
  "receipt_footer": "<h1>Thanks for stopping by<h1>",
  "receipt_style_class": "has-receipt-80",
  "invoice_prefix": "",
  "invoice_suffix": "",
  "invoice_sequence": 50,
  "register_open_count_sequence": "4",
  "register_open_time": "2013-09-23 23:30:42",
  "register_close_time": "",
  "quick_keys_template": {},
  "receipt": {
    "fields": {
      "label_invoice": "Invoice #:",
      "label_invoice_title": "Receipt / Tax Invoice",
      "label_served_by": "Served by:",
      "label_line_discount": "Less discount ",
      "label_sub_total": "Subtotal",
      "label_tax": "Tax",
      "label_to_pay": "TO PAY",
      "label_total": "TOTAL",
      "label_change": "Change",
      "header": "<h1>piotr<h1>",
      "footer": "<h1>Thanks for stopping by<h1>"
    }
  }
}
*/

package vend

import "fmt"

type Register struct {
	Id         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	RegisterId *string `json:"regset_id,omitempty"`
}

type RegisterResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Registers  *[]Register `json:"registers,omitempty"`
}

type RegisterService struct {
	client *Client
}

func (s *RegisterService) List() ([]Register, error) {

	resource := make([]Register, 0)

	regs, pagination, _, err := s.getPage(1, 50)

	if err != nil {
		return nil, err
	}

	resource = append(resource, *regs...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			regs, pg, _, err := s.getPage(*pagination.Page+1, 50)
			if err != nil {
				return nil, err
			}
			pagination = pg
			resource = append(resource, *regs...)
		}
	}

	return resource, err
}

func (s *RegisterService) getPage(p, ps int) (*[]Register, *Pagination, *Response, error) {
	u := fmt.Sprintf("registers?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(RegisterResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Registers
	return resource, pagination, resp, err
}
