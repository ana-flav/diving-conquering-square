package main

import (
	"math"
)


func LineQueryy(matrix [][]int, diagonalType string, xStart float64, xEnd float64) (int, int) {
	var resultRow, resultCol int
	var minValue = int(^uint(0) >> 1) 
	n := len(matrix)

	
	xFirst := int(math.Round(xStart * float64(n)))
	xLast := int(math.Round(xEnd * float64(n)))


	if diagonalType == "x=y" { 
        for i := xFirst; i < xLast; i++ {
            x := i
            y := n - x
            
            if matrix[x][y] < minValue {
                minValue = matrix[x][y]
                resultRow, resultCol = x, y
            }
        }

	} else if diagonalType == "y=-x+1" {
		for i := xFirst; i < xLast && i < n; i++ {
			x := i
			y := n - 1 - i
			if matrix[x][y] < minValue {
				minValue = matrix[x][y]
				resultRow, resultCol = x, y
			}
		}
	}

	return resultRow, resultCol
}


func CheckNeighbourhood(matrix [][]int, row, col int) bool {
	nRows, nCols := len(matrix), len(matrix[0])
	value := matrix[row][col]

	
	if row > 0 && matrix[row-1][col] < value {
		return false
	}
	if row < nRows-1 && matrix[row+1][col] < value {
		return false
	}
	if col > 0 && matrix[row][col-1] < value {
		return false
	}
	if col < nCols-1 && matrix[row][col+1] < value {
		return false
	}


	return true
}


func ProcedureParaDiagonal(matrix [][]int) (int, int) {
	row1, col1 := LineQueryy(matrix, "y=x", 0.0, 1.0)
	
	if CheckNeighbourhood(matrix, row1, col1) {
		return row1, col1
	}

	row2, col2 := LineQueryy(matrix, "y=-x+1", 0.0, 0.5)

	if CheckNeighbourhood(matrix, row2, col2) {
		return row2, col2
	}

	
}

