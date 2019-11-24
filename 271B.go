package main

import (
	"bufio"
	"fmt"
	"os"
)

var goodChars string
var k, result int

func main() {
	// file, _ := os.Open("input")
	file := os.Stdin
	reader := bufio.NewReader(file)
	var s string
	fmt.Fscanf(reader, "%s\n%s\n%d", &s, &goodChars, &k)

	trie := &trie{}
	n := len(s)
	for i := 0; i < n; i++ {
		subString := s[i:n]
		trie.insert(subString, 0, 0)
	}

	fmt.Println(result)
	return
}

type trie struct {
	chars [26]*trie
}

func (e *trie) insert(s string, i, v int) {
	if i == len(s) {
		return
	}

	pos := s[i] - 'a'
	if goodChars[pos] != '1' {
		v++
	}

	if e.chars[pos] == nil {
		newNode := &trie{}
		e.chars[pos] = newNode

		if v <= k {
			result++
		}
	}

	e.chars[pos].insert(s, i+1, v)
}
