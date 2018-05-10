package main

import (
	"fmt"
	"math/rand"
	"time"
)
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
		fmt.Print("you hit my battleship!")
		printBoard(resultBoard)
	}

	fmt.Print("End of Game")
}

func playGame(playboard [][]string, shipboard [][]string) ([][]string, bool) {
	var x int
	var y int
	
	printBoard(playboard)
	// ask for guess
	fmt.Print("guess x: ")
	fmt.Scan(&x)
	fmt.Print("guess y: ")
	fmt.Scan(&y)
	
	// check if hit 
	guessHit := checkForHit(shipboard, x, y)
	
	if guessHit {
		playboard[x][y] = "X"
		return playboard, true
	}

	playboard[x][y] = "M"
	playGame(playboard, shipboard)

	return playboard, false
}

func initGameBoard(playboard [][]string, n int) [][]string {

	// make n length slice of O's 
	gridRow := make([]string, n)
	for i := range gridRow {
		gridRow[i] = "O"
	}

	// add rows to grid
	for i := range playboard {
		tempGrid := append([]string(nil), gridRow...)
		playboard[i] = tempGrid
	}

	return playboard
}

func initShipBoard(shipboard [][]string, n int, s int) [][]string {

	// generate random loc on grid
	rand.Seed(time.Now().Unix())
	x := rand.Intn(n)
	y := rand.Intn(n)

	// replace loc with "S" and continue replacing row s number of times
	for i := 0; i < s; i++ {
		if y+i > n-1 { break } else {
			shipboard[x][y+i] = "S"
		}
	}
	return shipboard
}

func checkForHit(shipboard [][]string, x int, y int) bool {
	if shipboard[x][y] == "S" {
		return true
	}
	return false
}

func printBoard(board [][]string) {
	fmt.Println("\nGame Board")
	for i := range board {
		fmt.Println(board[i])
	}
}