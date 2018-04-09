// Copyright © 2018 budougumi0617 All Rights Reserved.

package routing

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/budougumi0617/react-golang/backend/task"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// GetTaskRouter returns simple JSON API server
//  curl -D - -X GET http://localhost:8888/tasks/5
//  HTTP/1.1 200 OK
//  Date: Mon, 09 Apr 2018 15:10:11 GMT
//  Content-Length: 141
//  Content-Type: text/plain; charset=utf-8
//
//  task &{ID:5 Title:Dummy Data5 Body:long long long long long CreatedAt:2018-04-09 14:20:21 +0000 UTC UpdatedAt:2018-04-09 14:20:21 +0000 UTC}
func GetTaskRouter() chi.Router {
	router := chi.NewRouter()
	// Set output for logging.
	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: newLogger(),
		},
	)
	router.Use(middleware.Logger)
	// 	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/tasks/{id}", getTaskByID)
	// 	router.Post("/todos", TodoCreate)
	// TODO Need to set default error response
	return router
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "chi-log: ", log.Lshortfile)
}

func getTaskByID(resp http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Println("could not get id")
		return
	}
	task, err := task.GetByID(id)
	if err != nil {
		resp.WriteHeader(404)
		resp.Write([]byte(err.Error()))
	}

	fmt.Fprintf(resp, "task %+v\n", task)
}
