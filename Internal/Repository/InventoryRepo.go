package Repository

import (
	"HOMEWORKAVENGINVAPI/Internal/Entity"

	_ "github.com/lib/pq"
)

type InventoryRepository struct {
	*BaseRepository
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{
		BaseRepository: NewBaseRepository(),
	}
}

func (r *InventoryRepository) CreateInventory(name string, itemcode string, stock int, description string, status string) (*Entity.Inventory, error) {
	query := `
	INSERT INTO avnginventory (
		name,
		itemcode,
		stock,
		description,
		status
	)
	VALUES ($1,$2,$3,$4,$5)
	`
	var inv Entity.Inventory
	_, err := r.db.Exec(query, name, itemcode, stock, description, status)
	if err != nil {
		return nil, err
	}

	return &inv, nil
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

func (r *InventoryRepository) GetInventoryByID(ID int) (*Entity.Inventory, error) {
	query := `
	select * FROM 
	avnginventory
	WHERE id = $1
	`
	var inv Entity.Inventory
	err := r.db.QueryRow(query, ID).Scan(
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

	return &inv, nil
}

func (r *InventoryRepository) UpdateInventory(id int, name string, itemcode string, stock int, description string, status string) (*Entity.Inventory, error) {
	query := `
	UPDATE avnginventory
	SET 
	name = $1,
	itemcode = $2,
	stock = $3,
	description = $4,
	status $5
	WHERE id = $6
	`
	var inv Entity.Inventory
	_, err := r.db.Exec(query, name, itemcode, stock, description, status, id)
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func (r *InventoryRepository) DeleteInventory(id int) error {
	query := `
	DELETE FROM avnginventory
	WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
