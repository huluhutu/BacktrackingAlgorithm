package main

import "awesomeProject/utils"

func knapsackMulti(weights []int, values []int, counts []int, capacity int) int {
	n := len(weights)
	var items []item
	for i := 0; i < n; i++ {
		for j := 1; j <= counts[i]; j *= 2 {
			counts[i] -= j
			items = append(items, item{weights[i] * j, values[i] * j})
		}
		if counts[i] > 0 {
			items = append(items, item{weights[i] * counts[i], values[i] * counts[i]})
		}
	}

	var maxValue int
	var path []int
	dfs(items, capacity, 0, 0, &maxValue, path)

	return maxValue
}

type item struct {
	weight int
	value  int
}

func dfs(items []item, capacity, start, value int, maxValue *int, path []int) {
	if value > *maxValue {
		*maxValue = value
	}

	if start == len(items) || capacity-items[start].weight < 0 {
		return
	}

	for i := start; i < len(items); i++ {
		path = append(path, i)
		dfs(items, capacity-items[i].weight, i+1, value+items[i].value, maxValue, path)
		path = path[:len(path)-1]
	}
}

func main() {
	n := 10
	capacity := 50
	maxWeight := 10
	maxValue := 20
	maxCount := 5
	weights, values, counts := utils.GetWeightValueMultil(n, maxWeight, maxValue, maxCount)

	res := knapsackMulti(counts, weights, values, capacity)

	utils.PrintResMultil(weights, values, counts, capacity, res)
}
