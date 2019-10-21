package repository

import (
	"context"

	"github.com/william-carvalho/go-crud-mysql-rest/models"
)

type PostRepo interface {
	getByID(ctx context.Context, id int64) (*models.Post, error)
	Fetch(ctx context.Context, num int64) ([]*models.Post, error)
	Create(ctx context.Context, p *models.Post) (int64, error)
	Update(ctx context.Context, p *models.Post) (*models.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
