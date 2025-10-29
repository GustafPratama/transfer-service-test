package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{conn: conn}
}

func (r *Repository) InitLocations(ctx context.Context) {
	r.conn.Exec(ctx, "INSERT INTO locations (name, capacity, used) VALUES ('LOC-A',10,2) ON CONFLICT DO NOTHING;")
	r.conn.Exec(ctx, "INSERT INTO locations (name, capacity, used) VALUES ('LOC-B',10,5) ON CONFLICT DO NOTHING;")
}

func (r *Repository) CreateTransfer(ctx context.Context, t *Transfer) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO transfers (pallet_id, from_location, to_location, note, status) VALUES ($1,$2,$3,$4,$5)",
		t.PalletID, t.FromLocation, t.ToLocation, t.Note, "REQUESTED")
	return err
}

func (r *Repository) UpdateStatus(ctx context.Context, id int, status string) error {
	_, err := r.conn.Exec(ctx, "UPDATE transfers SET status=$1 WHERE id=$2", status, id)
	return err
}

func (r *Repository) GetTransfer(ctx context.Context, id int) (*Transfer, error) {
	row := r.conn.QueryRow(ctx, "SELECT id,pallet_id,from_location,to_location,note,status FROM transfers WHERE id=$1", id)
	var t Transfer
	err := row.Scan(&t.ID, &t.PalletID, &t.FromLocation, &t.ToLocation, &t.Note, &t.Status)
	return &t, err
}

func (r *Repository) SaveOutbox(ctx context.Context, eventType string, data interface{}) error {
	payload, _ := json.Marshal(data)
	_, err := r.conn.Exec(ctx, "INSERT INTO outbox (event_type,payload) VALUES ($1,$2)", eventType, payload)
	return err
}

func (r *Repository) FlushOutbox(ctx context.Context) error {
	rows, _ := r.conn.Query(ctx, "SELECT id, event_type, payload FROM outbox WHERE processed=false")
	defer rows.Close()
	for rows.Next() {
		var id int
		var eventType string
		var payload []byte
		rows.Scan(&id, &eventType, &payload)
		filename := fmt.Sprintf("event_%d.json", id)
		_ = os.WriteFile(filename, payload, 0644)
		r.conn.Exec(ctx, "UPDATE outbox SET processed=true WHERE id=$1", id)
	}
	return nil
}
