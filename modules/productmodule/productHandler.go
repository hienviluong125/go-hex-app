package productmodule

import (
	"encoding/json"
	"fmt"
	"hienviluong125/go-hex-app/common"
	"hienviluong125/go-hex-app/errorhandler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Index(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.FindAll()

	if err != nil {
		panic(err)
	}

	common.RespondWithStatus(w, http.StatusOK, products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var productParams Product
	err := json.NewDecoder(r.Body).Decode(&productParams)

	if err != nil {
		fmt.Println(err.Error())
		panic(errorhandler.ErrBadRequest(err))
	}

	if err := h.service.Save(productParams); err != nil {
		panic(err)
	}

	common.WriteStatus(w, http.StatusOK)
}

func (h *ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["product_id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	product, err := h.service.FindById(id)

	if err != nil {
		panic(err)
	}

	common.RespondWithStatus(w, http.StatusOK, product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["product_id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	var updateProductParams UpdateProduct
	if err := json.NewDecoder(r.Body).Decode(&updateProductParams); err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	err = h.service.UpdateById(id, updateProductParams)

	if err != nil {
		panic(err)
	}

	common.WriteStatus(w, http.StatusOK)
}

func (h *ProductHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["product_id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	if err = h.service.DestroyById(id); err != nil {
		panic(err)
	}

	common.WriteStatus(w, http.StatusOK)
}
