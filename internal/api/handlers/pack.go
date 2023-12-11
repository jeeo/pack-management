package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	domainerrors "github.com/jeeo/pack-management/internal/errors"
	"github.com/jeeo/pack-management/internal/model"
)

type PackApplication interface {
	FindAll(ctx context.Context) ([]model.Pack, error)
	Create(context.Context, int) error
	Update(context.Context, model.Pack) error
	Delete(context.Context, string) error
}

type PackHandler struct {
	packApplication PackApplication
}

func (p PackHandler) Configure(router *mux.Router) {
	s := router.PathPrefix("/package").Subrouter()

	s.HandleFunc("", p.FindAll).Methods("GET")
	s.HandleFunc("", p.Create).Methods("POST")
	s.HandleFunc("/{id}", p.Update).Methods("PUT")
	s.HandleFunc("/{id}", p.Delete).Methods("DELETE")
}

func NewPackHandler(packApplication PackApplication) PackHandler {
	return PackHandler{
		packApplication: packApplication,
	}
}

func (p PackHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	packs, err := p.packApplication.FindAll(r.Context())
	if err != nil {
		log.Println("failed to find packs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dtoPacks := make([]PackDTO, 0, len(packs))
	for _, pack := range packs {
		dtoPacks = append(dtoPacks, ToPackDTO(pack))
	}

	response := FindAllPackResponse{
		Packs: dtoPacks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

func (p PackHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreatePackRequest
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

	err := p.packApplication.Create(r.Context(), request.Amount)
	if err != nil {
		if errors.As(err, &domainerrors.ErrPackNotFound{}) {
			log.Println("failed to update pack", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Println("failed to create pack", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p PackHandler) Update(w http.ResponseWriter, r *http.Request) {
	packId := mux.Vars(r)["id"]
	var request UpdatePackRequest
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("failed to decode response", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsValidUUID(packId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid pack id"))
		return
	}

	if request.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid amount"))
		return
	}

	pack := model.Pack{
		ID:     packId,
		Amount: request.Amount,
	}

	err := p.packApplication.Update(r.Context(), pack)
	if err != nil {
		log.Println("failed to create pack", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p PackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	packId := mux.Vars(r)["id"]

	if !IsValidUUID(packId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid pack id"))
		return
	}

	err := p.packApplication.Delete(r.Context(), packId)
	if err != nil {
		log.Println("failed to create pack", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
