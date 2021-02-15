func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
	return
}

func transpose(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := row + 1; col < len(matrix[row]); col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
	return
}

func reverse(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for colStart, colEnd := 0, len(matrix[row])-1; colStart < colEnd; colStart, colEnd = colStart+1, colEnd-1 {
			matrix[row][colStart], matrix[row][colEnd] = matrix[row][colEnd], matrix[row][colStart]
		}
	}
	return
}
