package main

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	// database connector
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Task TODO
// http://doc.gorm.io/models.html
type Task struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(db-server:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.LogMode(true)
	log.Println("has tasks table?" + strconv.FormatBool(db.HasTable(&Task{})))
	log.Println("has tasks table?" + strconv.FormatBool(db.HasTable("tasks")))
	var tasks []Task
	db.Find(&tasks)
	for k, v := range tasks {
		log.Printf("task[%d]  = %+v\n", k, v)
	}
	db.Where("body LIKE ?", "%data2").Find(&tasks)
	log.Println("---------------------LIKE-------------------")
	for k, v := range tasks {
		log.Printf("task[%d]  = %+v\n", k, v)
	}

}
