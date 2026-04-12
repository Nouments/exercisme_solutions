package chessboard

import "chessboard"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

type File struct {
	occupied []bool
}

type Chessboard struct {
	chessboard map[string]File
}

func CountInFile(cb Chessboard, file string) int {
	count:=0
	f,ok:=cb.chessboard[file]
	if !ok{
		return count
	}

	for _,occupied := range f.occupied {
		if occupied {
			count++
		}
	}
	return count 
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
