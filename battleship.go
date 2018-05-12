package main

import "fmt"

var size int
var shipSize int

func main() {
	fmt.Println("Welcome to Battleship!")
	
	fmt.Print("Enter the size of the board: \n")
	fmt.Scan(&size)
	
	shipSize = 3

	// make playboard
	playBoard := make([][]string, size)
	playBoard = initGameBoard(playBoard, size)

	// make shipboard
	shipBoard := make([][]string, size)
	shipBoard = initGameBoard(shipBoard, size)
	shipBoard = initShipBoard(shipBoard, size, shipSize)

	resultBoard, result := playGame(playBoard, shipBoard)

	if result {
		fmt.Print("\nYou hit my Battleship!")
		printBoard(resultBoard)
	}
	fmt.Print("End of Game")
}
