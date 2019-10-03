package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
