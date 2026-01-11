package router

import (
	"fmt"
	"net/http"
)

func LoadRouters(router *http.ServeMux) {
	router.Handle(
		fmt.Sprintf("%s/", monsterPrefix),
		http.StripPrefix(monsterPrefix, getMonsterRouters()),
	)
	router.Handle(
		fmt.Sprintf("%s/", commentPrefix),
		http.StripPrefix(commentPrefix, getCommentRouters()),
	)
}
