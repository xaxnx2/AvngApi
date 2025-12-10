package Router

import (
	"HOMEWORKAVENGINVAPI/Internal/Handler"

	"github.com/gorilla/mux"
)

type Routing struct {
	invhand *Handler.InventoryHandler
}

func NewRouting() *Routing {
	return &Routing{
		invhand: Handler.NewInventoryHandler(),
	}
}

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/Get/Inventories", NewRouting().invhand.GetInventory).Methods("GET")
	// r.HandleFunc("/Get/inventories/:id", Handler.GetInventoryByID).Methods("GET")
	// r.HandleFunc("/POST/inventories", Handler.PostInventory).Methods("POST")
	// r.HandleFunc("/PUT/inventories/:id", Handler.UpdateInvetory).Methods("PUT")
	// r.HandleFunc("/DELETE/inventories/:id", Handler.DeleteInvetory).Methods("DELETE")
	return r
}
