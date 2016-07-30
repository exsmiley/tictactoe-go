package main

import (
	_ "fmt"
	"log"
	"net/http"
)


func main() {
	port := "3000"
	
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}