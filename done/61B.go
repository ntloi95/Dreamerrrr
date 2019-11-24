package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// INPUT
	// file, _ := os.Open("input")
	file := os.Stdin
	reader := bufio.NewReader(file)
	arr := make([]string, 3)
	nStudent := 0
	for i := range arr {
		fmt.Fscanf(reader, "%s\n", &arr[i])
	}
	fmt.Fscanf(reader, "%d\n", &nStudent)
	answers := make([]string, nStudent)
	for i := 0; i < nStudent; i++ {
		fmt.Fscanf(reader, "%s\n", &answers[i])
	}

	// PREPROCESS
	for i := range arr {
		arr[i] = normalize(arr[i])
	}

	// BUILD HASH MAP ALL ACC ANSWER
	positions := []int{0, 1, 2}
	posPermutations := permutate(positions)
	permutations := map[string]bool{}

	for _, p := range posPermutations {
		concat := ""
		for _, v := range p {
			concat += arr[v]
		}
		permutations[strings.ToLower(concat)] = true
	}

	// PROCESS
	for _, answer := range answers {
		realAnswer := normalize(answer)

		if permutations[realAnswer] {
			fmt.Println("ACC")
		} else {
			fmt.Println("WA")
		}
	}

	return
}

func recursiveGenerate(arr []int, curLen, size int, result *[][]int) {
	if curLen == 1 {
		permutation := make([]int, size)
		copy(permutation, arr)
		*result = append(*result, permutation)
		return
	}

	for i := 0; i < curLen; i++ {
		recursiveGenerate(arr, curLen-1, size, result)

		if curLen%2 != 0 {
			arr[0], arr[curLen-1] = arr[curLen-1], arr[0]
		} else {
			arr[i], arr[curLen-1] = arr[curLen-1], arr[i]
		}
	}
}

func permutate(arr []int) [][]int {
	n := len(arr)
	permutations := [][]int{}

	recursiveGenerate(arr, n, n, &permutations)

	return permutations
}

func normalize(s string) string {
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, ";", "")

	return strings.ToLower(s)
}
