package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeeo/pack-management/internal/model"
)

type OrderHandler struct {
	orderApplication OrderApplication
}

type OrderApplication interface {
	CalculateOrderPack(context.Context, int) ([]model.OrderPack, error)
}

func NewOrderHandler(orderApplication OrderApplication) OrderHandler {
	return OrderHandler{
		orderApplication: orderApplication,
	}
}

func (p OrderHandler) Configure(router *mux.Router) {
	s := router.PathPrefix("/order").Subrouter()

	s.HandleFunc("/calculate", p.CalculateOrder).Methods("POST")
}

func (p OrderHandler) CalculateOrder(w http.ResponseWriter, r *http.Request) {
	var request CalculateOrderPackRequest
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("failed to decode response", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderPacks, err := p.orderApplication.CalculateOrderPack(r.Context(), request.Amount)
	if err != nil {
		log.Println("failed to calculate order", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := CalculateOrderPackResponse{
		OrderPacks: ToOrderPackDTOs(orderPacks),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(responseBytes); err != nil {
		log.Println("failed to write response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
