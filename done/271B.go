package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	// file, _ := os.Open("input")
	file := os.Stdin
	reader := bufio.NewReader(file)
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	a := make([][]int, n)
	count := make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		count[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d", &a[i][j])
		}
		fmt.Fscanf(reader, "\n")
	}

	// Create Eratos Primes
	size := 200000
	eratos := make([]bool, size)
	for i := 0; i < size; i++ {
		eratos[i] = true
	}

	for i := 2; i < size; i++ {
		if eratos[i] {
			for j := 2; j*i < size; j++ {
				eratos[i*j] = false
			}
		}
	}

	var primes []int
	for i := 2; i < size; i++ {
		if eratos[i] {
			primes = append(primes, i)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nextPrimePos := sort.SearchInts(primes, a[i][j])
			if primes[nextPrimePos] != a[i][j] {
				count[i][j] = primes[nextPrimePos] - a[i][j]
			} else {
				count[i][j] = 0
			}
		}
	}

	result := math.MaxInt32
	for i := 0; i < n; i++ {
		counter := 0
		for j := 0; j < m; j++ {
			counter += count[i][j]
		}

		result = min(result, counter)
	}

	for j := 0; j < m; j++ {
		counter := 0
		for i := 0; i < n; i++ {
			counter += count[i][j]
		}

		result = min(result, counter)
	}

	fmt.Print(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
