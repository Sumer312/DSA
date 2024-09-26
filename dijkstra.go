package main

import (
	"fmt"
	"math"
)

type node struct {
	v int
	w int
}

type MinHeap struct {
	data []node
}

func (h *MinHeap) Insert(value node) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) Delete() (node, bool) {
	if len(h.data) == 0 {
		return node{-1, -1}, true
	}
	value := h.data[0]
	if len(h.data) == 1 {
		h.data = []node{}
		return value, false
	}
	h.data[0] = h.data[len(h.data)-1]
	newSlice := h.data[:len(h.data)-1]
	h.data = newSlice
	h.heapifyDown(0)
	return value, false
}

func (h *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := h.parent(index)
	parentValue := h.data[parentIndex]
	currentValue := h.data[index]
	if currentValue.w < parentValue.w {
		h.data[index] = parentValue
		h.data[parentIndex] = currentValue
		h.heapifyUp(parentIndex)
	}
	return
}

func (h *MinHeap) heapifyDown(index int) {
	if index == len(h.data)-1 {
		return
	}
	lcIndex := h.leftChild(index)
	rcIndex := h.rightChild(index)
	if lcIndex >= len(h.data) || rcIndex >= len(h.data) {
		return
	}
	lcValue := h.data[lcIndex]
	rcValue := h.data[rcIndex]
	currentValue := h.data[index]
	if lcValue.w > rcValue.w && currentValue.w > rcValue.w {
		h.data[lcIndex] = currentValue
		h.data[index] = lcValue
		h.heapifyDown(lcIndex)
	} else {
		h.data[rcIndex] = currentValue
		h.data[index] = rcValue
		h.heapifyDown(rcIndex)
	}
	return
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}

func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

func dijkstra() {
	var length int = 5
	var source int = 1
	var dists []int = make([]int, length+1)
	list := make(map[int][]node)
	for i := 1; i <= length; i++ {
		dists[i] = math.MaxUint32
		list[i] = []node{}
	}

	list[1] = append(list[1], node{2, 0})
	list[1] = append(list[1], node{3, 0})
	list[1] = append(list[1], node{4, 0})

	list[3] = append(list[3], node{4, 0})

	list[4] = append(list[4], node{5, 0})

	h := MinHeap{[]node{}}
	dists[source] = 0

	h.Insert(node{source, 0})
	for len(h.data) > 0 {
		temp, flag := h.Delete()
		if flag {
			return
		}
		fmt.Printf("Popped: %d -> %v\n", source, temp)
		curr := temp.v
		weight := temp.w
		for i := 0; i < len(list[curr]); i++ {
			neighbor := list[curr][i].v
			w := weight - 1
			next := node{neighbor, w}
			if w < dists[neighbor] {
				dists[neighbor] = w
				fmt.Printf("Inserted: %d -> %v\n", source, next)
				h.Insert(next)
			}
		}
	}

	for i := 1; i < len(dists); i++ {
		fmt.Printf("%d -> %d = %d\n", source, i, dists[i])
	}

}
func main() {
	dijkstra()
}
