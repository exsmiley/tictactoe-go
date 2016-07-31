package main

import (
	"fmt"
	"log"
    "os"
	"net/http"
	"encoding/json"
    "tictactoe-go/player"
)

type Board struct {
	Board [][]string
    Winner string
}

func executeNextMove(b *Board) string{
    // check if other player won the game
    over, winner := player.IsGameOver(b.Board)

    if over {
        fmt.Println(winner, "won the game!")
        return winner
    }

    move := player.GetNextMove(b.Board)
	b.Board[move.X][move.Y] = "o"

    // check for new winner
    over, winner = player.IsGameOver(b.Board)

    if over {
        fmt.Println(winner, "won the game!")
        return winner
    }
    return ""
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
    b.Winner = executeNextMove(&b)

    fmt.Println(b.Board)

    json.NewEncoder(w).Encode(b)
}

func main() {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "3000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
	http.HandleFunc("/move", moveHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}