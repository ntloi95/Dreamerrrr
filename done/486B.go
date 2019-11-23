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
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
		c[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d", &b[i][j])
			a[i][j] = 1
			c[i][j] = 0
		}
		fmt.Fscanf(reader, "\n")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if b[i][j] == 0 {
				for ii := 0; ii < n; ii++ {
					a[ii][j] = 0
				}

				for jj := 0; jj < m; jj++ {
					a[i][jj] = 0
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 1 {
				for ii := 0; ii < n; ii++ {
					c[ii][j] = 1
				}

				for jj := 0; jj < m; jj++ {
					c[i][jj] = 1
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if b[i][j] != c[i][j] {
				fmt.Print("NO")
				return
			}
		}
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
