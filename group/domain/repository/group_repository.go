package repository

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/group/domain/model"
)

type GroupRepository interface {
	Fetch(ctx context.Context, offset int, num int, name string, order string, ordering string) (res []*model.Group, err error)
	GetByID(ctx context.Context, id int) (*model.Group, error)
	Create(ctx context.Context, u *model.Group) (*model.Group, error)
	Update(ctx context.Context, u *model.Group) (*model.Group, error)
	Delete(ctx context.Context, id int) error
}
