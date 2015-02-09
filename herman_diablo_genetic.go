package main

import (
	"bufio"
	"fmt"
	//"github.com/davecheney/profile"
	"math"
	"math/rand"
	//"reflect"
	"os"
	"strconv"
	"strings"
	"time"

	//"time"
)

var weights [16]uint8
var useWeightedEval bool

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	//start := time.Now()

	player, board := ReadBoard()
	useWeightedEval = player
	rand.Seed(time.Now().Unix())
	// EvalBoard weights:
	// 0 - 13 weight of each pit
	// 14, 15 - Weight of moves left to win
	for true {
		for i := 0; i < 16; i++ {
			weights[i] = uint8(rand.Intn(256))
		}
		playBoard := board
		for !GameFinished(playBoard) {
			playBoard = MakeMove(player, playBoard, DetermineMove(player, playBoard, 6))
			player = !player
			useWeightedEval = !useWeightedEval
			//fmt.Println(playBoard)
		}
		if playBoard[6] > playBoard[13] {
			fmt.Println("Good", weights)
		} else {
			fmt.Println("Bad", weights)
		}
	}
	//fmt.Println("Best weights", weights)
	//fmt.Println(DetermineMove(player, board, 10))
	//fmt.Println(time.Since(start).Seconds())
}

func DetermineMove(player bool, board [14]uint8, depth uint8) uint8 {
	scoredMoves := make(map[uint8]int32)
	total := SumSide(true, board) + SumSide(false, board)
	if total <= 30 && total > 22 {
		depth += 1
	} else if total <= 22 && total > 18 {
		depth += 2
	} else if total <= 18 && total > 12 {
		depth += 3
	} else if total <= 12 {
		depth += 4
	}
	pMoves := PotentialMoves(player, board)
	for i := 0; i < 6; i++ {
		move := uint8(i)
		if player {
			if pMoves[move] {
				if ExtraTurn(player, board, move) {
					scoredMoves[move] = AlphaBeta(true, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
				} else {
					scoredMoves[move] = AlphaBeta(false, MakeMove(player, board, move), math.MinInt32, math.MaxInt32, depth)
				}
			}
		} else {
			if pMoves[move] {
				if ExtraTurn(player, board, move+7) {
					scoredMoves[move] = AlphaBeta(false, MakeMove(player, board, move+7), math.MinInt32, math.MaxInt32, depth)
				} else {
					scoredMoves[move] = AlphaBeta(true, MakeMove(player, board, move+7), math.MinInt32, math.MaxInt32, depth)
				}
			}
		}
	}
	// Converts map of scoredMoves into slice of sortedMoves
	//fmt.Println(scoredMoves)
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
		return sortedMoves[0] + 7
	}
	return sortedMoves[0]
	//if !player {
	//	return sortedMoves[0] + 1
	//}
	//return sortedMoves[0] + 1
}

func GameFinished(board [14]uint8) (finished bool) {
	var total1 uint8
	var total2 uint8
	if board[6] > 24 || board[13] > 24 {
		return true
	}
	for i := 0; i < 6; i++ {
		total1 += board[i]
		total2 += board[i+7]
	}
	return total1 == 0 || total2 == 0
}

func PotentialMoves(player bool, board [14]uint8) (moves [6]bool) {
	if player {
		for i := 0; i < 6; i++ {
			if board[i] != 0 {
				moves[i] = true
			}
		}
	} else {
		for i := 0; i < 6; i++ {
			if board[i+7] != 0 {
				moves[i] = true
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
		for i := 0; i < 6; i++ {
			sum += board[i]
		}
	} else {
		for i := 7; i < 13; i++ {
			sum += board[i]
		}
	}
	return
}

func AlphaBeta(player bool, board [14]uint8, alpha int32, beta int32, depth uint8) int32 {
	if depth == 0 || GameFinished(board) {
		if !useWeightedEval {
			return EvalBoard(player, board)
		} else {
			return WeightedEvalBoard(player, board, weights)
		}
	} else {
		if player {
			pMoves := PotentialMoves(player, board)
			for i := 0; i < 6; i++ {
				move := uint8(i)
				if pMoves[move] {
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
			}
			return alpha
		} else {
			pMoves := PotentialMoves(player, board)
			for i := 0; i < 6; i++ {
				move := uint8(i)
				if pMoves[move] {
					newMove := MakeMove(player, board, move+7)
					if ExtraTurn(player, board, move+7) {
						beta = Min(beta, AlphaBeta(false, newMove, alpha, beta, depth-1))
					} else {
						beta = Min(beta, AlphaBeta(true, newMove, alpha, beta, depth-1))
					}
					if beta <= alpha {
						break
					}
				}
			}
			return beta
		}
	}
}

func PlayerCount(board [14]uint8) (count int8) {
	for i := 0; i < 6; i++ {
		count += int8(board[i])
		count -= int8(board[i+7])
	}
	return
}

func EvalBoard(player bool, board [14]uint8) int32 {
	//If  player wins, return max/min value
	if board[6] > 24 {
		return math.MaxInt32
	} else if board[13] > 24 {
		return math.MinInt32
	} else {
		if player {
			return int32(board[6])*5 + int32(board[13])*-5 + int32(SumSide(player, board))
		} else {
			return int32(board[13])*-5 + int32(board[6])*5 - int32(SumSide(player, board))
		}
	}
}

func WeightedEvalBoard(player bool, board [14]uint8, weights [16]uint8) (total int32) {
	//If  player wins, return max/min value
	if player {
		// Single pits
		total += int32(board[0] * weights[0])
		total += int32(board[1] * weights[1])
		total += int32(board[2] * weights[2])
		total += int32(board[3] * weights[3])
		total += int32(board[4] * weights[4])
		total += int32(board[5] * weights[5])
		total += int32(board[6] * weights[6])
		total -= int32(board[7] * weights[7])
		total -= int32(board[8] * weights[8])
		total -= int32(board[9] * weights[9])
		total -= int32(board[10] * weights[10])
		total -= int32(board[11] * weights[11])
		total -= int32(board[12] * weights[12])
		total -= int32(board[13] * weights[13])
		// Distance to winning
		total += int32((24 - board[6]) * weights[14])
		total -= int32((24 - board[13]) * weights[15])
		return total
	} else {
		// Single pits
		total -= int32(board[0] * weights[0])
		total -= int32(board[1] * weights[1])
		total -= int32(board[2] * weights[2])
		total -= int32(board[3] * weights[3])
		total -= int32(board[4] * weights[4])
		total -= int32(board[5] * weights[5])
		total -= int32(board[6] * weights[6])
		total += int32(board[7] * weights[7])
		total += int32(board[8] * weights[8])
		total += int32(board[9] * weights[9])
		total += int32(board[10] * weights[10])
		total += int32(board[11] * weights[11])
		total += int32(board[12] * weights[12])
		total += int32(board[13] * weights[13])
		// Distance to winning
		total -= int32((24 - board[6]) * weights[14])
		total += int32((24 - board[13]) * weights[15])
		return total
	}
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
	if player {
		for count > 0 {
			if pos == 13 {
				pos = (pos + 1) % 14
			}
			newBoard[pos] += 1
			count -= 1
			pos = (pos + 1) % 14
		}
	} else {
		for count > 0 {
			if pos == 6 {
				pos = (pos + 1) % 14
			}
			newBoard[pos] += 1
			count -= 1
			pos = (pos + 1) % 14
		}
	}
	// Steal
	// Need to decrement to undo the last move
	pos = (pos - 1) % 14
	if player && pos >= 0 && pos <= 5 && newBoard[pos] == 1 {
		newBoard[pos] += newBoard[12-pos]
		newBoard[12-pos] = 0
	} else if !player && pos >= 7 && pos <= 12 && newBoard[pos] == 1 {
		newBoard[pos] += newBoard[12-pos]
		newBoard[12-pos] = 0
	}

	// If the game is finished, sum the sides and add to mancala
	if GameFinished(newBoard) {
		for i := 0; i < 6; i++ {
			newBoard[6] += newBoard[i]
			newBoard[13] += newBoard[i+7]
			newBoard[i] = 0
			newBoard[i+7] = 0
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
	p2mancala, _ := strconv.Atoi(scanner.Text())
	board[13] = uint8(p2mancala)
	// Read Player 2 board
	scanner.Scan()
	for index, item := range strings.Split(scanner.Text(), " ") {
		temp, _ := strconv.Atoi(item)
		board[index+7] = uint8(temp)
	}
	return
}
