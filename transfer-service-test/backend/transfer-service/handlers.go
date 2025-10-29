package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux, svc *Service, repo *Repository) {
	r.Post("/transfers", func(w http.ResponseWriter, r *http.Request) {
		var t Transfer
		_ = json.NewDecoder(r.Body).Decode(&t)
		if err := svc.ValidateCapacity(context.Background(), t.ToLocation); err != nil {
			http.Error(w, "Lokasi penuh", 400)
			return
		}
		repo.CreateTransfer(context.Background(), &t)
		repo.SaveOutbox(context.Background(), "TransferCreated", t)
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(t)
	})

	r.Post("/transfers/{id}/accept", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		repo.UpdateStatus(context.Background(), id, "ACCEPTED")
		w.Write([]byte("accepted"))
	})

	r.Post("/transfers/{id}/complete", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		repo.UpdateStatus(context.Background(), id, "COMPLETED")
		repo.SaveOutbox(context.Background(), "TransferCompleted", map[string]int{"id": id})
		w.Write([]byte("completed"))
	})

	r.Get("/transfers/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		t, err := repo.GetTransfer(context.Background(), id)
		if err != nil {
			http.Error(w, "not found", 404)
			return
		}
		json.NewEncoder(w).Encode(t)
	})

	r.Post("/dev/flush-outbox", func(w http.ResponseWriter, r *http.Request) {
		repo.FlushOutbox(context.Background())
		w.Write([]byte("flushed"))
	})
}
