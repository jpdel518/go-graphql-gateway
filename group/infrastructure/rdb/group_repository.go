package rdb

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/group/domain/model"
	"github.com/jpdel518/go-graphql-gateway/group/domain/repository"
	"github.com/jpdel518/go-graphql-gateway/group/ent"
	"github.com/jpdel518/go-graphql-gateway/group/ent/group"
	"log"
)

type userRepository struct {
	client *ent.Client
}

func NewSiteRepository(client *ent.Client) repository.GroupRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Fetch(ctx context.Context, offset int, num int, name string, order string, ordering string) ([]*model.Group, error) {
	res := make([]*model.Group, 0)

	// fetch users
	query := r.client.Group.Query().
		Where(group.NameContains(name)).
		Offset(offset).
		Limit(num)
	if ordering == "asc" {
		query = query.Order(ent.Asc(order))
	} else {
		query = query.Order(ent.Desc(order))
	}
	sites, err := query.All(ctx)
	if err != nil {
		log.Printf("failed fetching users: %v", err)
		return res, err
	}

	// ent.Group -> model.Group
	for _, s := range sites {
		res = append(res, &model.Group{
			ID:          s.ID,
			Name:        s.Name,
			Description: s.Description,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
		})
	}
	return res, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*model.Group, error) {
	// get user
	s, err := r.client.Group.Get(ctx, id)
	if err != nil {
		log.Printf("failed getbyid user: %v", err)
		return nil, err
	}

	// ent.Group -> model.Group
	return &model.Group{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, s *model.Group) (*model.Group, error) {
	data, err := r.client.Group.Create().
		SetName(s.Name).
		SetDescription(s.Description).
		Save(ctx)

	if err != nil {
		log.Printf("failed creating user: %v", err)
		return nil, err
	}
	log.Printf("user was created: %v", data)

	// ent.Site -> model.Site
	res := &model.Group{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return res, err
}

func (r *userRepository) Update(ctx context.Context, s *model.Group) (*model.Group, error) {
	data, err := r.client.Group.Update().
		Where(group.ID(s.ID)).
		SetName(s.Name).
		SetDescription(s.Description).
		Save(ctx)

	if err != nil {
		log.Printf("failed updating user: %v", err)
		return nil, err
	}
	log.Printf("user was updated: %v", data)

	return s, err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	return r.client.Group.DeleteOneID(id).Exec(ctx)
}
