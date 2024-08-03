package main

import "fmt"

func subarray(arr []int, start int, end int) {
	if end == len(arr) {
		return
	} else if start > end {
		subarray(arr, 0, end+1)
	} else {
		fmt.Printf("[")
		for i := start; i < end; i++ {
			fmt.Printf("%d, ", arr[i])
		}
		fmt.Printf("%d]\n", arr[end])
		subarray(arr, start+1, end)
	}
	return
}

func main() {
  arr := []int{1, 2, 3, 4}
  subarray(arr, 0, 0);
}
