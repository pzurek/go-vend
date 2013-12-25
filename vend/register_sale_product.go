/*

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

*/

package vend

type RegisterSaleProduct struct {
	Id         *string      `json:"id,omitempty"`
	Name       *string      `json:"name,omitempty"`
	Sku        *string      `json:"sku,omitempty"`
	Quantity   *float64     `json:"quantity,omitempty"`
	Price      *float64     `json:"price,omitempty"`
	Tax        *float64     `json:"tax,omitempty"`
	Discount   *float64     `json:"discount,omitempty"`
	Loyalty    *float64     `json:"loyalty_value,omitempty"`
	TaxTotal   *float64     `json:"tax_total,omitempty"`
	PriceTotal *float64     `json:"price_total,omitempty"`
	Attributes *[]Attribute `json:"attributes,omitempty"`
}
