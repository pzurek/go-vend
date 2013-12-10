/*
   {
     "id": "f2ad2a3c-8f2e-2325-0980-948c10f0dcf9",
     "payment_type_id": "1",
     "register_id": "02f65d7d-0adc-11e3-a415-bc764e10976c",
     "retailer_payment_type_id": "02f8c16f-0adc-11e3-a415-bc764e10976c",
     "name": "Cash",
     "label": "Cash",
     "payment_date": "2013-11-14 20:41:01",
     "amount": 299
   }
*/

package vce

type Payment struct {
	Id     *string  `json:"id,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
}
