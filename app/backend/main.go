package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

// Sample User model, for which we'll create a table in Postgres
type User struct {
	gorm.Model // fields ID, CreatedAt, UpdatedAt, DeletedAt will be added
	Name string `gorm:"size:255"`
}

func main() {
	log.Println("Starting app")

	// Open Database Connection
	// Postgres User and Database name are both "postgres"
	log.Println("DB Pass:", os.Getenv("postgres_password"))
	db, err := gorm.Open("postgres", "host=localhost port=5432 sslmode=disable user=postgres dbname=postgres password=" + os.Getenv("postgres_password"))
	defer db.Close()
	if err!=nil{
		panic("DB Open Error: " + err.Error())
	} 

	// Create or migrate tables, as needed
	db.Debug().AutoMigrate(&User{}) 
	
	// Setup HTTP handler
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err = http.ListenAndServe(":9898", nil)
	if err!=nil{
		panic("Error starting web server: " + err.Error())
	} 
}
