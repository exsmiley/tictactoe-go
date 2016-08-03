package player

import (
    "fmt"
    "math/rand"
)

// IsGameOver returns true if the game is over and the name of the winner
func IsGameOver(board [][]string) (bool, string){
    for i := 0; i < 3; i++ {
        if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != ""{
            return true, board[i][0]
        } else if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != ""{
            return true, board[0][i]
        }
    }
    if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != ""{
        return true, board[0][0]
    }
    if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] != ""{
        return true, board[0][2]
    }
    // search for cat's game
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                return false, ""
            }
        }
    }
    return true, "cat"
}

type Move struct {
    X int
    Y int
}

// getAllMoves returns an array of all possible remaining moves
func getAllMoves(board [][]string) []Move{
    moves := []Move{}

    over, _ := IsGameOver(board)

    if over {
        return moves
    }

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                move := Move{i,j}
                moves = append(moves, move)
            }
        }
    }

    // want the moves to be scrambled for some fun
    scrambled := make([]Move, len(moves))

    perm := rand.Perm(len(moves))

    for i, v := range perm {
        scrambled[v] = moves[i]
    }

    return scrambled
}

// doMove returns a new board with the move executed by the player
func doMove(board [][]string, move Move, player string) [][]string{
    newBoard := [][]string{{"","",""},{"","",""},{"","",""}}
    copy(newBoard[0], board[0])
    copy(newBoard[1], board[1])
    copy(newBoard[2], board[2])
    newBoard[move.X][move.Y] = player
    return newBoard
}

func getNextPlayer(player string) string {
    if player == "x" {
        return "o"
    } else {
        return "x"
    }
}

func evaluateBoard(board [][]string, player string, depth int) int {
    over, winner := IsGameOver(board)

    if over && winner == player {
        return 10 + depth
    } else if over && winner == getNextPlayer(player) {
        return -10 - depth
    }

    return 0
}

type ABAction struct {
    value int
    move Move
}

func alphaBetaHelper(board [][]string, alpha int, beta int, player string, depth int, givenMove Move, results chan ABAction) {
    newAlpha := -beta
    newBeta := -alpha

    actionMove := Move{-1,-1}
    result := ABAction{newAlpha, actionMove}

    over, _ := IsGameOver(board)

    if over || depth == 0{
        result.value = evaluateBoard(board, player, depth)
        result.move = givenMove
        results <- result
        return
    }

    moves := getAllMoves(board)

    counter := len(moves)

    options := make(chan ABAction)

    for _, move := range moves {
        nextBoard := doMove(board, move, player)
        go alphaBetaHelper(nextBoard, newAlpha, newBeta, getNextPlayer(player), depth - 1, move, options)
    }

    for action := range options {
        val := -action.value
        centerMove := Move{1,1}
        if val > newAlpha {
            newAlpha = val
            result = action
            if newAlpha > newBeta {
                results <- action
                close(options)
                return
            }
        } else if val == newAlpha && action.move == centerMove {
            result = action
        }
        counter--
        if counter == 0 {
            // no more options so return best option
            close(options)
        }
    }

    results <- result
}

func GetNextMove(board [][]string) Move{
    alpha := -10000
    beta := 10000
    blocker := make(chan struct{})
    results := make(chan ABAction)
    move := Move{-1,-1}
    go alphaBetaHelper(board, alpha, beta, "o", 10, move, results)
    for action := range results {
        fmt.Println("I got a score of", action.value, "for", action.move)
        close(results)
        move = action.move
        close(blocker)
    }
    <- blocker
    return move
}