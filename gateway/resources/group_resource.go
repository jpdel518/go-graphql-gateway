package resources

import (
	"encoding/json"
	"fmt"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"io"
	"log"
	"strconv"
)

type GroupResource struct {
}

func NewGroupResource() *GroupResource {
	return &GroupResource{}
}

func (r *GroupResource) Parse(body io.ReadCloser) (*model.Group, error) {
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
	var group *model.Group
	if data != nil && code == 2000 {
		group = r.createModel(data["data"])
	} else {
		log.Printf("Error getting code: %v groups: %v", data["code"], data["data"])
		return nil, fmt.Errorf("Error getting code: %v groups: %v", data["code"], data["data"])
	}

	return group, nil
}

func (r *GroupResource) ParseToArray(body io.ReadCloser) ([]*model.Group, error) {
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

	// create response groups
	var groups []*model.Group
	if data != nil && code == 2000 {
		log.Printf("data type: %T", data["data"])
		for _, responseData := range data["data"].([]interface{}) {
			groups = append(groups, r.createModel(responseData))
		}
	} else {
		log.Printf("Error getting code: %v groups: %v", data["code"], data["data"])
		return nil, fmt.Errorf("error getting code: %v groups: %v", data["code"], data["data"])
	}

	return groups, nil
}

func (r *GroupResource) createModel(data interface{}) *model.Group {
	// group := &bodies.Group{}
	// responseJson, _ := json.Marshal(responseData)
	// err := json.Unmarshal(responseJson, group)
	// if err != nil {
	// 	log.Printf("Error unmarshalling response body: %v", err)
	// 	return nil, err
	// }
	// groups = append(groups, &model.Group{
	// 	ID:          strconv.Itoa(group.ID),
	// 	Name:        group.Name,
	// 	Description: group.Description,
	// 	CreatedAt:   group.CreatedAt,
	// 	UpdatedAt:   group.UpdatedAt,
	// })
	id, _ := data.(map[string]interface{})["id"].(json.Number).Int64()
	return &model.Group{
		ID:          strconv.Itoa(int(id)),
		Name:        data.(map[string]interface{})["name"].(string),
		Description: data.(map[string]interface{})["description"].(string),
		CreatedAt:   data.(map[string]interface{})["created_at"].(string),
		UpdatedAt:   data.(map[string]interface{})["created_at"].(string),
	}
}
