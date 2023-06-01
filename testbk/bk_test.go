package testbk

import (
	"awesomeProject/utils"
	"testing"
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

func BenchmarkFindMaxForm(b *testing.B) {
	strs := utils.GenerateStrs(5, 5, 1000)
	m, n := 3, 3
	for i := 0; i < b.N; i++ {
		findMaxForm(strs, m, n)
	}
}
