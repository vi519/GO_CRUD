package routes

import (
	"go-crud-live/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	//Fetch/Modify Accounts Routes
	router.HandleFunc("/accounts", controllers.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", controllers.DeleteAccout).Methods("DELETE")

	return router
}
