package model

type PurchaseOrder struct {
	id                      int
	number                  string
	planning_receiving_date string
	actual_receiving_time   string
	products                string
	status                  string
	cost                    int64
	vendor_id               int
}

func (c *PurchaseOrder) Create() {

}
func (c *PurchaseOrder) Read() {

}
func (c *PurchaseOrder) Update() {

}
func (c *PurchaseOrder) Delete() {

}
