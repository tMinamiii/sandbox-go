package main

import (
	"fmt"
	"math"
)

func max(lhs, rhs int) int {
	return int(math.Max(float64(lhs), float64(rhs)))
}

func main() {
	var N, M int64
	// input
	fmt.Scanf("%d %d", &N, &M)
	values, weights := make([]int64, N), make([]int64, N)
	for i := int64(0); i < N; i++ {
		fmt.Scanf("%d %d", &weights[i], &values[i])

	}
	// 32bitだとオーバーフローする場合があるので64bit

	dp := make([][]int64, N+1)
	for i := int64(0); i < N+1; i++ {
		dp[i] = make([]int64, M+1)
	}
}
