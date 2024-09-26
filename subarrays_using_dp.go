package main

import "fmt"

func subarray(arr []int, start int) {
  if start >= len(arr) {
    return
  } 
  for i := start; i < len(arr); i++ {
    fmt.Printf("%d, ", arr[i])
    subarray(arr, i + 1)
  }
  fmt.Println()
  return
}
