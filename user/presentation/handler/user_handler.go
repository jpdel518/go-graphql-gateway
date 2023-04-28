package handler

import (
	"encoding/json"
	"github.com/jpdel518/go-graphql-gateway/user/domain/model"
	"github.com/jpdel518/go-graphql-gateway/user/usecase"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Handler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *Handler {
	return &Handler{usecase}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		h.Fetch(w, r)
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	// get query parameters
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		log.Println(err)
		offset = 0
	}
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil {
		log.Println(err)
		num = 10
	}
	name := r.URL.Query().Get("name")
	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
	if err != nil {
		log.Println(err)
		sort = 1
	}
	group, err := strconv.Atoi(r.URL.Query().Get("group"))
	if err != nil {
		log.Println(err)
		group = 0
	}

	// fetch site data
	users, err := h.usecase.Fetch(r.Context(), offset, num, name, group, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3000, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: users}))
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get path parameters
	sub := strings.TrimPrefix(r.URL.Path, "/user")
	id, err := strconv.Atoi(filepath.Base(sub))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3101, Data: err.Error()}))
	}

	// fetch site data
	user, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3100, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: user}))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if r.Header.Get("Content-Type") != "application/json" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")

	// get parameters
	var user *model.User
	var avatar multipart.File
	var fh *multipart.FileHeader
	if r.Header.Get("Content-Type") != "application/json" {
		// multipart/form-data or application/x-www-form-urlencoded
		log.Printf("Content-Type: %v", r.Header.Get("Content-Type"))
		// err := r.ParseMultipartForm(32 << 20)
		// if err != nil {
		// 	log.Println(err)
		// 	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3201, Data: err.Error()}))
		// 	return
		// }
		log.Printf("response: %v", r.MultipartForm)
		firstName := r.FormValue("first_name")
		log.Printf("firstName: %v", r.FormValue("first_name"))
		lastName := r.FormValue("last_name")
		log.Printf("lastName: %v", r.FormValue("last_name"))
		address := r.FormValue("address")
		log.Printf("address: %v", r.FormValue("address"))
		email := r.FormValue("email")
		log.Printf("email: %v", r.FormValue("email"))
		age, err := strconv.Atoi(r.FormValue("age"))
		log.Printf("age: %v", r.FormValue("age"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3201, Data: err.Error()}))
			return
		}
		groupID, _ := strconv.Atoi(r.FormValue("group_id"))
		avatar, fh, err = r.FormFile("avatar")
		if err != nil {
			// avatarファイルは必須ではない
			log.Printf("failed to get avatar file %v", err)
		} else {
			log.Printf("avatar file: %v", fh.Filename)
		}
		user = &model.User{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Address:   address,
			Email:     email,
			GroupID:   groupID,
		}
	} else {
		// application/json
		user = &model.User{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3204, Data: err.Error()}))
			return
		}
		err = json.Unmarshal(body, user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3205, Data: err.Error()}))
			return
		}
	}

	// create site
	data, err := h.usecase.Create(r.Context(), user, avatar, fh)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3200, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: data}))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get parameters
	var user *model.User
	if r.Header.Get("Content-Type") != "application/json" {
		// multipart/form-data or application/x-www-form-urlencoded
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3301, Data: err.Error()}))
			return
		}
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		address := r.FormValue("address")
		email := r.FormValue("email")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3302, Data: err.Error()}))
			return
		}
		groupID, _ := strconv.Atoi(r.FormValue("group_id"))

		user = &model.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Address:   address,
			Email:     email,
			GroupID:   groupID,
		}
	} else {
		// application/json
		user = &model.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3206, Data: err.Error()}))
			return
		}
	}

	// update site
	data, err := h.usecase.Update(r.Context(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3300, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: data}))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get query parameters
	sub := strings.TrimPrefix(r.URL.Path, "/user")
	id, err := strconv.Atoi(filepath.Base(sub))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3401, Data: err.Error()}))
		return
	}

	// delete site
	err = h.usecase.Delete(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3400, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: "success"}))
}
