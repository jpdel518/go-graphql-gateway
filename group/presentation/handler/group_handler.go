package handler

import (
	"encoding/json"
	"github.com/jpdel518/go-graphql-gateway/group/domain/model"
	"github.com/jpdel518/go-graphql-gateway/group/usecase"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Handler struct {
	usecase usecase.GroupUsecase
}

func NewGroupHandler(usecase usecase.GroupUsecase) *Handler {
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

	// fetch site data
	groups, err := h.usecase.Fetch(r.Context(), offset, num, name, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3000, Data: err.Error()}))
		return
	}
	log.Printf("returned value: %+v", groups)

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: groups}))
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get path parameters
	sub := strings.TrimPrefix(r.URL.Path, "/group")
	id, err := strconv.Atoi(filepath.Base(sub))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3101, Data: err.Error()}))
	}

	// fetch site data
	group, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3100, Data: err.Error()}))
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 2000, Data: group}))
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
	var group *model.Group
	if r.Header.Get("Content-Type") != "application/json" {
		// multipart/form-data or application/x-www-form-urlencoded
		name := r.FormValue("name")
		description := r.FormValue("description")
		group = &model.Group{
			Name:        name,
			Description: description,
		}
	} else {
		// application/json
		group = &model.Group{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3204, Data: err.Error()}))
			return
		}
		err = json.Unmarshal(body, group)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3205, Data: err.Error()}))
			return
		}
	}

	// create site
	data, err := h.usecase.Create(r.Context(), group)
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
	var group *model.Group
	if r.Header.Get("Content-Type") != "application/json" {
		// multipart/form-data or application/x-www-form-urlencoded
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3301, Data: err.Error()}))
			return
		}
		name := r.FormValue("name")
		description := r.FormValue("description")

		group = &model.Group{
			ID:          id,
			Name:        name,
			Description: description,
		}
	} else {
		// application/json
		group = &model.Group{}
		err := json.NewDecoder(r.Body).Decode(group)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(CreateResponseJson(&ApiRequestResponse{Code: 3206, Data: err.Error()}))
			return
		}
	}

	// update site
	data, err := h.usecase.Update(r.Context(), group)
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
	sub := strings.TrimPrefix(r.URL.Path, "/group")
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
