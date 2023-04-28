package controllers

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"github.com/jpdel518/go-graphql-gateway/gateway/requests/group_requests"
	"github.com/jpdel518/go-graphql-gateway/gateway/resources"
	"log"
	"net/http"
)

type GroupController struct {
}

func NewGroupController() *GroupController {
	return &GroupController{}
}

func (g *GroupController) Fetch(ctx context.Context, name *string, offset *int, limit *int, sort *int) ([]*model.Group, error) {
	// create request
	endpoint := group_requests.NewFetchRequest(name, offset, limit, sort).MakeEndpoint()

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// send request
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewGroupResource().ParseToArray(response.Body)
}

func (g *GroupController) Get(ctx context.Context, id *string) (*model.Group, error) {
	// create request
	endpoint := group_requests.NewGetRequest(id).MakeEndpoint()

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// send request
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewGroupResource().Parse(response.Body)
}

func (g *GroupController) Store(ctx context.Context, input model.NewGroup) (*model.Group, error) {
	// create request
	r := group_requests.NewStoreRequest(input.Name, input.Description)
	endpoint := r.MakeEndpoint()
	body, contentType, err := r.MakeRequest()
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}
	request.Header.Add("Content-Type", contentType)

	// send request
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// resolve response
	return resources.NewGroupResource().Parse(response.Body)
}
