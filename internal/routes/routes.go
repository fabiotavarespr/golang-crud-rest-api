package routes

import (
	"net/http"

	"github.com/fabiotavarespr/golang-crud-rest-api/internal/controllers"
	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods(http.MethodDelete)
}
