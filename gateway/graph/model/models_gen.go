// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type Group struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Users       []*User `json:"users"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type NewGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID        string          `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Age       *int            `json:"age,omitempty"`
	Address   string          `json:"address"`
	Email     string          `json:"email"`
	GroupID   int             `json:"groupId"`
	Group     *Group          `json:"group"`
	Avator    *graphql.Upload `json:"avator,omitempty"`
	CreatedAt string          `json:"createdAt"`
	UpdatedAt string          `json:"updatedAt"`
}