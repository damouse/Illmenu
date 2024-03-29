package illmenu

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
)

// This is the actual server

// Testing:
//   https://github.com/stretchr/testify
//   https://golang.org/pkg/net/http/httptest/

// Version 1 of the api
type APIServerOne struct {
	Db  *gorm.DB
	Api *rest.Api
}

// Create a new instance of the given server/handler wrapper
// If a connection is passed in, it will be used as the database
func NewServer(db *gorm.DB) *APIServerOne {
	s := &APIServerOne{Db: db}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/query", s.Query),
	)

	panicOnError(err)
	api.SetApp(router)
	s.Api = api

	return s
}

func (s *APIServerOne) Serve() {
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", s.Api.MakeHandler()))
}

// Given a menu, search through the contents and find matching images
func (s *APIServerOne) Query(w rest.ResponseWriter, r *rest.Request) {
	log.Printf("Query HTTP call")
	w.WriteJson("")
}
