package task

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// database connector
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

// GetByID TODO
func GetByID(id int) (*Task, error) {
	db, err := gorm.Open("mysql", "root:@tcp(db-server:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.LogMode(true)
	task := &Task{}
	if err := db.Find(&task, id).Error; err != nil {
		return nil, err
	}
	return task, nil
}
