package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetWeightValueMultil(n, maxWeight, maxValue, maxCount int) ([]int, []int, []int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	weights := make([]int, n)
	values := make([]int, n)
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i] = rand.Intn(maxWeight) + 1
		values[i] = rand.Intn(maxValue) + 1
		counts[i] = rand.Intn(maxCount) + 1
	}
	return weights, values, counts
}

func PrintResMultil(weights, values, counts []int, capacity, res int) {
	fmt.Println("Weights:", weights)
	fmt.Println("Values:", values)
	fmt.Println("Counts:", counts)
	fmt.Println("Capacity:", capacity)

	fmt.Println("Max value:", res)
}
