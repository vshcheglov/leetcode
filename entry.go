package main

import (
	"./704_binary_search"
	"fmt"
)

func main() {
	result := _704_binary_search.BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Printf("%d", result)
}
