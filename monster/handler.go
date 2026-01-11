package monster

import (
	"errors"
	"log"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/entity"
	"github.com/dreamsofcode-io/nethttp/model"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) model.MonsterRes {
	log.Println("received request to create a monster")
	return model.MonsterRes{Desc: "Monster created!"}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) (entity.Monster, error) {
	log.Println("handling READ request - Method:", r.Method)
	monster, exists := loadMonsters()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return entity.Monster{}, errors.New("FindByID error")

	}
	return monster, nil
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling UPDATE request - Method:", r.Method)
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	log.Println("received DELETE request for monster")
}

func (h *Handler) PatchByID(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
}
