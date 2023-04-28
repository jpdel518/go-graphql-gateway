package handler

import (
	"encoding/json"
	"github.com/jpdel518/go-graphql-gateway/group/usecase"
	"log"
	"net/http"
)

func NewHandler(userUsecase usecase.GroupUsecase) {
	userHandler := NewGroupHandler(userUsecase)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})
	http.HandleFunc("/api/groups", userHandler.Handle)
	http.HandleFunc("/api/groups/", userHandler.Handle)
	http.HandleFunc("/api/groups/get-by-id/", userHandler.GetById)

	_ = http.ListenAndServe(":8081", nil)
}

// ApiRequestResponse response json
type ApiRequestResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// CreateResponseJson create response data as json
func CreateResponseJson(a *ApiRequestResponse) []byte {
	js, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("create response json error: %v", err)
	}
	return js
}
