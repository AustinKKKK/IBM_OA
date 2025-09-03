package main

import "fmt"

func calcJobs(m int, gateways []int, jobnums []int ) map[int]int {
	background := 1
	jobSum := make(map[int]int)

	for i, job := range jobnums {
		gate := gateways[i]
		jobSum[gate] += (job + background)
	}

	for gate := 1; gate <= m; gate++ {
		if _, exists := jobSum[gate]; !exists {
			jobSum[gate] = 0
		}
	}

	return jobSum
}

func main() {
	m := 3
	jobnums := []int{5, 4, 3}
	gateways := []int{1, 2, 2}

	result := calcJobs(m, gateways, jobnums)
	fmt.Println(result)
}