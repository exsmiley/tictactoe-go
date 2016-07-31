package player

import (
    "fmt"
)

// IsGameOver returns true if the game is over and the name of the winner
func IsGameOver(board [][]string) (bool, string){
    for i := 0; i < 3; i++ {
        if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
            return true, board[i][0]
        } else if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
            return true, board[0][i]
        }
    }
    if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
        return true, board[0][0]
    }
    if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
        return true, board[0][0]
    }
    if len(getAllMoves(board)) == 0 {
        return true, "cat"
    }
    return false, ""
}

type Move struct {
    X int
    Y int
}

// getAllMoves returns an array of all possible remaining moves
func getAllMoves(board [][]string) []Move{
    moves := []Move{}

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                move := Move{i,j}
                moves = append(moves, move)
            }
        }
    }

    return moves
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

func evaluateBoard(board [][]string, player string) int {
    val := 0

    over, winner := IsGameOver(board)

    if over && winner == player {
        fmt.Println("I won")
        val += 1000
    } else if over && winner == "cat" {
        fmt.Println("I tied")
        val += 0
    } else if over {
        fmt.Println("I lost")
        val -= 1000
    }

    // overall value of certain places
    val += potentialOf3(player, board[0][0], board[0][1], board[0][2])
    val += potentialOf3(player, board[1][0], board[1][1], board[1][2])
    val += potentialOf3(player, board[2][0], board[2][1], board[2][2])
    val += potentialOf3(player, board[0][0], board[1][0], board[2][0])
    val += potentialOf3(player, board[0][1], board[1][1], board[2][1])
    val += potentialOf3(player, board[0][2], board[1][2], board[2][2])
    val += potentialOf3(player, board[0][0], board[1][1], board[2][2])
    val += potentialOf3(player, board[2][0], board[1][1], board[0][2])

    if board[1][1] == player {
        val += 100
    }

    return val
}

func addPotential(win int, lose int, maybe int) int {
    if lose == 2 && win == 1 {
        return 40
    } else if lose == 2 && maybe == 1 {
        return -20
    } else if win == 2 && lose == 1 {
        return 5
    } else if win == 1 && lose == 1 {
        return 5
    } else if lose == 3 {
        return -100
    } else {
        return 0
    }
}

func potentialOf3(player string, spot1 string, spot2 string, spot3 string) int{
    win, lose, maybe := 0, 0, 0

    if spot1 == player {
        win += 1
    } else if spot1 == "" {
        maybe += 1
    } else {
        lose += 1
    }

    if spot2 == player {
        win += 1
    } else if spot2 == "" {
        maybe += 1
    } else {
        lose += 1
    }

    if spot3 == player {
        win += 1
    } else if spot3 == "" {
        maybe += 1
    } else {
        lose += 1
    }

    return addPotential(win, lose, maybe)
}

func alphaBetaHelper(board [][]string, alpha int, beta int, player string) (int, Move) {
    newAlpha := -beta
    newBeta := -alpha

    actionAlpha := newAlpha
    actionMove := Move{-1,-1}
    for _, move := range getAllMoves(board) {
        nextBoard := doMove(board, move, getNextPlayer(player))
        val, _ := alphaBetaHelper(nextBoard, newAlpha, newBeta, getNextPlayer(player))
        if val > newAlpha {
            newAlpha = val
            actionAlpha = newAlpha
            actionMove = move
            if newAlpha > newBeta {
                fmt.Println("Weird situation", newAlpha, newBeta, actionMove, player)
                return actionAlpha, actionMove
            }
        }
    }

    if len(getAllMoves(board)) == 0 {
        return evaluateBoard(board, player), actionMove
    }
    fmt.Println(actionAlpha, actionMove, player)
    return actionAlpha, actionMove
}

func GetNextMove(board [][]string) Move{
    alpha := -10000
    beta := 10000
    _, move := alphaBetaHelper(board, alpha, beta, "o")
    return move
}