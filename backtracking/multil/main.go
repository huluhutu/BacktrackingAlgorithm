package main

import "awesomeProject/utils"

var maxVal int
var path []int

func backtrack(curWeight, curValue, idx int, w, v, M []int, C int) {
	n := len(w)
	if idx == n {
		maxVal = max(maxVal, curValue)
		return
	}
	for i := 0; i <= M[idx]; i++ {
		if curWeight+i*w[idx] <= C {
			for j := 0; j < i; j++ {
				path = append(path, idx)
			}
			backtrack(curWeight+i*w[idx], curValue+i*v[idx], idx+1, w, v, M, C)
			for j := 0; j < i; j++ {
				path = path[:len(path)-1]
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 10
	capacity := 50
	maxWeight := 10
	maxValue := 20
	maxCount := 5
	weights, values, counts := utils.GetWeightValueMultil(n, maxWeight, maxValue, maxCount)

	backtrack(0, 0, 0, weights, values, counts, capacity)

	utils.PrintResMultil(weights, values, counts, capacity, maxVal)
}
