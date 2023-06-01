package main

import (
	"awesomeProject/utils"
	"sort"
)

func knapsack01(weights []int, values []int, capacity int) int {
	n := len(weights)
	items := make([]item, n)
	for i := 0; i < n; i++ {
		items[i] = item{weights[i], values[i], float64(values[i]) / float64(weights[i])}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].valueDensity > items[j].valueDensity
	})

	var path []int
	var maxValue int
	dfs(items, capacity, 0, 0, &maxValue, path)

	return maxValue
}

type item struct {
	weight       int
	value        int
	valueDensity float64
}

func dfs(items []item, capacity, start, value int, maxValue *int, path []int) {
	if value > *maxValue {
		*maxValue = value
	}

	if start == len(items) || capacity-items[start].weight < 0 {
		return
	}

	for i := start; i < len(items); i++ {
		if capacity-items[i].weight < 0 {
			continue
		}
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
	weights, values := utils.GetWeightValue(n, maxWeight, maxValue)

	res := knapsack01(weights, values, capacity)

	utils.PrintRes01(weights, values, capacity, res)
}
