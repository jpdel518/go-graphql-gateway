package group_requests

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
	return fmt.Sprintf("http://app2:8081/api/groups/get-by-id/%v", *r.id)
}
