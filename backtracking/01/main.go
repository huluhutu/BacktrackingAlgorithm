package main

import "awesomeProject/utils"

var maxVal int
var path []int

func backtrack(curWeight, curValue, idx int, w, v []int, C int) {
	n := len(w)
	if idx == n {
		maxVal = max(maxVal, curValue)
		return
	}
	if curWeight+w[idx] <= C {
		path = append(path, idx)
		backtrack(curWeight+w[idx], curValue+v[idx], idx+1, w, v, C)
		path = path[:len(path)-1]
	}
	backtrack(curWeight, curValue, idx+1, w, v, C)
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
	weights, values := utils.GetWeightValue(n, maxWeight, maxValue)

	backtrack(0, 0, 0, weights, values, capacity)

	utils.PrintRes01(weights, values, capacity, maxVal)
}
