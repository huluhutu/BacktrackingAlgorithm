package main

import (
	"awesomeProject/utils"
)

func knapsackComplete(C int, w []int, v []int) int {
	n := len(w)
	// 创建一个一维数组 memo，用于保存已经计算过的状态
	memo := make([]int, C+1)
	// 调用记忆化搜索函数，返回最大价值
	return completeKnapsackHelper(C, w, v, n, memo)
}

// 记忆化搜索函数，返回从 0 到 index 可选物品中，容量为 capacity 时的最大价值
func completeKnapsackHelper(C int, w []int, v []int, index int, memo []int) int {
	// 如果已经计算过该状态，直接返回 memo 中保存的结果
	if memo[C] != 0 {
		return memo[C]
	}
	// 如果已经处理完了所有物品或者背包容量为 0，返回 0
	if index == 0 || C == 0 {
		return 0
	}
	// 对于第index个物品，枚举其选择的数量k从0到C/w[index]
	res := 0
	for k := 0; k <= C/w[index-1]; k++ {
		// 递归地求解memo[C-k*w[index]]，并加上k*v[index]
		choose := k*v[index-1] + completeKnapsackHelper(C-k*w[index-1], w, v, index-1, memo)
		res = max(res, choose)
	}
	// 将计算结果保存到 memo 中
	memo[C] = res
	return res
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

	res := knapsackComplete(capacity, weights, values)

	utils.PrintRes01(weights, values, capacity, res)
}
