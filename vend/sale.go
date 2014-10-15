package vend

type Sale struct {
	Id            *string               `json:"id,omitempty"`
	RegisterId    *string               `json:"register_id,omitempty"`
	SaleDate      *string               `json:"sale_date,omitempty"`
	InvoiceNumber *string               `json:"invoice_number,omitempty"`
	Note          *string               `json:"note,omitempty"`
	CustomerName  *string               `json:"customer_name,omitempty"`
	Customer      *Customer             `json:"customer,omitempty"`
	UserName      *string               `json:"user_name,omitempty"`
	Products      []RegisterSaleProduct `json:"register_sale_products,omitempty"`
	Payments      []Payment             `json:"register_sale_payments,omitempty"`
	Totals        *Totals               `json:"totals,omitempty"`
	Status        *string               `json:"status,omitempty"`
}

type RegisterSaleProduct struct {
	Id         *string     `json:"id,omitempty"`
	Name       *string     `json:"name,omitempty"`
	Sku        *string     `json:"sku,omitempty"`
	Quantity   *float64    `json:"quantity,omitempty"`
	Price      *float64    `json:"price,omitempty"`
	Tax        *float64    `json:"tax,omitempty"`
	Discount   *float64    `json:"discount,omitempty"`
	Loyalty    *float64    `json:"loyalty_value,omitempty"`
	TaxTotal   *float64    `json:"tax_total,omitempty"`
	PriceTotal *float64    `json:"price_total,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Payment struct {
	Id     *string  `json:"id,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
}

type Totals struct {
	Tax   *float64 `json:"total_tax,omitempty"`
	Price *float64 `json:"total_price,omitempty"`
	Paid  *float64 `json:"total_payment,omitempty"`
	ToPay *float64 `json:"total_to_pay,omitempty"`
}
