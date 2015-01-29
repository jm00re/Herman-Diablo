package main

import (
	"math"
	"testing"
)

func TestMakeMove_1(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(true, board, 3)
	resultBoard := [14]uint8{4, 4, 4, 0, 5, 5, 1, 5, 4, 4, 4, 4, 4, 0}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_2(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(true, board, 2)
	resultBoard := [14]uint8{4, 4, 0, 5, 5, 5, 1, 4, 4, 4, 4, 4, 4, 0}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_3(t *testing.T) {
	board := [14]uint8{13, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(true, board, 0)
	resultBoard := [14]uint8{6, 5, 5, 5, 5, 5, 1, 5, 5, 5, 5, 5, 0, 0}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_4(t *testing.T) {
	board := [14]uint8{14, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(true, board, 0)
	resultBoard := [14]uint8{1, 6, 5, 5, 5, 5, 1, 5, 5, 5, 5, 5, 5, 0}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_5(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(false, board, 7)
	resultBoard := [14]uint8{4, 4, 4, 4, 4, 4, 0, 0, 5, 5, 5, 5, 4, 0}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_6(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(false, board, 9)
	resultBoard := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 0, 5, 5, 5, 1}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_7(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 13, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(false, board, 7)
	resultBoard := [14]uint8{5, 5, 5, 5, 5, 0, 0, 6, 5, 5, 5, 5, 5, 1}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_8(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 1, 0}
	newBoard := MakeMove(false, board, 12)
	resultBoard := [14]uint8{0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 1}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestMakeMove_9(t *testing.T) {
	board := [14]uint8{0, 0, 0, 0, 0, 1, 0, 4, 4, 4, 4, 4, 4, 0}
	newBoard := MakeMove(true, board, 5)
	resultBoard := [14]uint8{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 24}
	if newBoard != resultBoard {
		t.Error(board)
		t.Error(newBoard)
		t.Error(resultBoard)
	}
}

func TestAlphaBeta(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	AlphaBeta(true, board, math.MinInt32, math.MaxInt32, 11)
}

func TestDetermineMove(t *testing.T) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	DetermineMove(true, board, 10)
}

func BenchmarkAlphaBeta(b *testing.B) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	b.ResetTimer()
	AlphaBeta(true, board, math.MinInt32, math.MaxInt32, 11)
}

func BenchmarkDetermineMove(b *testing.B) {
	board := [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	b.ResetTimer()
	DetermineMove(true, board, 10)
}
