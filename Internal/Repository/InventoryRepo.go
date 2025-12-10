package Repository

import (
	"HOMEWORKAVENGINVAPI/Internal/Entity"

	_ "github.com/lib/pq"
)

// type InventoryStatus string

// const (
// 	StatusActive InventoryStatus = "Active"
// 	StatusBroken InventoryStatus = "Broken"
// )

// type Inventory struct {
// 	ID          int
// 	Name        string
// 	Itemcode    string
// 	Stock       string
// 	Description string
// 	Status      InventoryStatus
// }

type InventoryRepository struct {
	*BaseRepository
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *InventoryRepository) GetAllInventory() ([]Entity.Inventory, error) {
	query := `select * FROM avnginventory`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory []Entity.Inventory
	for rows.Next() {
		var inv Entity.Inventory
		err := rows.Scan(
			&inv.ID,
			&inv.Name,
			&inv.Itemcode,
			&inv.Stock,
			&inv.Description,
			&inv.Status,
		)
		if err != nil {
			return nil, err
		}
		inventory = append(inventory, inv)
	}

	return inventory, nil
}
