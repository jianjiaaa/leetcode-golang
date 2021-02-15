func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]int, len(board))
	column := make([]map[byte]int, len(board))
	box := make([]map[byte]int, len(board))
	for i := 0; i < len(board); i++ {
		row[i] = map[byte]int{}
		column[i] = map[byte]int{}
		box[i] = map[byte]int{}
	}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if string(board[r][c]) != "." {
				_, existInRow := row[r][board[r][c]]
				_, existInColumn := column[c][board[r][c]]
				_, existInBox := box[(r/3)*3+c/3][board[r][c]]
				if existInRow || existInColumn || existInBox {
					return false
				}
				row[r][board[r][c]]++
				column[c][board[r][c]]++
				box[(r/3)*3+c/3][board[r][c]]++
			}
		}
	}
	return true
}
