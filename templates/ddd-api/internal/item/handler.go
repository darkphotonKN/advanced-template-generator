package item

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/darkphotonKN/go-template-generator/internal/utils/errorutils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
	logger  *slog.Logger
}

type Service interface {
	CreateItem(ctx context.Context, req CreateItemRequest) (*Item, error)
	ListItems(ctx context.Context, limit, offset int) ([]*Item, int64, error)
	GetItem(ctx context.Context, id uuid.UUID) (*Item, error)
	UpdateItem(ctx context.Context, id uuid.UUID, req UpdateItemRequest) (*Item, error)
	DeleteItem(ctx context.Context, id uuid.UUID) error
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateItem(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind request", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	item, err := h.service.CreateItem(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, errorutils.ErrDuplicateResource):
			c.JSON(409, gin.H{"error": "Item already exists"})
		case errors.Is(err, errorutils.ErrInvalidInput):
			c.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, errorutils.ErrConstraintViolation):
			c.JSON(400, gin.H{"error": "Invalid data provided"})
		default:
			h.logger.Error("failed to create item", slog.String("error", err.Error()))
			c.JSON(500, gin.H{"error": "Failed to create item"})
		}
		return
	}

	c.JSON(201, item)
}

func (h *Handler) ListItems(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	items, total, err := h.service.ListItems(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list items"})
		return
	}

	response := ListItemsResponse{
		Items:  items,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}

	c.JSON(200, response)
}

func (h *Handler) GetItem(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "item ID is required"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item ID"})
		return
	}

	item, err := h.service.GetItem(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, errorutils.ErrNotFound) {
			c.JSON(404, gin.H{"error": "Item not found"})
		} else {
			h.logger.Error("failed to get item", slog.String("error", err.Error()))
			c.JSON(500, gin.H{"error": "Failed to retrieve item"})
		}
		return
	}

	c.JSON(200, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "item ID is required"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item ID"})
		return
	}

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("failed to bind request", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	item, err := h.service.UpdateItem(c.Request.Context(), id, req)
	if err != nil {
		switch {
		case errors.Is(err, errorutils.ErrNotFound):
			c.JSON(404, gin.H{"error": "Item not found"})
		case errors.Is(err, errorutils.ErrDuplicateResource):
			c.JSON(409, gin.H{"error": "Item with that name already exists"})
		case errors.Is(err, errorutils.ErrInvalidInput):
			c.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, errorutils.ErrConstraintViolation):
			c.JSON(400, gin.H{"error": "Invalid data provided"})
		default:
			h.logger.Error("failed to update item", slog.String("error", err.Error()))
			c.JSON(500, gin.H{"error": "Failed to update item"})
		}
		return
	}

	c.JSON(200, item)
}

func (h *Handler) DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "item ID is required"})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid item ID"})
		return
	}

	if err := h.service.DeleteItem(c.Request.Context(), id); err != nil {
		if errors.Is(err, errorutils.ErrNotFound) {
			c.JSON(404, gin.H{"error": "Item not found"})
		} else {
			h.logger.Error("failed to delete item",
				slog.String("error", err.Error()),
				slog.String("id", id.String()))
			c.JSON(500, gin.H{"error": "Failed to delete item"})
		}
		return
	}

	c.Status(204)
}