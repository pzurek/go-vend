/*
{
  "id": "3ec35710-3aae-a502-cd6b-9e8539c9eccf",
  "product_id": "add3ff78-587e-11e3-a29a-bc305bf5da20",
  "price_book_id": "02ff5bf1-0adc-11e3-a415-bc764e10976c",
  "price_book_name": "General Price Book (All Products)",
  "type": "BASE",
  "outlet_name": "",
  "outlet_id": "",
  "customer_group_name": "All Customers",
  "customer_group_id": "02fa2d73-0adc-11e3-a415-bc764e10976c",
  "price": 675.65,
  "loyalty_value": null,
  "tax": 101.35,
  "tax_id": "41ab729c-44d4-11e3-a29a-bc305bf5da20",
  "tax_rate": 0.15,
  "tax_name": "NZ GST",
  "display_retail_price_tax_inclusive": 1,
  "min_units": "",
  "max_units": "",
  "valid_from": "",
  "valid_to": ""
}
*/

package vend

type PriceBookEntry struct {
	Id           *string  `json:"id,omitempty"`
	ProductId    *string  `json:"product_id,omitempty"`
	PriceBookId  *string  `json:"price_book_id,omitempty"`
	Type         *string  `json:"type,omitempty"`
	OutletName   *string  `json:"outlet_name,omitempty"`
	OutletId     *string  `json:"outlet_id,omitempty"`
	Price        *float64 `json:"price,omitempty"`
	LoyaltyValue *float64 `json:"loyalty_value,omitempty"`
	Tax          *float64 `json:"tax,omitempty"`
	TaxId        *string  `json:"tax_id,omitempty"`
	TaxRate      *float64 `json:"tax_rate,omitempty"`
	TaxName      *string  `json:"tax_name,omitempty"`
	TaxInclusive *int     `json:"display_retail_price_tax_inclusive,omitempty"`
	MinUnits     *string  `json:"min_units,omitempty"`
	MaxUnits     *string  `json:"max_units,omitempty"`
	ValidFrom    *string  `json:"valid_from,omitempty"`
	ValidTo      *string  `json:"valid_to,omitempty"`
}
