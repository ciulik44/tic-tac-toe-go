package main

import (
	"fmt"
)

func main() {
	player := 'X'
	gameOver := false
	var scanerRow int
	var scanerCol int

	board := [3][3]rune{}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = ' '
		}
	}

	for !gameOver {
		printBoard(board)
		fmt.Printf("Player %c enter row: ", player)
		fmt.Scan(&scanerRow)
		fmt.Printf("Player %c enter column: ", player)
		fmt.Scan(&scanerCol)

		if player == 'X' {
			player = 'O'
		} else {
			player = 'X'
		}
	}

}

func printBoard(board [3][3]rune) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			fmt.Printf("%c | ", board[row][col])
		}
		fmt.Println()
	}
}
