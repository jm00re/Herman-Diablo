package main

import (
	"bufio"
	"fmt"
	"github.com/davecheney/profile"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	player, board := ReadBoard()
	fmt.Println(DetermineMove(player, board, 11))
}

func DetermineMove(player bool, board [14]uint8, depth uint8) uint8 {
	scoredMoves := make(map[uint8]int32)
	total := SumSide(true, board) + SumSide(false, board)
	if total <= 30 && total > 24 {
		depth += 1
	} else if total <= 24 && total > 18 {
		depth += 2
	} else if total <= 18 && total > 12 {
		depth += 3
	} else if total <= 12 {
		depth += 4
	}
	for _, move := range PotentialMoves(player, board) {
		if player {
			if ExtraTurn(player, board, move) {
				scoredMoves[move] = AlphaBeta(true, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
			} else {
				scoredMoves[move] = AlphaBeta(false, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
			}
		} else {
			if ExtraTurn(player, board, move) {
				scoredMoves[move] = AlphaBeta(false, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
			} else {
				scoredMoves[move] = AlphaBeta(true, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
			}
		}
	}
	// Converts map of scoredMoves into slice of sortedMoves
	var sortedMoves []uint8
	for len(scoredMoves) != 0 {
		var tempKey uint8
		var tempVal int32
		// Selects an item from the map
		for key, val := range scoredMoves {
			tempKey = key
			tempVal = val
			break
		}
		for key, val := range scoredMoves {
			if player {
				if val > tempVal {
					tempKey = key
					tempVal = val
				}
			} else {
				if val < tempVal {
					tempKey = key
					tempVal = val
				}
			}
		}
		sortedMoves = append(sortedMoves, tempKey)
		delete(scoredMoves, tempKey)
	}
	if !player {
		return sortedMoves[0] - 6
	}
	return sortedMoves[0] + 1
}

func GameFinished(board [14]uint8) (finished bool) {
	finished = true
	if board[6] > 24 || board[13] > 24 {
		return
	}
	for _, content := range board[0:6] {
		if content != 0 {
			finished = false
		}
	}
	if finished {
		return
	}
	finished = true
	for _, content := range board[7:13] {
		if content != 0 {
			finished = false
		}
	}
	return
}

func PotentialMoves(player bool, board [14]uint8) (moves []uint8) {
	if player {
		for move, content := range board[0:6] {
			if content != 0 {
				moves = append(moves, uint8(move))
			}
		}
	} else {
		for move, content := range board[7:13] {
			if content != 0 {
				moves = append(moves, uint8(move+7))
			}
		}
	}
	return
}

func Max(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int32, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func SumSide(player bool, board [14]uint8) (sum uint8) {
	if player {
		for _, content := range board[0:6] {
			sum += content
		}
	} else {
		for _, content := range board[7:13] {
			sum += content
		}
	}
	return
}

func AlphaBeta(player bool, board [14]uint8, alpha int32, beta int32, depth uint8) int32 {
	if depth == 0 || GameFinished(board) {
		return EvalBoard(board)
	} else {
		if player {
			for _, move := range PotentialMoves(player, board) {
				newMove := MakeMove(player, board, move)
				if ExtraTurn(player, board, move) {
					alpha = Max(alpha, AlphaBeta(true, newMove, alpha, beta, depth-1))
				} else {
					alpha = Max(alpha, AlphaBeta(false, newMove, alpha, beta, depth-1))
				}
				if beta <= alpha {
					break
				}
			}
			//fmt.Println(alpha)
			return alpha
		} else {
			for _, move := range PotentialMoves(player, board) {
				newMove := MakeMove(player, board, move)
				if ExtraTurn(player, board, move) {
					beta = Min(beta, AlphaBeta(false, newMove, alpha, beta, depth-1))
				} else {
					beta = Min(beta, AlphaBeta(true, newMove, alpha, beta, depth-1))
				}
				if beta <= alpha {
					break
				}
			}
			return beta
		}
	}
}

func EvalBoard(board [14]uint8) int32 {
	//If  player wins, return max/min value
	if GameFinished(board) {
		if board[6] > board[13] {
			return math.MaxInt32
		} else if board[13] > board[6] {
			return math.MinInt32
		} else {
			return 0
		}
	} else {
		//If a player has more than 24, they win
		if board[6] > 24 {
			return math.MaxInt32
		} else if board[13] > 24 {
			return math.MinInt32
		}
	}
	return int32((board[6]*5)+(SumSide(true, board))) - int32((board[13]*5)+(SumSide(false, board)))
}

func ExtraTurn(player bool, board [14]uint8, move uint8) bool {
	// If last move ends on players mancala, ExtraTurn == true
	if player {
		return (board[move]+move)%14 == 6
	} else {
		return (board[move]+move)%14 == 13
	}
}

func MakeMove(player bool, board [14]uint8, move uint8) (newBoard [14]uint8) {
	// Basically this function is a mess but it's all cool
	count := board[move]
	pos := (move + 1) % 14
	for i := 0; i < 14; i++ {
		newBoard[i] = board[i]
	}
	newBoard[move] = 0
	for count > 0 {
		if player {
			if pos == 13 {
				pos = (pos + 1) % 14
			}
			newBoard[pos] += 1
			count -= 1
			if count == 0 && pos >= 0 && pos <= 5 && newBoard[pos] == 1 {
				newBoard[pos] += newBoard[12-pos]
				newBoard[12-pos] = 0
			}
			pos = (pos + 1) % 14
		} else {
			if pos == 6 {
				pos = (pos + 1) % 14
			}
			newBoard[pos] += 1
			count -= 1
			if count == 0 && pos >= 7 && pos <= 12 && newBoard[pos] == 1 {
				newBoard[pos] += newBoard[12-pos]
				newBoard[12-pos] = 0
			}
			pos = (pos + 1) % 14
		}
	}
	// If the game is finished, sum the sides and add to mancala
	if GameFinished(newBoard) {
		for index, content := range newBoard[0:6] {
			newBoard[6] += content
			newBoard[index] = 0
		}
		for index, content := range newBoard[7:13] {
			newBoard[13] += content
			newBoard[index+7] = 0
		}
	}
	return
}

func ReadBoard() (player bool, board [14]uint8) {
	scanner := bufio.NewScanner(os.Stdin)
	// Read player number
	scanner.Scan()
	tempPlayer, _ := strconv.Atoi(scanner.Text())
	if tempPlayer == 1 {
		player = true
	} else {
		player = false
	}
	// Read Player 1 mancala
	scanner.Scan()
	p1mancala, _ := strconv.Atoi(scanner.Text())
	board[6] = uint8(p1mancala)
	// Read player 1 board
	scanner.Scan()
	//var p1board [14]uint8
	for index, item := range strings.Split(scanner.Text(), " ") {
		temp, _ := strconv.Atoi(item)
		board[index] = uint8(temp)
	}
	// Read Player 2 mancala
	scanner.Scan()
	board[13] = uint8(p1mancala)
	// Read Player 2 board
	scanner.Scan()
	for index, item := range strings.Split(scanner.Text(), " ") {
		temp, _ := strconv.Atoi(item)
		board[index+7] = uint8(temp)
	}
	return
}
