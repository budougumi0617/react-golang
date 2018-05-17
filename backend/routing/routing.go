// Copyright © 2018 budougumi0617 All Rights Reserved.

package routing

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	router.Get("/tasks", getAllTasks)
	router.HandleFunc("/tasks/{id}", getTaskByID)
	router.Post("/tasks", addTask)
	// TODO Need to set default error response
	return router
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "chi-log: ", log.Lshortfile)
}

func addTask(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		log.Println("parse request error")
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		return
	}
	var t struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.Unmarshal(body, &t); err != nil {
		resp.WriteHeader(http.StatusBadRequest) // unprocessable entity
		resp.Write([]byte(err.Error()))
		return
	}

	result, err := task.Create(t.Title, t.Body)
	if err != nil {
		log.Println("Save data error")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(resp).Encode(result); err != nil {
		log.Println("could not marshal JSON")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
	}
}

func getTaskByID(resp http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Println("could not get id")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}
	task, err := task.GetByID(id)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(task)
	if err != nil {
		log.Println("could not marshal JSON from task")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "task %+v\n", string(b))
}

func getAllTasks(resp http.ResponseWriter, r *http.Request) {
	tasks, err := task.All()
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(resp).Encode(tasks); err != nil {
		log.Println("could not marshal JSON")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
	}
}
