package item

import (
	"context"

	"github.com/darkphotonKN/go-template-generator/internal/utils/errorutils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, item *Item) error {
	query := `
		INSERT INTO items (id, name, description, created_at, updated_at)
		VALUES (:id, :name, :description, NOW(), NOW())
		RETURNING created_at, updated_at
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, item, item)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}

	return nil
}

func (r *repository) List(ctx context.Context, limit, offset int) ([]*Item, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM items`
	err := r.db.GetContext(ctx, &total, countQuery)
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

	var items []*Item
	err = r.db.SelectContext(ctx, &items, query, limit, offset)
	if err != nil {
		return nil, 0, errorutils.AnalyzeDBErr(err)
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
	err := r.db.GetContext(ctx, item, query, id)
	if err != nil {
		return nil, errorutils.AnalyzeDBErr(err)
	}

	return item, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, item *Item) error {
	query := `
		UPDATE items
		SET name = :name, description = :description, updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at
	`

	// Set the ID for the update
	item.ID = id

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return errorutils.AnalyzeDBErr(err)
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &item.UpdatedAt, item)
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

