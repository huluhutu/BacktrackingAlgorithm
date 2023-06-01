package main

import (
	"awesomeProject/utils"
	"time"
)

func findMaxForm(strs []string, m int, n int) int {
	// 定义状态
	dp := make([][][]int, len(strs)+1)
	for i := 0; i <= len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	// 状态转移
	for i := 1; i <= len(strs); i++ {
		zerosCnt, onesCnt := utils.CountZerosOnes(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= zerosCnt && k >= onesCnt {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-zerosCnt][k-onesCnt]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

// 求最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 生成1万个随机字符串
	m := 10
	n := 10
	num := 10000

	strs := utils.GenerateStrs(m, n, num)
	maxm, maxn := 5, 5
	// 回溯算法求解
	start := time.Now()
	res := findMaxForm(strs, maxm, maxn)
	utils.Print474(start, "backtrackingAlgorithm", res)
}
