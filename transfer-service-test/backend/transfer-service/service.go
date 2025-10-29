package main

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ValidateCapacity(ctx context.Context, to string) error {
	row := s.repo.conn.QueryRow(ctx, "SELECT capacity, used FROM locations WHERE name=$1", to)
	var capacity, used int
	if err := row.Scan(&capacity, &used); err != nil {
		return err
	}
	if used >= capacity {
		return errors.New("destination full")
	}
	return nil
}
