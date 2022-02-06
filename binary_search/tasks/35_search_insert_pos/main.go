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
	//arr := []int{1,3,5,6}
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

	fmt.Printf("Search insert position: Elem %d, insert position %d\n", num, searchInsert(arr, num))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
        m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

    return l
}
