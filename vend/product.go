package vend

import (
	"fmt"
)

type Product struct {
	Id                       *string          `json:"id,omitempty"`
	SourceId                 *string          `json:"source_id,omitempty"`
	VariantSourceId          *string          `json:"variant_source_id,omitempty"`
	Handle                   *string          `json:"handle,omitempty"`
	Type                     *string          `json:"type,omitempty"`
	Active                   *bool            `json:"active,omitempty"`
	Name                     *string          `json:"name,omitempty"`
	Description              *string          `json:"description,omitempty"`
	Image                    *string          `json:"image,omitempty"`
	ImageLarge               *string          `json:"image_large,omitempty"`
	Sku                      *string          `json:"sku,omitempty"`
	Tags                     *string          `json:"tags,omitempty"`
	BrandId                  *string          `json:"brand_id,omitempty"`
	BrandName                *string          `json:"brand_name,omitempty"`
	SupplierName             *string          `json:"supplier_name,omitempty"`
	SupplierCode             *string          `json:"supplier_code,omitempty"`
	SupplyPrice              *string          `json:"supply_price,omitempty"`
	AccountCodePurchase      *string          `json:"account_code_purchase,omitempty"`
	AccountCodeSales         *string          `json:"account_code_sales,omitempty"`
	Inventory                []InventoryItem  `json:"inventory,omitempty"`
	Composites               []Composite      `json:"composites,omitempty"`
	PriceBookEntries         []PriceBookEntry `json:"price_book_entries,omitempty"`
	Price                    *float64         `json:"price,omitempty"`
	Tax                      *float64         `json:"tax,omitempty"`
	TaxId                    *string          `json:"tax_id,omitempty"`
	TaxRate                  *float64         `json:"tax_rate,omitempty"`
	TaxName                  *string          `json:"tax_name,omitempty"`
	DisplayPriceTaxInclusive *int             `json:"display_retail_price_tax_inclusive,omitempty"`
	UpdatedAt                *string          `json:"updated_at,omitempty"`
}

type InventoryItem struct {
	OutletId     *string `json:"outlet_id,omitempty"`
	OutletName   *string `json:"outlet_name,omitempty"`
	Count        *string `json:"count,omitempty"`
	ReorderPoint *string `json:"reorder_point,omitempty"`
	RestockLevel *string `json:"restock_level,omitempty"`
}

type Composite struct {
	Id     *string `json:"id,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Sku    *string `json:"sku,omitempty"`
	Count  *string `json:"count,omitempty"`
}

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

type ProductResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Products   []Product   `json:"products,omitempty"`
}

type ProductService struct {
	client *Client
}

// List returns a slice of all products
func (s *ProductService) List() ([]Product, error) {

	var resource []Product

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

func (s *ProductService) getPage(p, ps int) ([]Product, *Pagination, *Response, error) {
	url := fmt.Sprintf("products?page=%v&page_size=%v&sort_by=id", p, ps)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(ProductResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	pagination := response.Pagination
	resource := response.Products
	return resource, pagination, resp, err
}
