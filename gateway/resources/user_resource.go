package resources

import (
	"encoding/json"
	"fmt"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"io"
	"log"
	"strconv"
)

type UserResource struct {
}

func NewUserResource() *UserResource {
	return &UserResource{}
}

func (r *UserResource) Parse(body io.ReadCloser) (*model.User, error) {
	var data map[string]interface{}
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return nil, err
	}
	log.Printf("Response body: %v", data)

	// get response code
	code, err := data["code"].(json.Number).Int64()
	if err != nil {
		log.Printf("Error getting code: %v", err)
		return nil, err
	}

	// create response user
	// log.Printf("Success getting code: %v", data["code"])
	var user *model.User
	if data != nil && code == 2000 {
		user = r.createModel(data["data"])
	} else {
		log.Printf("Error getting code: %v user: %v", data["code"], data["data"])
		return nil, fmt.Errorf("error getting code: %v user: %v", data["code"], data["data"])
	}

	return user, nil
}

func (r *UserResource) ParseToArray(body io.ReadCloser) ([]*model.User, error) {
	var data map[string]interface{}
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return nil, err
	}
	log.Printf("Response body: %v", data)

	// get response code
	code, err := data["code"].(json.Number).Int64()
	if err != nil {
		log.Printf("Error getting code: %v", err)
		return nil, err
	}

	// create response user
	// log.Printf("Success getting code: %v", data["code"])
	var users []*model.User
	if data != nil && code == 2000 {
		for _, user := range data["data"].([]interface{}) {
			// append users
			users = append(users, r.createModel(user))
		}
	} else {
		log.Printf("Error getting code: %v users: %v", data["code"], data["data"])
		return nil, fmt.Errorf("error getting code: %v users: %v", data["code"], data["data"])
	}

	return users, nil
}

func (r *UserResource) createModel(data interface{}) *model.User {
	// id
	id, _ := data.(map[string]interface{})["id"].(json.Number).Int64()
	// age
	age64, _ := data.(map[string]interface{})["age"].(json.Number).Int64()
	age := int(age64)
	// user
	user := &model.User{
		ID:        strconv.Itoa(int(id)),
		FirstName: data.(map[string]interface{})["first_name"].(string),
		LastName:  data.(map[string]interface{})["last_name"].(string),
		Age:       &age,
		Address:   data.(map[string]interface{})["address"].(string),
		Email:     data.(map[string]interface{})["email"].(string),
		CreatedAt: data.(map[string]interface{})["created_at"].(string),
		UpdatedAt: data.(map[string]interface{})["created_at"].(string),
	}

	return user
}
