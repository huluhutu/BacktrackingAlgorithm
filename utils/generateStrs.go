package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机字符串数组
func GenerateStrs(m, n, num int) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	strs := make([]string, num)
	for i := 0; i < num; i++ {
		s := ""
		for j := 0; j < rand.Intn(m+n)+1; j++ {
			if rand.Intn(2) == 0 {
				s += "0"
			} else {
				s += "1"
			}
		}
		strs[i] = s
	}
	return strs
}

func CountZerosOnes(s string) (int, int) {
	var zeros, ones int
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}

// 打印运行时间
func printTime(start time.Time, name string) {
	fmt.Printf("%s: %v\n", name, time.Since(start))
}

func Print474(start time.Time, name string, res int) {
	printTime(start, name)
	fmt.Printf("Max number of selected strings: %d\n", res)
}
