package main

import (
	"awesomeProject/utils"
)

func knapsack01Memo(weights []int, values []int, capacity int) int {
	n := len(weights)
	// 创建一个二维数组 memo，用于保存已经计算过的状态
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}
	// 调用记忆化搜索函数，返回最大价值
	return knapsack01MemoHelper(weights, values, capacity, n-1, memo)
}

// 记忆化搜索函数，返回从 0 到 index 可选物品中，容量为 capacity 时的最大价值
func knapsack01MemoHelper(weights []int, values []int, capacity, index int, memo [][]int) int {
	// 如果已经计算过该状态，直接返回 memo 中保存的结果
	if memo[index][capacity] != 0 {
		return memo[index][capacity]
	}
	// 如果已经处理完了所有物品或者背包容量为 0，返回 0
	if index < 0 || capacity == 0 {
		return 0
	}
	// 如果当前物品的重量大于背包容量，跳过该物品，递归处理剩下的物品
	if weights[index] > capacity {
		memo[index][capacity] = knapsack01MemoHelper(weights, values, capacity, index-1, memo)
		return memo[index][capacity]
	}
	// 选择当前物品和不选择当前物品两种情况中的最大值
	choose := values[index] + knapsack01MemoHelper(weights, values, capacity-weights[index], index-1, memo)
	notChoose := knapsack01MemoHelper(weights, values, capacity, index-1, memo)
	result := max(choose, notChoose)
	// 将计算结果保存到 memo 中
	memo[index][capacity] = result
	return result
}

// 辅助函数，返回两个数中的最大值
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

	res := knapsack01Memo(weights, values, capacity)

	utils.PrintRes01(weights, values, capacity, res)
}
