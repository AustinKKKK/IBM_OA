package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func minCost(nums []int) int {
	pq:= &IntHeap{}
	for _, v := range nums {
		*pq = append(*pq, v)
	}
	heap.Init(pq)

	cost := 0
	for pq.Len() > 1 {
		a := heap.Pop(pq).(int)
		b := heap.Pop(pq).(int)
		sum := a + b
		cost += sum
		heap.Push(pq, sum)
	}

	return cost
}

func main() {
	nums := []int{5, 3, 8}
	fmt.Println(minCost(nums))
}