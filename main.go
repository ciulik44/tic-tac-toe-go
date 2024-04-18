package main

import (
	"fmt"
)

func main() {
	board := [3][3]rune{}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = ' '
		}
	}
	printBoard(board)
}

func printBoard(board [3][3]rune) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			fmt.Printf("%c | ", board[row][col])
		}
		fmt.Println()
	}
}
