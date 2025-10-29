package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@db:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(conn)
	repo.InitLocations(ctx)
	svc := NewService(repo)

	r := chi.NewRouter()
	RegisterRoutes(r, svc, repo)

	fmt.Println("Transfer Service running on :8080")
	http.ListenAndServe(":8080", r)
}
