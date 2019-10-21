package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/william-carvalho/go-crud-mysql-rest/driver"
	models "github.com/william-carvalho/go-crud-mysql-rest/models"
	repository "github.com/william-carvalho/go-crud-mysql-rest/repository"
	post "github.com/william-carvalho/go-crud-mysql-rest/repository/post"
)

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

type Post struct {
	repo repository.PostRepo
}

func (p *Post) getByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.getByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	} else {
		respondwithJSON(w, http.StatusOK, payload)
	}

}

func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	fmt.Println(newID)
	fmt.Println(&post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}

}

func (p *Post) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		respondwithJSON(w, http.StatusOK, payload)
	}

}

func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
	}

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
