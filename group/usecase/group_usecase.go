package usecase

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/group/domain/model"
	"github.com/jpdel518/go-graphql-gateway/group/domain/repository"
	"time"
)

type GroupUsecase interface {
	Fetch(ctx context.Context, offset int, num int, name string, order int) ([]*model.Group, error)
	GetByID(ctx context.Context, id int) (*model.Group, error)
	Create(ctx context.Context, u *model.Group) (*model.Group, error)
	Update(ctx context.Context, u *model.Group) (*model.Group, error)
	Delete(ctx context.Context, id int) error
}

type groupUsecase struct {
	groupRepo      repository.GroupRepository
	contextTimeout time.Duration
}

// NewGroupUsecase will create new an groupUsecase object
func NewGroupUsecase(u repository.GroupRepository, timeout time.Duration) GroupUsecase {
	return &groupUsecase{
		groupRepo:      u,
		contextTimeout: timeout,
	}
}

// Fetch will retrieve sites from the repository
func (usecase *groupUsecase) Fetch(c context.Context, offset int, num int, name string, sort int) ([]*model.Group, error) {
	// デフォルト値を設定
	if num <= 0 {
		num = 10
	}

	ordering := "desc"
	order := "updated_at"
	if sort%2 == 0 {
		ordering = "asc"
	}
	if sort < 2 {
		order = "updated_at"
	} else {
		order = "name"
	}

	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.groupRepo.Fetch(ctx, offset, num, name, order, ordering)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID will find a site by id
func (usecase *groupUsecase) GetByID(c context.Context, id int) (*model.Group, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.groupRepo.GetByID(ctx, id)
	if err != nil {
		return &model.Group{}, err
	}

	return res, nil
}

// Create will register a site
func (usecase *groupUsecase) Create(c context.Context, u *model.Group) (*model.Group, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	// create user
	site, err := usecase.groupRepo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return site, nil
}

// Update will update a site
func (usecase *groupUsecase) Update(c context.Context, s *model.Group) (*model.Group, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	_, err := usecase.groupRepo.Update(ctx, s)
	return s, err
}

// Delete will delete a site by id
func (usecase *groupUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	return usecase.groupRepo.Delete(ctx, id)
}
