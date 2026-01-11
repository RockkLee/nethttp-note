package router

import (
	"encoding/json"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/monster"
)

const monsterPrefix = "/monster"

func getMonsterRouters() http.Handler {
	router := http.NewServeMux()
	monsterHdlr := &monster.Handler{}

	router.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		monsterRes := monsterHdlr.Create(w, r)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(monsterRes.Desc))
	})

	router.HandleFunc("PUT /update/{id}", func(w http.ResponseWriter, r *http.Request) {
		monsterHdlr.UpdateByID(w, r)
	})

	router.HandleFunc("GET /get/{id}", func(w http.ResponseWriter, r *http.Request) {
		monsterEntity, err := monsterHdlr.FindByID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(monsterEntity)
	})

	router.HandleFunc("DELETE /delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		monsterHdlr.DeleteByID(w, r)
	})

	return router
}
