package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for k, v := range cb[file] {
		fmt.Println("key: ", k, "value: ", v)
		if v {
			count++
		}

	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank <= 1 && rank >= 8 {
		return 0
	}
	for k, _ := range cb {
		for i, v := range cb[k] {
			if i == rank-1 && v == true {
				fmt.Println(i, v)
				count++
			}
		}

	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for k, _ := range cb {
		for i, v := range cb[k] {
			if v == true {
				fmt.Println(i, v)
				count++
			}
		}

	}
	return count
}

func main() {
	//var board Chessboard

	//fmt.Println(CountInFile(board, "A"))

}
