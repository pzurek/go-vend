/*
{
  "outlet_id": "02f399e9-0adc-11e3-a415-bc764e10976c",
  "outlet_name": "Main Outlet",
  "count": "200.00000",
  "reorder_point": "45",
  "restock_level": "15"
}
*/

package vend

type InventoryItem struct {
	OutletId     *string `json:"outlet_id,omitempty"`
	OutletName   *string `json:"outlet_name,omitempty"`
	Count        *string `json:"count,omitempty"`
	ReorderPoint *string `json:"reorder_point,omitempty"`
	RestockLevel *string `json:"restock_level,omitempty"`
}
