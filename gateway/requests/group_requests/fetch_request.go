package group_requests

import (
	"log"
	"strconv"
)

type FetchRequest struct {
	name   *string
	offset *int
	limit  *int
	sort   *int
	user   *string
}

func NewFetchRequest(name *string, offset *int, limit *int, sort *int) *FetchRequest {
	return &FetchRequest{
		name:   name,
		offset: offset,
		limit:  limit,
		sort:   sort,
	}
}

func (r *FetchRequest) MakeEndpoint() string {
	const URL = "http://app2:8081/api/groups"

	endpoint := URL
	if r.name != nil {
		endpoint += "?name=" + *r.name
	}
	if r.offset != nil {
		if endpoint == URL {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "offset=" + strconv.Itoa(*r.offset)
	}
	if r.limit != nil {
		if endpoint == URL {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "num=" + strconv.Itoa(*r.limit)
	}
	if r.sort != nil {
		if endpoint == URL {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "sort=" + strconv.Itoa(*r.sort)
	}
	log.Printf("URL: %v", endpoint)

	return endpoint
}
