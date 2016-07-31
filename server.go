package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Board struct {
	Board [][]string
}

func executeNextMove(b *Board) {
	b.Board[0][0] = "o"
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
    b := Board{}

    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
    err := json.NewDecoder(r.Body).Decode(&b)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    // mutates b and makes the next move
    executeNextMove(&b)

    fmt.Println(b.Board)

    json.NewEncoder(w).Encode(b)
}

func main() {
	port := "3000"
	http.HandleFunc("/move", moveHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}