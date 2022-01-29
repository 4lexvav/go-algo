package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1, 2, 4, 5, 8, 9, 11, 12, 13, 15, 18, 19, 23, 32, 35, 38, 54, 89, 233}
	fmt.Printf("Given an array: %v\n", arr)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose the digit: ")

	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(strings.Trim(str, " \n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Looping: Elem %d, found at index %d\n", num, searchLoop(arr, num))
	fmt.Printf("Recursive: Elem %d, found at index %d\n", num, searchRecursive(arr, num))
}

func searchLoop(arr []int, n int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (r + l) / 2
		if arr[m] == n { return m }
		if arr[m] > n { r = m - 1 }
		if arr[m] < n { l = m + 1 }
	}
	return -1
}

func searchRecursive(arr []int, n int) int {
	if len(arr) == 0 { return -1 }
	m := (len(arr) - 1) / 2
	if arr[m] == n { return m }

	if arr[m] > n { return searchRecursive(arr[0:m], n) }
	pos := searchRecursive(arr[m + 1:], n)
	if pos < 0 { return pos }
	return m + 1 + pos
}
