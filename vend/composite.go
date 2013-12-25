/*

{
  "id": "756a4a75-dfcd-11e2-a415-bc764e10976c",
  "handle": "VinaZorzalGarnacha",
  "sku": "8414319002297",
  "count": "3.00000"
}

*/

package vend

type Composite struct {
	Id     *string `json:"id,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Sku    *string `json:"sku,omitempty"`
	Count  *string `json:"count,omitempty"`
}
