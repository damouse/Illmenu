package illmenu

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// This is the actual server

// Testing:
//   https://github.com/stretchr/testify
//   https://golang.org/pkg/net/http/httptest/

// Version 1 of the api
type ServerApiV1 struct{}

func Serve() {
	v1 := ServerApiV1{}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/query", v1.Query),
	)

	panicOnError(err)
	api.SetApp(router)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

// Given a menu, search through the contents and find matching images
func (s *ServerApiV1) Query(w rest.ResponseWriter, r *rest.Request) {
	log.Printf("Query HTTP call")
}
