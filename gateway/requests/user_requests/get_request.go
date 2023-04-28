package user_requests

import (
	"fmt"
)

type GetRequest struct {
	id *string
}

func NewGetRequest(id *string) *GetRequest {
	return &GetRequest{
		id: id,
	}
}

func (r *GetRequest) MakeEndpoint() string {
	return fmt.Sprintf("http://app1:8080/api/users/%v", *r.id)
}
