package Entity

type InventoryStatus string

const (
	Active InventoryStatus = "Active"
	Broken InventoryStatus = "Broken"
)

type Inventory struct {
	ID          int             `json:"ID"`
	Name        string          `json:"Name"`
	Itemcode    string          `json:"ItemCode"`
	Stock       int             `json:"Stock"`
	Description string          `json:"Description"`
	Status      InventoryStatus `json:"Status"`
}
