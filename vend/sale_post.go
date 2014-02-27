/*
{
    "register_id": "6e1c67ff-80e4-11df-b0bf-4040f540b50a",
    "customer_id": "null",
    "sale_date": "2010-10-15 05:39:15",
    "user_name": "admin",
    "total_price": 4248.16,
    "total_tax": 637.22,
    "tax_name": "NZ GST",
    "status": "CLOSED",
    "invoice_number": "25",
    "invoice_sequence": 25,
    "note": null,
    "register_sale_products": [
    {
        "product_id": "0515649e-80e5-11df-b0bf-4040f540b50a",
        "quantity": 1,
        "price": 1999.13,
        "tax": 299.87,
        "tax_id": "53b3501c-887c-102d-8a4b-a9cf13f17faa",
        "tax_total": 299.87
    },
    {
        "product_id": "507d497e-80e5-11df-b0bf-4040f540b50a",
        "quantity": 1,
        "price": 2249.03,
        "tax": 337.35,
        "tax_id": "53b3501c-887c-102d-8a4b-a9cf13f17faa",
        "tax_total": 337.35,
        "attributes" : [
                {
                    "name"  : "line_note",
                    "value" : "This is a line item note for a single product"
                }
        ]
    }],
    "register_sale_payments": [
    {
        "retailer_payment_type_id": "6e303e1e-80e4-11df-b0bf-4040f540b50a",
        "payment_date": "2010-10-15 05:39:13",
        "amount": 1000
    },
    {
        "retailer_payment_type_id": "a689e6de-80e4-11df-b0bf-4040f540b50a",
        "payment_date": "2010-10-15 05:39:15",
        "amount": 3885.38
    }]
}
*/

package vend

type SalePost struct {
	RegisterId      string            `json:"register_id,omitempty"`
	MarketId        string            `json:"market_id,omitempty"`
	CustomerId      string            `json:"customer_id,omitempty"`
	SaleDate        string            `json:"sale_date,omitempty"`
	UserName        string            `json:"user_name,omitempty"`
	TotalPrice      float64           `json:"total_price,omitempty"`
	TotalTax        float64           `json:"total_tax,omitempty"`
	TaxName         string            `json:"tax_name,omitempty"`
	Status          string            `json:"status,omitempty"`
	InvoiceNumber   string            `json:"invoice_number,omitempty"`
	InvoiceSequence int               `json:"invoice_sequence,omitempty"`
	Note            string            `json:"note,omitempty"`
	Products        []SaleProductPost `json:"register_sale_products,omitempty"`
	Payments        []SalePaymentPost `json:"register_sale_payments,omitempty"`
}

type SaleProductPost struct {
	ProductId  string      `json:"product_id,omitempty"`
	Quantity   float64     `json:"quantity,omitempty"`
	Price      float64     `json:"price,omitempty"`
	Tax        float64     `json:"tax,omitempty"`
	TaxId      string      `json:"tax_id,omitempty"`
	TaxTotal   float64     `json:"tax_total,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type SalePaymentPost struct {
	PaymentTypeId string  `json:"retailer_payment_type_id,omitempty"`
	PaymentDate   string  `json:"payment_date,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
}
