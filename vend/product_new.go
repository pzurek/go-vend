/*

{
  "id": "add3ff78-587e-11e3-a29a-bc305bf5da20",
  "source_id": "185134285",
  "variant_source_id": "423922789",
  "handle": "new-fancy-product",
  "type": "test",
  "variant_parent_id": "",
  "variant_option_one_name": "",
  "variant_option_one_value": "",
  "variant_option_two_name": "",
  "variant_option_two_value": "",
  "variant_option_three_name": "",
  "variant_option_three_value": "",
  "active": true,
  "name": "New fancy product",
  "description": "<p>Description from Vend</p>",
  "image": "https://piotr.vendhq.com/images/placeholder/product/no-image-white-thumb.png",
  "image_large": "https://piotr.vendhq.com/images/placeholder/product/no-image-white-original.png",
  "sku": "1000000",
  "tags": "",
  "brand_id": "",
  "brand_name": "",
  "supplier_name": "Shopify",
  "supplier_code": "",
  "supply_price": "0.00",
  "account_code_purchase": "",
  "account_code_sales": "",
  "composites": [
    {
      "id": "756a4a75-dfcd-11e2-a415-bc764e10976c",
      "handle": "VinaZorzalGarnacha",
      "sku": "8414319002297",
      "count": "3.00000"
    }
  ],
  "inventory": [
    {
      "outlet_id": "02f399e9-0adc-11e3-a415-bc764e10976c",
      "outlet_name": "Main Outlet",
      "count": "200.00000",
      "reorder_point": "45",
      "restock_level": "15"
    }
  ],
  "price_book_entries": [
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
  ],
  "price": 675.65,
  "tax": 101.35,
  "tax_id": "41ab729c-44d4-11e3-a29a-bc305bf5da20",
  "tax_rate": 0.15,
  "tax_name": "NZ GST",
  "display_retail_price_tax_inclusive": 1,
  "updated_at": "2013-12-05 09:00:55",
  "deleted_at": ""
}

*/

package vend

import (
	"fmt"
)

type NewProduct struct {
	Id                      *string           `json:"id,omitempty"`
	SourceId                *string           `json:"source_id,omitempty"`
	VariantSourceId         *string           `json:"variant_source_id,omitempty"`
	Handle                  *string           `json:"handle,omitempty"`
	Type                    *string           `json:"type,omitempty"`
	VariantParentId         *string           `json:"variant_parent_id"`
	VariantOptionOneName    *string           `json:"variant_option_one_name"`
	VariantOptionOneValue   *string           `json:"variant_option_one_value"`
	VariantOptionTwoName    *string           `json:"variant_option_two_name"`
	VariantOptionTwoValue   *string           `json:"variant_option_two_value"`
	VariantOptionThreeName  *string           `json:"variant_option_three_name"`
	VariantOptionThreeValue *string           `json:"variant_option_three_value"`
	Active                  *bool             `json:"active,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	Description             *string           `json:"description,omitempty"`
	Image                   *string           `json:"image,omitempty"`
	ImageLarge              *string           `json:"image_large,omitempty"`
	Sku                     *string           `json:"sku,omitempty"`
	Tags                    *string           `json:"tags,omitempty"`
	BrandId                 *string           `json:"brand_id,omitempty"`
	BrandName               *string           `json:"brand_name,omitempty"`
	SupplierName            *string           `json:"supplier_name,omitempty"`
	SupplierCode            *string           `json:"supplier_code,omitempty"`
	SupplyPrice             *string           `json:"supply_price,omitempty"`
	AccountCodePurchase     *string           `json:"account_code_purchase,omitempty"`
	AccountCodeSales        *string           `json:"account_code_sales,omitempty"`
	Inventory               *[]InventoryItem  `json:"inventory,omitempty"`
	Composites              *[]Composite      `json:"composites,omitempty"`
	PriceBookEntries        *[]PriceBookEntry `json:"price_book_entries,omitempty"`
	Price                   *float64          `json:"price,omitempty"`
	Tax                     *float64          `json:"tax,omitempty"`
	TaxId                   *string           `json:"tax_id,omitempty"`
	TaxRate                 *float64          `json:"tax_rate,omitempty"`
	TaxName                 *string           `json:"tax_name,omitempty"`
	TaxInclusive            *int              `json:"display_retail_price_tax_inclusive,omitempty"`
	SequenceId              *int              `json:"sequence_id"`
	UpdatedAt               *string           `json:"updated_at,omitempty"`
}

type NewProductResponse struct {
	Products *[]NewProduct `json:"products,omitempty"`
}

type NewProductService struct {
	client *Client
}

// List returns a slice of all products
func (s *NewProductService) List() ([]NewProduct, error) {

	resource := make([]NewProduct, 0)

	rp, sequence, _, err := s.getPage(0)

	if err != nil {
		return nil, err
	}

	lastPageSize := len(*rp)
	resource = append(resource, *rp...)

	for lastPageSize > 0 {
		rp, sq, _, err := s.getPage(sequence)
		if err != nil {
			return nil, err
		}
		lastPageSize = len(*rp)
		sequence = sq
		resource = append(resource, *rp...)
	}

	return resource, err
}

func (s *NewProductService) getPage(ps int) (*[]NewProduct, int, *Response, error) {
	u := fmt.Sprintf("products?after=%v&order_by=sequence_id&order_direction=asc", ps)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	response := new(NewProductResponse)

	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, 0, resp, err
	}

	resource := response.Products
	sequence := ps

	for _, product := range *resource {
		if *product.SequenceId > sequence {
			sequence = *product.SequenceId
		}
	}

	return resource, sequence, resp, err
}
