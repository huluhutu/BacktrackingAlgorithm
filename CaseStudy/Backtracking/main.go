package main

import (
	"awesomeProject/utils"
	"time"
)

func findMaxForm(strs []string, m int, n int) int {
	var res int
	// 回溯函数
	var backtrack func(pos, zeros, ones, cnt int)
	backtrack = func(pos, zeros, ones, cnt int) {
		// 达到字符串长度上限或者已经超过要求
		if pos == len(strs) || zeros > m || ones > n {
			return
		}
		if cnt > res {
			res = cnt
		}
		// 不选当前字符串
		backtrack(pos+1, zeros, ones, cnt)
		// 选当前字符串
		zerosCnt, onesCnt := utils.CountZerosOnes(strs[pos])
		backtrack(pos+1, zeros+zerosCnt, ones+onesCnt, cnt+1)
	}
	backtrack(0, 0, 0, 0)
	return res
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
