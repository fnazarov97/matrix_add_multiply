package main

import "fmt"

func main() {
	fmt.Println("enter two  number to make matrix for ADD: ")
	add()
	fmt.Println("enter two  number to make matrix for MULTIPLY: ")
	multiply()
}

//FUNCTION FOR ADD TWO MATRIX

func add() {
	var n, m int
	fmt.Print("n = ")
	fmt.Scan(&n)
	fmt.Print("m = ")
	fmt.Scan(&m)
	var matrix1 = makeMatrix(n, m)
	var matrix2 = makeMatrix(n, m)
	var added = make([][]float64, n)
	for i := 0; i < n; i++ {
		added[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			added[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}
	fmt.Println(added)
}

//FUNCTION FOR MULTIPLY TWO MATRIX

func multiply() {
	var n, m int
	fmt.Print("n = ")
	fmt.Scan(&n)
	fmt.Print("m = ")
	fmt.Scan(&m)
	var matrix1 = makeMatrix(n, m)
	var matrix2 = makeMatrix(n, m)
	var multiplied = make([][]float64, n)
	for i := 0; i < n; i++ {
		multiplied[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			multiplied[i][j] = 0
			for k := 0; k < n; k++ {
				multiplied[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	fmt.Println(multiplied)
}

//FUNCTION FOR MAKE MATRIX

func makeMatrix(n, m int) [][]float64 {
	var matrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			fmt.Printf("[%d][%d] : ", i, j)
			fmt.Scanf("%f", &matrix[i][j])
		}
	}
	return matrix
}
