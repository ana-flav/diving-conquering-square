package main

import (
	"math"
)


func LineQuery(matrix [][]int, diagonalType string, xStart float64, xEnd float64) (int, int) {
    var resultRow, resultCol int
    var minValue = int(^uint(0) >> 1) 
    	

    n := len(matrix)

    xFirst := int(math.Round(xStart * float64(n)))
	xLast := int(math.Round(xEnd * float64(n)))

    if diagonalType == "x+y=1" { 
        for i := xFirst; i < xLast; i++ {
            if matrix[i][i] < minValue {
                minValue = matrix[i][i]
                resultRow, resultCol = i, i
            }
        }
    } else if diagonalType == "x=y" { 
        for i := xFirst; i < xLast; i++ {
            x := i
            y := n - x
            
            if matrix[x][y] < minValue {
                minValue = matrix[x][y]
                resultRow, resultCol = x, y
            }
        }

    } else if diagonalType == "y=x+1/2" {
		for i := xFirst; i < xLast && i < n; i++ {
			
			col := int(math.Round(float64(i) + 0.5)) 
			if col < n && matrix[i][col] < minValue {
				minValue = matrix[i][col]
				resultRow, resultCol = i, col
			}
		}
    }

    return resultRow, resultCol
}


func Check(matrix [][]int, row, col int) bool {
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


func ProcedureDiagonal(matrix [][]int) (int, int) {

    row1, col1 := LineQuery(matrix, "x=y", 0, 1)
    
  
    if Check(matrix, row1, col1) {
        return row1, col1 

    }
    row2, col2 := LineQuery(matrix, "x+y=1", 0, 0.5)


    if Check(matrix, row2, col2) {
        return row2, col2 
    }


    row3, col3 := LineQuery(matrix, "y=x+1/2", 0, 0.25)


    if Check(matrix, row3, col3) {
        return row3, col3 
    }

    row4, col4 := LineQuery(matrix, "y=-x", 0, 0.25)

	if Check(matrix, row4, col4) {
		return row4, col4
	}

	if row4 == col4 {
		return ProcedureDiamond()
	} else {
		return ProcedureDiamond()
	}

}


// func ProcedureDiamond() (int, int) {
//     return 0, 0
// }
