/*
{
  "pagination": {
    "results": 19,
    "page": 1,
    "page_size": 1,
    "pages": 19
  },
  "register_sales": [
    {
      "id": "87d6ec9a-572b-d8c1-5fd6-d0656c032d7c",
      "register_id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
      "market_id": "1",
      "customer_id": "02fb2229-0adc-11e3-a415-bc764e10976c",
      "customer_name": "",
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
      "user_id": "02f9bd97-0adc-11e3-a415-bc764e10976c",
      "user_name": "p.zurek@gmail.com",
      "sale_date": "2013-11-14 20:41:01",
      "created_at": "2013-11-14 20:41:02",
      "updated_at": "2013-11-14 20:41:02",
      "total_price": 260,
      "total_cost": 0,
      "total_tax": 39,
      "tax_name": "NZ GST",
      "note": "",
      "status": "CLOSED",
      "short_code": "of1yvp",
      "invoice_number": "45",
      "register_sale_products": [
        {
          "id": "89edaa26-d2b7-e417-0116-b118d7543084",
          "product_id": "7dbac57b-23be-11e3-a415-bc764e10976c",
          "register_id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
          "sequence": "0",
          "handle": "SUP635OT",
          "sku": "SUP635OT",
          "name": "America 49 Light Havana",
          "quantity": 1,
          "price": 260,
          "price_set": 0,
          "discount": 0,
          "loyalty_value": 0,
          "tax": 39,
          "tax_id": "41ab729c-44d4-11e3-a29a-bc305bf5da20",
          "tax_name": "NZ GST",
          "tax_rate": 0.15,
          "tax_total": 39,
          "price_total": 260,
          "display_retail_price_tax_inclusive": "1",
          "status": "CONFIRMED",
          "attributes": [
            {
              "name": "line_note",
              "value": ""
            }
          ]
        }
      ],
      "totals": {
        "total_tax": 39,
        "total_price": 260,
        "total_payment": 299,
        "total_to_pay": 0
      },
      "register_sale_payments": [
        {
          "id": "f2ad2a3c-8f2e-2325-0980-948c10f0dcf9",
          "payment_type_id": "1",
          "register_id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
          "retailer_payment_type_id": "02f8c16f-0adc-11e3-a415-bc764e10976c",
          "name": "Cash",
          "label": "Cash",
          "payment_date": "2013-11-14 20:41:01",
          "amount": 299
        }
      ],
      "taxes": [
        {
          "id": "622d3131-3562-2d34-3031-646231356665",
          "tax": 39,
          "name": "NZ GST",
          "rate": 0.15
        }
      ]
    }
  ]
}
*/

package vend

type Sale struct {
	Id            *string                `json:"id,omitempty"`
	RegisterId    *string                `json:"register_id,omitempty"`
	SaleDate      *string                `json:"sale_date,omitempty"`
	InvoiceNumber *string                `json:"invoice_number,omitempty"`
	Note          *string                `json:"note,omitempty"`
	CustomerName  *string                `json:"customer_name,omitempty"`
	Customer      *Customer              `json:"customer,omitempty"`
	UserName      *string                `json:"user_name,omitempty"`
	Products      *[]RegisterSaleProduct `json:"register_sale_products,omitempty"`
	Payments      *[]Payment             `json:"register_sale_payments,omitempty"`
	Totals        *Totals                `json:"totals,omitempty"`
	Status        *string                `json:"status,omitempty"`
}
