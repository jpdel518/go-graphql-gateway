package handler

import (
	"encoding/json"
	"github.com/jpdel518/go-graphql-gateway/user/usecase"
	"log"
	"net/http"
)

func NewHandler(userUsecase usecase.UserUsecase) {
	userHandler := NewUserHandler(userUsecase)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})
	http.HandleFunc("/api/users", userHandler.Handle)
	http.HandleFunc("/api/users/", userHandler.Handle)
	http.HandleFunc("/api/users/get-by-id/", userHandler.GetById)

	_ = http.ListenAndServe(":8080", nil)
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
