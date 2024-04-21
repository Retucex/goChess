package main

import "fmt"

const (
	ansiEscape = "\u001b"
	resetModes = ansiEscape + "[0m"
	defaultBg  = ansiEscape + "[49m"
	defaultFg  = ansiEscape + "[39m"
	blackBg    = ansiEscape + "[40m"
	whiteBg    = ansiEscape + "[47m"
	whiteFg    = ansiEscape + "[97m"
)

const (
	Empty = iota
	Pawn
	Rook
	Bishop
	Knight
	Queen
	King
)

const (
	White = 0b0100_0000
	Black = 0b0000_0000
)

const (
	Occupied   = 0b1000_0000
	Unoccupied = 0b0000_0000
)

func initBoard() [8][8]byte {
	board := [8][8]byte{}

	board[0][0] = Occupied | Black | Rook
	board[1][0] = Occupied | Black | Knight
	board[2][0] = Occupied | Black | Bishop
	board[3][0] = Occupied | Black | Queen
	board[4][0] = Occupied | Black | King
	board[5][0] = Occupied | Black | Bishop
	board[6][0] = Occupied | Black | Knight
	board[7][0] = Occupied | Black | Rook

	for i := 0; i < 8; i++ {
		board[i][1] = Occupied | Black | Pawn
	}

	board[0][7] = Occupied | White | Rook
	board[1][7] = Occupied | White | Knight
	board[2][7] = Occupied | White | Bishop
	board[3][7] = Occupied | White | Queen
	board[4][7] = Occupied | White | King
	board[5][7] = Occupied | White | Bishop
	board[6][7] = Occupied | White | Knight
	board[7][7] = Occupied | White | Rook

	for i := 0; i < 8; i++ {
		board[i][6] = Occupied | White | Pawn
	}

	return board
}

func printBoardBinary(board *[8][8]byte) {
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("[%08b]", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initBoard()
	printBoardBinary(&board)
}
