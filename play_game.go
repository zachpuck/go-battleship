package main

import "fmt"

func playGame(playboard [][]string, shipboard [][]string) ([][]string, bool) {
	var x int
	var y int
	var result bool

	for !result {
		printBoard(playboard)
		
		// ask for guess
		fmt.Print("guess x: ")
		fmt.Scan(&x)
		fmt.Print("guess y: ")
		fmt.Scan(&y)

		// check if hit 
		hit := checkForHit(shipboard, x, y)

		if hit {
			playboard[x][y] = "X"
			result = true
		} else {
			playboard[x][y] = "M"
		}
	}
	return playboard, result
}

func checkForHit(shipboard [][]string, x int, y int) bool {
	hit := false
	if shipboard[x][y] == "S" {
		hit = true
	}
	return hit
}