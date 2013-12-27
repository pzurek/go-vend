
# vend
    import "github.com/pzurek/go-vend/vend"






## func CheckResponse
``` go
func CheckResponse(r *http.Response) error
```


## type Attribute
``` go
type Attribute struct {
    Name  *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}
```










## type Callback
``` go
type Callback struct{}
```










## type Client
``` go
type Client struct {
    UserAgent string

    Config    *ConfigService
    Products  *ProductService
    Taxes     *TaxService
    Users     *UserService
    Outlets   *OutletService
    Registers *RegisterService
    Suppliers *SupplierService
    Customers *CustomerService
    // contains filtered or unexported fields
}
```
Vend client which talks to the Vend API









### func NewClient
``` go
func NewClient(st, usrname, passwd string, httpClient *http.Client) *Client
```



### func (\*Client) Do
``` go
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error)
```


### func (\*Client) NewRequest
``` go
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error)
```


## type Composite
``` go
type Composite struct {
    Id     *string `json:"id,omitempty"`
    Handle *string `json:"handle,omitempty"`
    Sku    *string `json:"sku,omitempty"`
    Count  *string `json:"count,omitempty"`
}
```










## type Config
``` go
type Config struct {
    RetailerId                     *string      `json:"retailer_id,omitempty"`
    RetailerName                   *string      `json:"retailer_name,omitempty"`
    AccountState                   *string      `json:"account_state,omitempty"`
    DomainPrefix                   *string      `json:"domain_prefix,omitempty"`
    DisplayRetailPriceTaxInclusive *string      `json:"display_retail_price_tax_inclusive,omitempty"`
    UserId                         *string      `json:"user_id,omitempty"`
    UserName                       *string      `json:"user_name,omitempty"`
    UserHash                       *string      `json:"user_hash,omitempty"`
    UserDisplayName                *string      `json:"user_display_name,omitempty"`
    NotifyUserOfNumpad             *bool        `json:"notify_user_of_numpad,omitempty"`
    AccountType                    *string      `json:"account_type,omitempty"`
    OutletName                     *string      `json:"outlet_name,omitempty"`
    OutletId                       *string      `json:"outlet_id,omitempty"`
    Version                        *string      `json:"version,omitempty"`
    Migrations                     *[]Migration `json:"migrate_sql,omitempty"`
    LastSync                       *string      `json:"last_sync,omitempty"`
    CurrencyName                   *string      `json:"currency_name,omitempty"`
    Currency                       *string      `json:"currency,omitempty"`
    Culture                        *string      `json:"culture,omitempty"`
    DefaultCustomerGroupId         *string      `json:"default_customer_group_id,omitempty"`
    DefaultCustomerId              *string      `json:"default_customer_id,omitempty"`
    DiscountProductId              *string      `json:"discount_product_id,omitempty"`
    CashierDiscount                *bool        `json:"cashier_discount,omitempty"`
    EnableLoyalty                  *int         `json:"enable_loyalty,omitempty"`
}
```










## type ConfigResponse
``` go
type ConfigResponse struct {
    Config *Config `json:"config"`
}
```










## type ConfigService
``` go
type ConfigService struct {
    // contains filtered or unexported fields
}
```










## type Contact
``` go
type Contact struct {
    FirstName         *string `json:"first_name,omitempty"`
    LastName          *string `json:"last_name,omitempty"`
    CompanyName       *string `json:"company_name,omitempty"`
    Phone             *string `json:"phone,omitempty"`
    Mobile            *string `json:"mobile,omitempty"`
    Fax               *string `json:"fax,omitempty"`
    Email             *string `json:"email,omitempty"`
    Twitter           *string `json:"twitter,omitempty"`
    Website           *string `json:"website,omitempty"`
    PhysicalAddress1  *string `json:"physical_address1,omitempty"`
    PhysicalAddress2  *string `json:"physical_address2,omitempty"`
    PhysicalSuburb    *string `json:"physical_suburb,omitempty"`
    PhysicalCity      *string `json:"physical_city,omitempty"`
    PhysicalPostcode  *string `json:"physical_postcode,omitempty"`
    PhysicalState     *string `json:"physical_state,omitempty"`
    PhysicalCountryId *string `json:"physical_country_id,omitempty"`
    PostalAddress1    *string `json:"postal_address1,omitempty"`
    PostalAddress2    *string `json:"postal_address2,omitempty"`
    PostalSuburb      *string `json:"postal_suburb,omitempty"`
    PostalCity        *string `json:"postal_city,omitempty"`
    PostalPostcode    *string `json:"postal_postcode,omitempty"`
    PostalState       *string `json:"postal_state,omitempty"`
    PostalCountryId   *string `json:"postal_country_id,omitempty"`
}
```
Generic Contact type used by other types











## type Customer
``` go
type Customer struct {
    Id                *string          `json:"id,omitempty"`
    Name              *string          `json:"name,omitempty"`
    CustomerCode      *string          `json:"customer_code,omitempty"`
    CustomerGroupId   *string          `json:"customer_group_id,omitempty"`
    CustomerGroupName *string          `json:"customer_group_name,omitempty"`
    FirstName         *string          `json:"first_name,omitempty"`
    LastName          *string          `json:"last_name,omitempty"`
    CompanyName       *string          `json:"company_name,omitempty"`
    Phone             *string          `json:"phone,omitempty"`
    Mobile            *string          `json:"mobile,omitempty"`
    Fax               *string          `json:"fax,omitempty"`
    Email             *string          `json:"email,omitempty"`
    Twitter           *string          `json:"twitter,omitempty"`
    Website           *string          `json:"website,omitempty"`
    PhysicalAddress1  *string          `json:"physical_address1,omitempty"`
    PhysicalAddress2  *string          `json:"physical_address2,omitempty"`
    PhysicalSuburb    *string          `json:"physical_suburb,omitempty"`
    PhysicalCity      *string          `json:"physical_city,omitempty"`
    PhysicalPostcode  *string          `json:"physical_postcode,omitempty"`
    PhysicalState     *string          `json:"physical_state,omitempty"`
    PhysicalCountryId *string          `json:"physical_country_id,omitempty"`
    PostalAddress1    *string          `json:"postal_address1,omitempty"`
    PostalAddress2    *string          `json:"postal_address2,omitempty"`
    PostalSuburb      *string          `json:"postal_suburb,omitempty"`
    PostalCity        *string          `json:"postal_city,omitempty"`
    PostalPostcode    *string          `json:"postal_postcode,omitempty"`
    PostalState       *string          `json:"postal_state,omitempty"`
    PostalCountryId   *string          `json:"postal_country_id,omitempty"`
    EnableLoyalty     *int             `json:"enable_loayaly,omitempty"`
    LoyaltyBalance    *string          `json:"loyalty_balance,omitempty"`
    UpdatedAt         *string          `json:"updated_at,omitempty"`
    DeletedAt         *string          `json:"deleted_at,omitempty"`
    Balance           *string          `json:"balance,omitempty"`
    YearToDate        *string          `json:"year_to_date,omitempty"`
    DateOfBirth       *string          `json:"date_of_birth,omitempty"`
    Sex               *string          `json:"sex,omitempty"`
    CustomField1      *string          `json:"custom_field_1,omitempty"`
    CustomField2      *string          `json:"custom_field_2,omitempty"`
    CustomField3      *string          `json:"custom_field_3,omitempty"`
    CustomField4      *string          `json:"custom_field_4,omitempty"`
    Contact           *CustomerContact `json:"contact,omitempty"`
}
```










## type CustomerContact
``` go
type CustomerContact struct {
    CompanyName *string `json:"company_name,omitempty"`
    Phone       *string `json:"phone,omitempty"`
    Email       *string `json:"email,omitempty"`
}
```
Contact type used by the Customer type











## type CustomerResponse
``` go
type CustomerResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Customers  *[]Customer `json:"customers,omitempty"`
}
```










## type CustomerService
``` go
type CustomerService struct {
    // contains filtered or unexported fields
}
```










### func (\*CustomerService) List
``` go
func (s *CustomerService) List() ([]Customer, error)
```


## type ErrorResponse
``` go
type ErrorResponse struct {
    Response *http.Response
    Message  string `json:"error,omitempty"`
}
```










### func (\*ErrorResponse) Error
``` go
func (r *ErrorResponse) Error() string
```


## type InventoryItem
``` go
type InventoryItem struct {
    OutletId     *string `json:"outlet_id,omitempty"`
    OutletName   *string `json:"outlet_name,omitempty"`
    Count        *string `json:"count,omitempty"`
    ReorderPoint *string `json:"reorder_point,omitempty"`
    RestockLevel *string `json:"restock_level,omitempty"`
}
```










## type Migration
``` go
type Migration struct {
    Version *string `json:"version"`
    Table   *string `json:"table"`
    Query   *string `json:"query"`
}
```










## type Outlet
``` go
type Outlet struct {
    Id                *string  `json:"id,omitempty"`
    RetailerId        *string  `json:"retailer_id,omitempty"`
    Name              *string  `json:"name,omitempty"`
    TimeZone          *string  `json:"timezone,omitempty"`
    Contact           *Contact `json:"contact,omitempty"`
    Email             *string  `json:"email,omitempty"`
    PhysicalAddress1  *string  `json:"physical_address1,omitempty"`
    PhysicalAddress2  *string  `json:"physical_address2,omitempty"`
    PhysicalSuburb    *string  `json:"physical_suburb,omitempty"`
    PhysicalCity      *string  `json:"physical_city,omitempty"`
    PhysicalPostcode  *string  `json:"physical_postcode,omitempty"`
    PhysicalState     *string  `json:"physical_state,omitempty"`
    PhysicalCountryId *string  `json:"physical_country_id,omitempty"`
}
```










## type OutletResponse
``` go
type OutletResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Outlets    *[]Outlet   `json:"outlets,omitempty"`
}
```










## type OutletService
``` go
type OutletService struct {
    // contains filtered or unexported fields
}
```










### func (\*OutletService) List
``` go
func (s *OutletService) List() ([]Outlet, error)
```


## type Pagination
``` go
type Pagination struct {
    Results  *int `json:"results"`
    Page     *int `json:"page"`
    PageSize *int `json:"page_size"`
    Pages    *int `json:"pages"`
}
```










## type Payment
``` go
type Payment struct {
    Id     *string  `json:"id,omitempty"`
    Name   *string  `json:"name,omitempty"`
    Amount *float64 `json:"amount,omitempty"`
}
```










## type PriceBookEntry
``` go
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
```










## type Product
``` go
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
```










## type ProductResponse
``` go
type ProductResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Products   *[]Product  `json:"products,omitempty"`
}
```










## type ProductService
``` go
type ProductService struct {
    // contains filtered or unexported fields
}
```










### func (\*ProductService) List
``` go
func (s *ProductService) List() ([]Product, error)
```


## type Register
``` go
type Register struct {
    Id         *string `json:"id,omitempty"`
    Name       *string `json:"name,omitempty"`
    RegisterId *string `json:"regset_id,omitempty"`
}
```










## type RegisterResponse
``` go
type RegisterResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Registers  *[]Register `json:"registers,omitempty"`
}
```










## type RegisterSaleProduct
``` go
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
```










## type RegisterService
``` go
type RegisterService struct {
    // contains filtered or unexported fields
}
```










### func (\*RegisterService) List
``` go
func (s *RegisterService) List() ([]Register, error)
```


## type Response
``` go
type Response struct {
    *http.Response
}
```










## type Sale
``` go
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
```










## type Supplier
``` go
type Supplier struct {
    Id          *string  `json:"id,omitempty"`
    RetailerId  *string  `json:"retailer_id,omitempty"`
    Name        *string  `json:"name,omitempty"`
    Description *string  `json:"description,omitempty"`
    Contact     *Contact `json:"contact,omitempty"`
}
```










## type SupplierResponse
``` go
type SupplierResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Suppliers  *[]Supplier `json:"suppliers,omitempty"`
}
```










## type SupplierService
``` go
type SupplierService struct {
    // contains filtered or unexported fields
}
```










### func (\*SupplierService) List
``` go
func (s *SupplierService) List() ([]Supplier, error)
```


## type Tax
``` go
type Tax struct {
    Id      *string    `json:"id,omitempty"`
    Name    *string    `json:"tax,omitempty"`
    Rate    *float64   `json:"tax_rate,omitempty"`
    Default *bool      `json:"default,omitempty"`
    Rates   *[]TaxRate `json:"rates,omitempty"`
}
```










## type TaxRate
``` go
type TaxRate struct {
    Id   *string  `json:"id,omitempty"`
    Name *string  `json:"tax,omitempty"`
    Rate *float64 `json:"tax_rate,omitempty"`
}
```










## type TaxResponse
``` go
type TaxResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Taxes      *[]Tax      `json:"taxes,omitempty"`
}
```










## type TaxService
``` go
type TaxService struct {
    // contains filtered or unexported fields
}
```










### func (\*TaxService) List
``` go
func (s *TaxService) List() ([]Tax, error)
```


## type Totals
``` go
type Totals struct {
    Tax   *float64 `json:"total_tax,omitempty"`
    Price *float64 `json:"total_price,omitempty"`
    Paid  *float64 `json:"total_payment,omitempty"`
    ToPay *float64 `json:"total_to_pay,omitempty"`
}
```










## type User
``` go
type User struct {
    Id          *string `json:"id,omitempty"`
    Username    *string `json:"username,omitempty"`
    Name        *string `json:"name,omitempty"`
    OutletId    *string `json:"outlet_id,omitempty"`
    OutletName  *string `json:"outlet_,omitempty"`
    AccountType *string `json:"account_type,omitempty"`
}
```










## type UserResponse
``` go
type UserResponse struct {
    Pagination *Pagination `json:"pagination,omitempty"`
    Users      *[]User     `json:"users,omitempty"`
}
```










## type UserService
``` go
type UserService struct {
    // contains filtered or unexported fields
}
```










### func (\*UserService) List
``` go
func (s *UserService) List() ([]User, error)
```








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)