package comment

import (
	"log"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/middleware"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) bool {
	userId, ok := r.Context().Value(middleware.AuthUserID).(string)
	log.Println("creating comment for user:", userId)
	return ok
}
