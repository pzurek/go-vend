/*
 "totals": {
        "total_tax": 39,
        "total_price": 260,
        "total_payment": 299,
        "total_to_pay": 0
      }
*/

package vend

type Totals struct {
	Tax   *float64 `json:"total_tax,omitempty"`
	Price *float64 `json:"total_price,omitempty"`
	Paid  *float64 `json:"total_payment,omitempty"`
	ToPay *float64 `json:"total_to_pay,omitempty"`
}
