package main

import (
	"fmt"
	"math"
)

type ponto [2]int

func lineQuery(matriz [][]int, linha []ponto) (int, int) {
	menorValor := matriz[linha[0][0]][linha[0][1]]
	menorLinha := linha[0][0]
	menorColuna := linha[0][1]
	for i := 0; i < len(linha); i++ {
		valorAtual := matriz[linha[i][0]][linha[i][1]]
		if valorAtual < menorValor {
			menorValor = valorAtual
			menorLinha = linha[i][0]
			menorColuna = linha[i][1]
		}
	}
	return menorLinha, menorColuna
}

// func getDiagonal(valoresX []int) []ponto {
// 	var linha []ponto
// 	for i := 0; i < len(valoresX); i++ {
// 		linha = append(linha, ponto{valoresX[i], valoresY[i]})
// 	}

// 	return linha
// }

func checkNeighbors(matriz [][]int, x int, y int) bool {
	valor := matriz[x][y]

	if x > 0 {
		if y > 0 {
			if valor > matriz[x][y-1] {
				return false
			}
			if valor > matriz[x-1][y-1] {
				return false
			}
		}

		if valor > matriz[x-1][y] {
			return false
		}
		if valor > matriz[x-1][y+1] {
			return false
		}
	}

	if x < len(matriz)-1 && y > 0 {
		if valor > matriz[x+1][y-1] {
			return false
		}
	}

	if y < len(matriz[0])-1 {
		if valor > matriz[x][y+1] {
			return false
		}
	}

	if x < len(matriz)-1 {
		if valor > matriz[x+1][y] {
			return false
		}
		if valor > matriz[x+1][y+1] {
			return false
		}
	}

	return true
}

func getLista(s int, e int) []int {
	var lista = make([]int, 0)
	if e >= s {
		for i := s; i <= e; i++ {
			lista = append(lista, i)
		}
	} else {
		for i := s; i >= e; i-- {
			lista = append(lista, i)
		}
	}

	return lista
}

func getDiagonal(valoresX, valoresY []int) []ponto {
	var lista []ponto
	// if tipo == "x+y=1" {
	// 	for i := inicioX; i <= fimX; i++ {
	// 		lista = append(lista, ponto{i, i})
	// 	}
	// }

	// if tipo == "x=y" {
	// 	for i := inicioX; i <= fimX; i++ {
	// 		j := n - i
	// 		lista = append(lista, ponto{i, j})
	// 	}
	// }

	aux := valoresX
	if len(valoresY) < len(valoresX) {
		aux = valoresY
	}

	for i := 0; i < len(aux); i++ {
		fmt.Printf("%v ", valoresX[i])
		fmt.Printf("%v \n", valoresY[i])
		lista = append(lista, ponto{valoresX[i], valoresY[i]})
	}
	return lista
}

func ProcedureDiamond(matriz [][]int, rowStart int, rowEnd int, colStart int, colEnd int) (int, int) {
	if rowStart == rowEnd || colStart == colEnd {
		fmt.Println("cai aqui", rowStart, rowEnd, colStart, colEnd)
		if checkNeighbors(matriz, rowStart, colStart) {
			return rowStart, colStart
		}
		return rowStart, colStart + 1
	}

	colRange := colEnd - colStart
	rowRange := rowEnd - rowStart
	fmt.Println("antes: ", rowStart, colStart, colEnd)
	valoresX := getLista(
		int(math.Floor(float64(colRange)*0.25))+colStart,
		int(math.Ceil(float64(colEnd)*0.75)),
	)
	valoresY := getLista(
		int(math.Floor(float64(rowRange)*0.25))+rowStart,
		int(math.Ceil(float64(rowEnd)*0.75)),
	)
	diagonal := getDiagonal(valoresX, valoresY)
	// diagonal := getDiagonal(
	// 	int(math.Floor(float64(colRange)*0.25))+colStart,
	// 	int(math.Ceil(float64(colEnd)*0.75)),
	// 	"x+y=1",
	// 	rowEnd,
	// )
	fmt.Println("diagonal 1: ", diagonal)
	row, col := lineQuery(matriz, diagonal)
	if checkNeighbors(matriz, row, col) {
		return row, col
	}

	valoresX = getLista(
		int(math.Floor(float64(colEnd)*0.25)),
		int(math.Ceil(float64(colEnd)*0.50)),
	)
	valoresY = getLista(
		int(math.Floor(float64(rowRange)*0.50)+float64(rowStart)),
		int(math.Ceil(float64(rowEnd)*0.25)),
	)

	diagonal = getDiagonal(valoresX, valoresY)
	fmt.Println("diagonal 2: ", diagonal)
	fmt.Println("checando", diagonal)
	row, col = lineQuery(matriz, diagonal)
	if checkNeighbors(matriz, row, col) {
		return row, col
	}

	// nRowStart := int(float64(rowEnd) * 0.25)
	nRowStart := int(float64(rowEnd)*0.50) + 1
	nColEnd := int(float64(colEnd) * 0.75)
	nColStart := int(float64(colEnd) * 0.25)

	fmt.Println("depois: ", nRowStart, nColStart, nColEnd)

	return ProcedureDiamond(matriz, nRowStart, rowEnd, nColStart, nColEnd)
}