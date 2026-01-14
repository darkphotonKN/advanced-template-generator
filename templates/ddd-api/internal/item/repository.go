package item

import (
	"context"
	"database/sql"

	"github.com/darkphotonKN/go-template-generator/internal/utils/errorutils"
	"github.com/google/uuid"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, item *Item) error {
	query := `
		INSERT INTO items (id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query, item.ID, item.Name, item.Description).Scan(&item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}

	return nil
}

func (r *repository) List(ctx context.Context, limit, offset int) ([]*Item, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM items`
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, errorutils.AnalyzeDBErr(err)
	}

	// Get items
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM items
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, errorutils.AnalyzeDBErr(err)
	}
	defer rows.Close()

	var items []*Item
	for rows.Next() {
		item := &Item{}
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, 0, errorutils.AnalyzeDBErr(err)
		}
		items = append(items, item)
	}

	return items, total, nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*Item, error) {
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM items
		WHERE id = $1
	`

	item := &Item{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, errorutils.AnalyzeDBErr(err)
	}

	return item, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, item *Item) error {
	query := `
		UPDATE items
		SET name = $2, description = $3, updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`

	err := r.db.QueryRowContext(ctx, query, id, item.Name, item.Description).Scan(&item.UpdatedAt)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}

	if rowsAffected == 0 {
		return errorutils.ErrNotFound
	}

	return nil
}

