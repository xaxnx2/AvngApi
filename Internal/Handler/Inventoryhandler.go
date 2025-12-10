package Handler

import (
	entity "HOMEWORKAVENGINVAPI/Internal/Entity"
	"HOMEWORKAVENGINVAPI/Internal/Service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func GetVillain(w http.ResponseWriter, r *http.Request) {
// 	villain := Service.GetVillainAll()

//		json.NewEncoder(w).Encode(villain)
//	}
type InventoryHandler struct {
	invserv *Service.InventoryService
}

func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{
		invserv: Service.NewInventoryService(),
	}
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	if h.invserv == nil {
		log.Println("Error: Inventory repository is not initialized")
		http.Error(w, "Inventory repository is not initialized", http.StatusInternalServerError)
		return
	}

	inventory, error := h.invserv.GetInventory()
	if error != nil {
		log.Printf("Error retrieving inventory: %v\n", error)
		http.Error(w, fmt.Sprintf("Error retrieving inventory: %v", error), http.StatusInternalServerError)
		response := entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error Getting Inventory",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	response := entity.Response{
		StatusCode: http.StatusOK,
		Message:    "Users fetched successfully",
		Data:       inventory,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		response := entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error Getting Inventory",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
}
