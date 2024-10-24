package main

import (
	"fmt"
)


func main() {

	matrix := [][]int{
		{10, 7, 2, 13},
		{6, 21, 9, 8},
		{12, 45, 11, 2},
		{16, 18, 20, 19},
	}

	fmt.Println("%d  %d", len(matrix)-1, len(matrix[0])-1 )
	row, col := ProcedureRowColumn(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)


	fmt.Printf("(%d, %d)", row, col)
}
