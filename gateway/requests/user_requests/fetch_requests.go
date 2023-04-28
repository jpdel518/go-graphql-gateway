package user_requests

import (
	"log"
	"strconv"
)

type FetchRequest struct {
	name   *string
	offset *int
	limit  *int
	sort   *int
	group  *string
}

func NewFetchRequest(opts ...FetchRequestOption) *FetchRequest {
	r := &FetchRequest{
		name:   nil,
		offset: nil,
		limit:  nil,
		sort:   nil,
		group:  nil,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *FetchRequest) MakeEndpoint() string {
	endpoint := "http://app1:8080/api/users"
	if r.name != nil {
		endpoint += "?name=" + *r.name
	}
	if r.offset != nil {
		if endpoint == "http://app1:8080/api/users" {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "offset=" + strconv.Itoa(*r.offset)
	}
	if r.limit != nil {
		if endpoint == "http://app1:8080/api/users" {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "num=" + strconv.Itoa(*r.limit)
	}
	if r.sort != nil {
		if endpoint == "http://app1:8080/api/users" {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "sort=" + strconv.Itoa(*r.sort)
	}
	if r.group != nil {
		if endpoint == "http://app1:8080/api/users" {
			endpoint += "?"
		} else {
			endpoint += "&"
		}
		endpoint += "group=" + *r.group
	}
	log.Printf("URL: %v", endpoint)

	return endpoint
}

type FetchRequestOption func(*FetchRequest)

func WithName(name *string) FetchRequestOption {
	return func(r *FetchRequest) {
		r.name = name
	}
}

func WithOffset(offset *int) FetchRequestOption {
	return func(r *FetchRequest) {
		r.offset = offset
	}
}

func WithLimit(limit *int) FetchRequestOption {
	return func(r *FetchRequest) {
		r.limit = limit
	}
}

func WithSort(sort *int) FetchRequestOption {
	return func(r *FetchRequest) {
		r.sort = sort
	}
}

func WithGroup(group *string) FetchRequestOption {
	return func(r *FetchRequest) {
		r.group = group
	}
}
