package main

import "awesomeProject/utils"

func knapsackMulti(C int, w []int, v []int, m []int) int {
	n := len(w)
	// 创建一个二维数组 memo，用于保存已经计算过的状态
	memo := make([][]int, C+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	// 调用记忆化搜索函数，返回最大价值
	return knapsackMultiHelper(C, w, v, m, n, memo)
}

// 记忆化搜索函数，返回从 0 到 index 可选物品中，容量为 capacity 时的最大价值
func knapsackMultiHelper(C int, w []int, v []int, m []int, index int, memo [][]int) int {
	// 如果已经计算过该状态，直接返回 memo 中保存的结果
	if memo[C][index] != 0 {
		return memo[C][index]
	}
	// 如果已经处理完了所有物品或者背包容量为 0，返回 0
	if index == 0 || C == 0 {
		return 0
	}
	// 对于第 index 个物品，枚举其选择的数量 k 从 0 到 m[index-1] 和 C/w[index-1] 中的较小值
	res := 0
	for k := 0; k <= m[index-1] && k*w[index-1] <= C; k++ {
		// 递归地求解 memo[C-k*w[index-1]][index-1]，并加上 k*v[index-1]
		choose := k*v[index-1] + knapsackMultiHelper(C-k*w[index-1], w, v, m, index-1, memo)
		res = max(res, choose)
	}
	// 将计算结果保存到 memo 中
	memo[C][index] = res
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
	maxCount := 5
	weights, values, counts := utils.GetWeightValueMultil(n, maxWeight, maxValue, maxCount)

	res := knapsackMulti(capacity, weights, values, counts)

	utils.PrintResMultil(weights, values, counts, capacity, res)
}
