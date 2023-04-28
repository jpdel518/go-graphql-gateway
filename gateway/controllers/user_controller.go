package controllers

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"github.com/jpdel518/go-graphql-gateway/gateway/requests/user_requests"
	"github.com/jpdel518/go-graphql-gateway/gateway/resources"
	"log"
	"net/http"
	"runtime"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// Fetch fetches users
func (u *UserController) Fetch(ctx context.Context, name *string, offset *int, limit *int, sort *int) ([]*model.User, error) {
	endpoint := user_requests.NewFetchRequest(user_requests.WithName(name), user_requests.WithOffset(offset),
		user_requests.WithLimit(limit), user_requests.WithSort(sort)).MakeEndpoint()

	// request
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewUserResource().ParseToArray(response.Body)
}

// FetchByGroup is the resolver for the user field.
func (u *UserController) FetchByGroup(ctx context.Context, obj *model.Group) ([]*model.User, error) {
	// request
	endpoint := user_requests.NewFetchRequest(user_requests.WithGroup(&obj.ID)).MakeEndpoint()

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewUserResource().ParseToArray(response.Body)
}

func (u *UserController) Get(ctx context.Context, id *string) (*model.User, error) {
	// request
	endpoint := user_requests.NewGetRequest(id).MakeEndpoint()

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewUserResource().Parse(response.Body)
}

// Store is the resolver for the createUser field.
func (u *UserController) Store(ctx context.Context, firstName string, lastName string, age *int, address string,
	email string, groupID int, avatar *graphql.Upload) (*model.User, error) {
	log.Println("Start  CreateUser: ", firstName, lastName, age, address, email, groupID, avatar)

	// メモリ計測1
	rm := runtime.MemStats{}
	runtime.ReadMemStats(&rm)
	log.Printf("before alloc memory %v", rm.Alloc)
	log.Printf("before total alloc memory %v", rm.TotalAlloc)
	log.Printf("before system memory %v", rm.Sys)

	// create request
	r := user_requests.NewStoreRequest(firstName, lastName, age, address, email, groupID, avatar)
	endpoint := r.MakeEndpoint()
	body, contentType, err := r.MakeRequest()
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// send request
	request, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	request.Header.Add("Content-Type", contentType)

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// メモリ計測2
	runtime.ReadMemStats(&rm)
	log.Printf("after alloc memory %v", rm.Alloc)
	log.Printf("after total alloc memory %v", rm.TotalAlloc)
	log.Printf("after system memory %v", rm.Sys)

	// resolve response
	log.Printf("Response : %v", response)
	return resources.NewUserResource().Parse(response.Body)
}
