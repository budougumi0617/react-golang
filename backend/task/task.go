package task

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	// database connector
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// BASE defines mysql URL.
	BASE = "root:@tcp(db-server:3306)"

	// ENDPOINT defines database URL
	ENDPOINT = "/todo?charset=utf8&parseTime=True&loc=Local"
)

// Task is a basic model.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DueDate
	// Completed
}

// Tasks is Task list.
type Tasks []Task

// Create inserts the Task in DB.
func Create(title, body string) (*Task, error) {
	task := Task{
		Title: title,
		Body:  body,
	}
	url := BASE + ENDPOINT
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.LogMode(true)
	if err := db.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// Delete removes the Task in DB.
func Delete(task Task) error {
	if task.ID == 0 {
		return errors.New("must need primary key, but ID was " + strconv.Itoa(task.ID))
	}
	url := BASE + ENDPOINT
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.LogMode(true)
	if err := db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

// GetByID returns a task by ID.
func GetByID(id int) (*Task, error) {
	url := BASE + ENDPOINT
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.LogMode(true)
	var task Task
	if err := db.Find(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// All returns tasks
func All() ([]Task, error) {
	url := BASE + ENDPOINT
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.LogMode(true)
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
