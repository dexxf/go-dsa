package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 5, 7, 10}
	target := 10

	for i := range len(arr) {
		if arr[i] == target {
			fmt.Println("Found target at index:", i)
			return
		}
	}

	fmt.Println("Target not found")
}
