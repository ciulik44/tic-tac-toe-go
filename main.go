package main

import (
	"fmt"
)

func main() {
	player := 'X'
	gameOver := false
	var scanerRow int
	var scanerCol int
	//var result string

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

		if board[scanerRow][scanerCol] == ' ' {
			board[scanerRow][scanerCol] = player
			gameOver = whoWon(board, player)
			if gameOver {
				fmt.Printf("Player %c have won!! ", player)
			} else {
				if player == 'X' {
					player = 'O'
				} else {
					player = 'X'
				}
			}
		} else {
			fmt.Println("Invalid move try again with different bracket")
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

func whoWon(board [3][3]rune, player rune) bool {
	for row := 0; row < len(board); row++ {
		if board[row][0] == player && board[row][1] == player && board[row][2] == player {
			return true
		}
	}
	for col := 0; col < len(board); col++ {
		if board[0][col] == player && board[1][col] == player && board[2][col] == player {
			return true
		}
	}

	return false
}

// func saveScore(result string) {

// }
