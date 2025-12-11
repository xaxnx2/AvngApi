package Entity

type InventoryStatus string

type Inventory struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Itemcode    string `json:"ItemCode"`
	Stock       int    `json:"Stock"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
}
