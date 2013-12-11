package vend

type Config struct {
	RetailerId                     *string     `json:"retailer_id,omitempty"`
	RetailerName                   *string     `json:"retailer_name,omitempty"`
	AccountState                   *string     `json:"account_state,omitempty"`
	DomainPrefix                   *string     `json:"domain_prefix,omitempty"`
	DisplayRetailPriceTaxInclusive *string     `json:"display_retail_price_tax_inclusive,omitempty"`
	UserId                         *string     `json:"user_id,omitempty"`
	UserName                       *string     `json:"user_name,omitempty"`
	UserHash                       *string     `json:"user_hash,omitempty"`
	UserDisplayName                *string     `json:"user_display_name,omitempty"`
	NotifyUserOfNumpad             *bool       `json:"notify_user_of_numpad,omitempty"`
	AccountType                    *string     `json:"account_type,omitempty"`
	OutletName                     *string     `json:"outlet_name,omitempty"`
	OutletId                       *string     `json:"outlet_id,omitempty"`
	Version                        *string     `json:"version,omitempty"`
	Migrations                     *Migrations `json:"migrate_sql,omitempty"`
	LastSync                       *string     `json:"last_sync,omitempty"`
	CurrencyName                   *string     `json:"currency_name,omitempty"`
	Currency                       *string     `json:"currency,omitempty"`
	Culture                        *string     `json:"culture,omitempty"`
	DefaultCustomerGroupId         *string     `json:"default_customer_group_id,omitempty"`
	DefaultCustomerId              *string     `json:"default_customer_id,omitempty"`
	DiscountProductId              *string     `json:"discount_product_id,omitempty"`
	CashierDiscount                *bool       `json:"cashier_discount,omitempty"`
	EnableLoyalty                  *int        `json:"enable_loyalty,omitempty"`
	// TODO: Delete it or fix it.
	//Callbacks                      *Callbacks  `json:"callbacks,omitempty"`
}

type ConfigResponse struct {
	Config *Config `json:"config"`
}

type ConfigService struct {
	client *Client
}
