package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/internal/service"
	"ecommerce/pkg/response"
	"ecommerce/pkg/apperror"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHanler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// .... Request / Reponse DTOs------
// why seperate structs? we dont want to expose internal fields like
//CreateAt, or ID to the user when they create a product,

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json: "description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest

	//step 1: Decode json the request body into out struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// step 2: call the service (business logic lives there, not here)
	product, err := h.svc.AddProduct(req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
        response.HandleError(w, err)
        return
    }

	response.JSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	product, err := h.svc.GetProduct(id)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.ListProducts()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, products)
}
