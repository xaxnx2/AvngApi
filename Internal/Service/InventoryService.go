package Service

import (
	"HOMEWORKAVENGINVAPI/Internal/Entity"
	"HOMEWORKAVENGINVAPI/Internal/Repository"
)

// func GetVillainAll() []entity.Villain {
// 	return Repository.GetAllVillain()
// }

type InventoryService struct {
	invRepo *Repository.InventoryRepository
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		invRepo: Repository.NewInventoryRepository(),
	}
}

func (s *InventoryService) GetInventory() ([]Entity.Inventory, error) {
	return s.invRepo.GetAllInventory()
}
