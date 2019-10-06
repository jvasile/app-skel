package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Pringln("Starting up")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
