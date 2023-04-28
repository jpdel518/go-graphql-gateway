package usecase

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/user/domain/model"
	"github.com/jpdel518/go-graphql-gateway/user/domain/repository"
	"mime/multipart"
	"time"
)

type UserUsecase interface {
	Fetch(ctx context.Context, offset int, num int, name string, group int, order int) ([]*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	Create(ctx context.Context, u *model.User, avatar multipart.File, fh *multipart.FileHeader) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
}

type userUsecase struct {
	userRepo       repository.UserRepository
	fileRepo       repository.UserFileRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object
func NewUserUsecase(u repository.UserRepository, f repository.UserFileRepository, timeout time.Duration) UserUsecase {
	return &userUsecase{
		userRepo:       u,
		fileRepo:       f,
		contextTimeout: timeout,
	}
}

// Fetch will retrieve sites from the repository
func (usecase *userUsecase) Fetch(c context.Context, offset int, num int, name string, group int, sort int) ([]*model.User, error) {
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

	res, err := usecase.userRepo.Fetch(ctx, offset, num, name, group, order, ordering)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID will find a site by id
func (usecase *userUsecase) GetByID(c context.Context, id int) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.userRepo.GetByID(ctx, id)
	if err != nil {
		return &model.User{}, err
	}

	return res, nil
}

// Create will register a site
func (usecase *userUsecase) Create(c context.Context, u *model.User, avatar multipart.File, fh *multipart.FileHeader) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	if avatar != nil {
		_, err := usecase.fileRepo.UploadFile(avatar, fh)
		if err != nil {
			return nil, err
		}
	}

	// create user
	site, err := usecase.userRepo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return site, nil
}

// Update will update a site
func (usecase *userUsecase) Update(c context.Context, s *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	_, err := usecase.userRepo.Update(ctx, s)
	return s, err
}

// Delete will delete a site by id
func (usecase *userUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	return usecase.userRepo.Delete(ctx, id)
}
