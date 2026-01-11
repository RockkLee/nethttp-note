package main

import (
	"fmt"
	// "log"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/middleware"
	"github.com/dreamsofcode-io/nethttp/router"
)

// func handleOther(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Received a non domain request")
// 	w.Write([]byte("Hello, stranger..."))
// }

// func handle(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Received a request at my domain")
// 	w.Write([]byte("Hello, Domain name!"))
// }

func main() {
	rootRouter := http.NewServeMux()
	router.LoadRouters(rootRouter)
	// rootRouter.HandleFunc("/", handleOther)
	// rootRouter.HandleFunc("dreamsofcode.foo/", handle)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
		middleware.IsAuthenticated,
		middleware.CheckPermissions,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: middlewareStack(rootRouter),
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
