package router

import (
	"log"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/comment"
)

const commentPrefix = "/comment"

func getCommentRouters() http.Handler {
	router := http.NewServeMux()
	commentHdlr := &comment.Handler{}

	router.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		ok := commentHdlr.Create(w, r)
		if !ok {
			log.Println("invalid user ID")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	return router
}
