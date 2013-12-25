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
  "composites": [],
  "inventory": [],
  "price_book_entries": [],
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

type Product struct {
	Id                       *string           `json:"id,omitempty"`
	SourceId                 *string           `json:"source_id,omitempty"`
	VariantSourceId          *string           `json:"variant_source_id,omitempty"`
	Handle                   *string           `json:"handle,omitempty"`
	Type                     *string           `json:"type,omitempty"`
	Active                   *bool             `json:"active,omitempty"`
	Name                     *string           `json:"name,omitempty"`
	Description              *string           `json:"description,omitempty"`
	Image                    *string           `json:"image,omitempty"`
	ImageLarge               *string           `json:"image_large,omitempty"`
	Sku                      *string           `json:"sku,omitempty"`
	Tags                     *string           `json:"tags,omitempty"`
	BrandId                  *string           `json:"brand_id,omitempty"`
	BrandName                *string           `json:"brand_name,omitempty"`
	SupplierName             *string           `json:"supplier_name,omitempty"`
	SupplierCode             *string           `json:"supplier_code,omitempty"`
	SupplyPrice              *string           `json:"supply_price,omitempty"`
	AccountCodePurchase      *string           `json:"account_code_purchase,omitempty"`
	AccountCodeSales         *string           `json:"account_code_sales,omitempty"`
	Inventory                *[]InventoryItem  `json:"inventory,omitempty"`
	Composites               *[]Composite      `json:"composites,omitempty"`
	PriceBookEntries         *[]PriceBookEntry `json:"price_book_entries,omitempty"`
	Price                    *float64          `json:"price,omitempty"`
	Tax                      *float64          `json:"tax,omitempty"`
	TaxId                    *string           `json:"tax_id,omitempty"`
	TaxRate                  *float64          `json:"tax_rate,omitempty"`
	TaxName                  *string           `json:"tax_name,omitempty"`
	DisplayPriceTaxInclusive *int              `json:"display_retail_price_tax_inclusive,omitempty"`
	UpdatedAt                *string           `json:"updated_at,omitempty"`
}

type ProductResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Products   *[]Product  `json:"products,omitempty"`
}

type ProductService struct {
	client *Client
}

func (s *ProductService) List() ([]Product, error) {

	products := make([]Product, 0)

	prod, pagination, _, err := s.getProductPage(1, 200)

	if err != nil {
		return nil, err
	}

	products = append(products, *prod...)

	if pagination != nil {
		for *pagination.Page < *pagination.Pages {
			prod, pag, _, err := s.getProductPage(*pagination.Page+1, 200)
			if err != nil {
				return nil, err
			}
			pagination = pag
			products = append(products, *prod...)
		}
	}

	return products, err
}

func (s *ProductService) getProductPage(p, ps int) (*[]Product, *Pagination, *Response, error) {
	u := fmt.Sprintf("products?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(ProductResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	products := response.Products
	return products, pagination, resp, err
}
