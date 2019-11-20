package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("input")
	file := os.Stdin
	reader := bufio.NewReader(file)
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	a := make([][]int, n)
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d", &b[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	hasOneA := false
	hasOneB := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if b[i][j] == 1 {
				hasOneB = true
				fRow := checkRow(b, i, m)
				fCol := checkCol(b, j, n)

				if !fRow && !fCol {
					fmt.Println("NO")
					return
				}

				if fRow && fCol {
					a[i][j] = 1
					hasOneA = true
				} else {
					a[i][j] = 0
				}
			}
		}
	}

	if hasOneA != hasOneB {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%v ", a[i][j])
		}
		fmt.Println("")
	}

	return
}

func checkRow(mx [][]int, r, m int) bool {
	for j := 0; j < m; j++ {
		if mx[r][j] == 0 {
			return false
		}
	}

	return true
}

func checkCol(mx [][]int, c, n int) bool {
	for i := 0; i < n; i++ {
		if mx[i][c] == 0 {
			return false
		}
	}

	return true
}
