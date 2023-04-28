package rdb

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/user/domain/model"
	"github.com/jpdel518/go-graphql-gateway/user/domain/repository"
	"github.com/jpdel518/go-graphql-gateway/user/ent"
	"github.com/jpdel518/go-graphql-gateway/user/ent/user"
	"log"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Fetch(ctx context.Context, offset int, num int, name string, group int, order string, ordering string) ([]*model.User, error) {
	res := make([]*model.User, 0)

	// fetch users
	query := r.client.User.Query().
		Where(user.FirstNameContains(name)).
		Offset(offset).
		Limit(num)
	if group > 0 {
		query = query.Where(user.GroupID(group))
	}
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
		res = append(res, &model.User{
			ID:        s.ID,
			FirstName: s.FirstName,
			LastName:  s.LastName,
			Age:       s.Age,
			Address:   s.Address,
			Email:     s.Email,
			GroupID:   s.GroupID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		})
	}
	return res, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	// get user
	s, err := r.client.User.Get(ctx, id)
	if err != nil {
		log.Printf("failed getbyid user: %v", err)
		return nil, err
	}

	// ent.Group -> model.Group
	return &model.User{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Age:       s.Age,
		Address:   s.Address,
		Email:     s.Email,
		GroupID:   s.GroupID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, s *model.User) (*model.User, error) {
	data, err := r.client.User.Create().
		SetFirstName(s.FirstName).
		SetLastName(s.LastName).
		SetAge(s.Age).
		SetAddress(s.Address).
		SetEmail(s.Email).
		SetGroupID(s.GroupID).
		Save(ctx)

	if err != nil {
		log.Printf("failed creating user: %v", err)
		return nil, err
	}
	log.Printf("user was created: %v", data)

	// ent.Site -> model.Site
	res := &model.User{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Age:       data.Age,
		Address:   data.Address,
		Email:     data.Email,
		GroupID:   data.GroupID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return res, err
}

func (r *userRepository) Update(ctx context.Context, s *model.User) (*model.User, error) {
	data, err := r.client.User.Update().
		Where(user.ID(s.ID)).
		SetFirstName(s.FirstName).
		SetLastName(s.LastName).
		SetAge(s.Age).
		SetAddress(s.Address).
		SetEmail(s.Email).
		SetGroupID(s.GroupID).
		Save(ctx)

	if err != nil {
		log.Printf("failed updating user: %v", err)
		return nil, err
	}
	log.Printf("user was updated: %v", data)

	return s, err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}
