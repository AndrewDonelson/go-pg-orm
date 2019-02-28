package controllers

import (
	goPg "github.com/AndrewDonelson/go-pg-orm"
	"net/http"
	"github.com/AndrewDonelson/example/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"strconv"
	"errors"
)

// Decoder is use to decode the schema
var Decoder = schema.NewDecoder()

type Handler struct {
	db *goPg.Model
}

func New(db *goPg.Model) *Handler  {
	return &Handler{db:db}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	articles := []model.Article{}
	err := h.db.GetAllModels(&articles)
	if err != nil {
		return
	}
	response, _ := json.Marshal(articles)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}


func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	article := model.Article{}

	isOK := decode(r, &article)
	if !isOK {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article.BeforeInsert()
	errU := h.db.SaveModel(&article)
	if errU != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["id"]

	id, err := checkID(idS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := model.Article{ID: id}
	errG := h.db.GetModel(&article)
	if errG != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) Delete (w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["id"]

	id, err := checkID(idS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := model.Article{ID: id}
	errG := h.db.DeleteModel(&article)
	if errG != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["id"]

	id, err := checkID(idS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := model.Article{ID: id}
	h.db.GetModel(&article)

	isOK := decode(r, &article)
	if !isOK {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	article.BeforeUpdate()
	errU := h.db.UpdateModel(&article)
	if errU != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func checkID (idS string) (int, error) {
	if idS == "" {
		return 0, errors.New("missing ID")
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func decode(req *http.Request, model interface{}) bool {
	err := req.ParseForm()
	if err != nil {
		return false
	}

	if err := Decoder.Decode(model, req.PostForm); err != nil {
		return false
	}

	return true
}


