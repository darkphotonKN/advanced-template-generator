package item

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type service struct {
	repo   Repository
	logger *slog.Logger
}

type Repository interface {
	Create(ctx context.Context, item *Item) error
	List(ctx context.Context, limit, offset int) ([]*Item, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Item, error)
	Update(ctx context.Context, id uuid.UUID, item *Item) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewService(repo Repository, logger *slog.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateItem(ctx context.Context, req CreateItemRequest) (*Item, error) {
	item := &Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, item); err != nil {
		s.logger.Error("failed to create item", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	return item, nil
}

func (s *service) ListItems(ctx context.Context, limit, offset int) ([]*Item, int64, error) {
	items, total, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		s.logger.Error("failed to list items", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("failed to list items: %w", err)
	}

	return items, total, nil
}

func (s *service) GetItem(ctx context.Context, id uuid.UUID) (*Item, error) {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get item",
			slog.String("error", err.Error()),
			slog.String("id", id.String()))
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (s *service) UpdateItem(ctx context.Context, id uuid.UUID, req UpdateItemRequest) (*Item, error) {
	// Get existing item
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find item: %w", err)
	}

	// Update fields
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Description != "" {
		item.Description = req.Description
	}

	// Save updates
	if err := s.repo.Update(ctx, id, item); err != nil {
		s.logger.Error("failed to update item",
			slog.String("error", err.Error()),
			slog.String("id", id.String()))
		return nil, fmt.Errorf("failed to update item: %w", err)
	}

	return item, nil
}

func (s *service) DeleteItem(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error("failed to delete item",
			slog.String("error", err.Error()),
			slog.String("id", id.String()))
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}

