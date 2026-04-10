package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok {
		return 0
	}

	count := 0
	for _, v := range f {
		if v {
			count++
		}
	}
	return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	index := rank - 1

	count := 0

	for _, file := range cb {
		if index < len(file) && file[index] {
			count++
		}
	}

	return count
}

func CountAll(cb Chessboard) int {
	count := 0

	for _, file := range cb {
		count += len(file)
	}

	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0

	for _, file := range cb {
		for _, v := range file {
			if v {
				count++
			}
		}
	}

	return count
}