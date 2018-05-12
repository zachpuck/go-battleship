package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func printBoard(board [][]string) {
	fmt.Println("\nGame Board")
	for i := range board {
		fmt.Println(board[i])
	}
}
