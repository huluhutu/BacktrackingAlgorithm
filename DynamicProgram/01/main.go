package main

import "awesomeProject/utils"

func knapsack(C int, w []int, v []int) int {
	n := len(w)
	// 创建一个二维数组 dp，用于保存已经计算过的状态
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, C+1)
	}
	// 初始化第一行和第一列，表示不选取任何物品或背包容量为0时的情况
	for j := 0; j <= C; j++ {
		dp[0][j] = 0
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	// 逐行填充 dp 数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			// 如果第i个物品重量大于背包容量j，则不能选择该物品
			if w[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 否则，选择第 i 个物品或不选取该物品两种情况中的最大值
				dp[i][j] = max(dp[i-1][j], v[i-1]+dp[i-1][j-w[i-1]])
			}
		}
	}
	// 返回最大价值
	return dp[n][C]
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

	res := knapsack(capacity, weights, values)

	utils.PrintRes01(weights, values, capacity, res)
}
