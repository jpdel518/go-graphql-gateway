package group_requests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type StoreRequest struct {
	name        string
	description string
}

func NewStoreRequest(name string, description string) *StoreRequest {
	return &StoreRequest{
		name:        name,
		description: description,
	}
}

func (r *StoreRequest) MakeEndpoint() string {
	return "http://app2:8081/api/groups"
}

func (r *StoreRequest) MakeRequest() (body io.Reader, contentType string, err error) {
	// application/x-www-form-urlencodedで送りたい場合
	// values := url.Values{}
	// values.Add("name", r.name)
	// values.Add("description", r.description)
	// contentType = "application/x-www-form-urlencoded"

	// application/jsonで送りたい場合
	values := map[string]string{
		"name":        r.name,
		"description": r.description,
	}
	contentType = "application/json"

	// jsonへ変換
	data, err := json.Marshal(values)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return nil, "", err
	}
	// io.Readerを作成
	body = bytes.NewReader(data)

	return body, contentType, nil
}
