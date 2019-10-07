package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"golang.org/x/crypto/scrypt"
	"log"
	"net/http"
	"os"
)

// Sample User model, for which we'll create a table in Postgres
type User struct {
	gorm.Model // fields ID, CreatedAt, UpdatedAt, DeletedAt will be added
	Name string `gorm:"size:255"`
	Email string `gorm:"size:255"`
	PassHash []byte `gorm:"size:32"` //scrypt.Key(pass, salt, 65536, 8, 1, 32)
}

func testHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("TEST"))
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/test", testHandler)
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
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
	router := gin.Default()
	initializeRoutes(router)
	err = router.Run(":9898")
	if err != nil{
		panic("Error starting web server: " + err.Error())
	}
}
