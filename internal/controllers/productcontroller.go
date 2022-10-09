package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fabiotavarespr/golang-crud-rest-api/internal/database"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/entities"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/logger"
	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Handling POST - /api/products - %s", r.URL))
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
	logger.Info(fmt.Sprintf("CreateProduct - %v", product))
}

func checkIfProductExists(productId string) bool {
	logger.Info(fmt.Sprintf("checkIfProductExists - %s", productId))
	var product entities.Product
	database.Instance.First(&product, productId)
	return product.ID != 0
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Handling GET - /api/products/{id} - %s", r.URL))
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		logger.Info(fmt.Sprintf("Product With %s Not Found!", productId))
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	logger.Info(fmt.Sprintf("GetProductById - %v", product))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Handling GET - /api/products - %s", r.URL))
	var products []entities.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	logger.Info(fmt.Sprintf("GetProducts - size %v", len(products)))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Handling PUT - /api/products/{id} - %s", r.URL))
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		logger.Info(fmt.Sprintf("Product With %s Not Found!", productId))
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	logger.Info(fmt.Sprintf("UpdateProduct - %v", product))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Handling DELETE - /api/products/{id} - %s", r.URL))
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		logger.Info(fmt.Sprintf("Product With %s Not Found!", productId))
		return
	}
	var product entities.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
	logger.Info("Product Deleted Successfully!")
}
