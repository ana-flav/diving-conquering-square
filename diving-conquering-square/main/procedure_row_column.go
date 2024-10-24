package main

import (
	"fmt"
	"math"
)


func ColumnQuery(matrix [][]int, col int, rowStart int, rowEnd int) (int, int) {
	minRow := rowStart
	minValue := matrix[rowStart][col]
	for i := rowStart; i <= rowEnd; i++ {
		if matrix[i][col] < minValue {
			minValue = matrix[i][col]
			minRow = i
		}
	}
	return minRow, minValue
}

func RowQuery(matrix [][]int, row int, colStart int, colEnd int) (int, int) {
	minCol := colStart
	minValue := matrix[row][colStart]
	for j := colStart; j <= colEnd; j++ {
		if matrix[row][j] < minValue {
			minValue = matrix[row][j]
			minCol = j
		}
	}
	return minCol, minValue
}

func CheckLocalMinimum(matrix [][]int, row int, col int) bool {
	value := matrix[row][col]
	fmt.Println(value)
	nRows, nCols := len(matrix), len(matrix[0])

	if row > 0 && matrix[row-1][col] < value {
		// verifica se o valor na linha de cima e maior que o valor atual
		fmt.Println("1")
		return false
	}
	if row < nRows-1 && matrix[row+1][col] < value {
		// verifica se o valor na linha de baixo e maior que o valor atual
		fmt.Println("2")

		return false
	}
	if col > 0 && matrix[row][col-1] < value {
		// verifica se o valor na coluna da esquerda e maior que o valor atual
		return false
	}
	if col < nCols-1 && matrix[row][col+1] < value {
		fmt.Println("4")
		return false
	}
	return true
}


func ProcedureRowColumn(matrix [][]int, rowStart int, rowEnd int, colStart int, colEnd int) (int, int) {

	if rowStart == rowEnd && colStart == colEnd {
		return rowStart, colStart
	}

	
	midCol := (colStart + colEnd) / 2
	rowMin, colMinValue := ColumnQuery(matrix, midCol, rowStart, rowEnd)
	fmt.Println(rowMin, colMinValue)

	leftValue := math.MaxInt64
	rightValue := math.MaxInt64
	if midCol > colStart {
		leftValue = matrix[rowMin][midCol-1]
	}
	if midCol < colEnd {
		rightValue = matrix[rowMin][midCol+1]
	}

	fmt.Println("testetestet")
	fmt.Println(rowMin, midCol)
	if CheckLocalMinimum(matrix, rowMin, midCol) {
		return rowMin, midCol
	}

	midRow := (rowStart + rowEnd) / 2
	fmt.Println(midRow)
	colMin, rowMinValue := RowQuery(matrix, midRow, colStart, colEnd)
	fmt.Println("snsnsnnsns")
	fmt.Println(colMin, rowMinValue)

	upperValue := math.MaxInt64
	
	if midRow > rowStart {
		upperValue = matrix[midRow-1][colMin]
	}
	


	if CheckLocalMinimum(matrix, midRow, colMin) {
		return midRow, colMin
	}


	
	if midCol > colStart && colMinValue > leftValue {
		return ProcedureRowColumn(matrix, rowStart, rowEnd, colStart, midCol-1)
	} else if midCol < colEnd && colMinValue > rightValue {
		return ProcedureRowColumn(matrix, rowStart, rowEnd, midCol+1, colEnd)
	} else if midRow > rowStart && rowMinValue > upperValue {
		return ProcedureRowColumn(matrix, rowStart, midRow-1, colStart, colEnd)
	} else {
		return ProcedureRowColumn(matrix, midRow+1, rowEnd, colStart, colEnd)
	}
}
