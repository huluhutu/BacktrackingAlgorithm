package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetWeightValue(n, maxWeight, maxValue int) ([]int, []int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	weights := make([]int, n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i] = rand.Intn(maxWeight) + 1
		values[i] = rand.Intn(maxValue) + 1
	}
	return weights, values
}

func PrintRes01(weights, values []int, capacity, res int) {
	fmt.Println("Weights:", weights)
	fmt.Println("Values:", values)
	fmt.Println("Capacity:", capacity)

	fmt.Println("Max value:", res)
}
