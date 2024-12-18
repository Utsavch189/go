package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Utsavch189/api_mysql/internal/controller"
	"github.com/Utsavch189/api_mysql/internal/models/requests"
	"github.com/Utsavch189/api_mysql/internal/models/responses"
	"github.com/Utsavch189/api_mysql/internal/utils"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := controller.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "fetching failed!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.GetAllProductResponse(products))
}

func GetAProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	product, err := controller.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "fetching failed!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.GetAProductResponse(product))
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct requests.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "creation failed!"))
		return
	}
	product := requests.NewProduct(newProduct.ProductName, newProduct.Price)
	if err := utils.ValidateStruct(product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "creation failed!", utils.FormatValidationError(err)))
		return
	}
	createdProduct, cerr := controller.AddProduct(product)
	if cerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(cerr, "creation failed!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.CreateProductRepsonse(createdProduct, "creation successful!"))
}

func UpdateAProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product requests.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "updation failed!"))
		return
	}
	if err := utils.ValidateStruct(product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "updation failed!", utils.FormatValidationError(err)))
		return
	}
	updatedProduct, err := controller.UpdateProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "updation failed!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.UpdateProductRepsonse(updatedProduct, "updation successful!"))
}

func DeleteAProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	err := controller.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "deletion failed!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.DeleteProductRepsonse(fmt.Sprintf("%s successfully deleted!", id)))
}
