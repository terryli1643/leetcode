package main

import "fmt"

func main() {
	dp := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, -1, 0, 0, 0, -1, 0},
		{0, -1, 0, -1, -1, 0, 0, 0},
		{0, 0, 0, 0, 0, -1, 0, 0},
		{0, 0, -1, 0, 0, -1, 0, -1},
		{0, 0, 0, -1, 0, 0, 0, 0},
		{0, -1, 0, 0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	dp[0][0] = 1
	for i := 0; i < 8; i++ {
		dp[i][0] = 1
		dp[0][i] = 1
	}

	for i := 1; i < 8; i++ {
		for j := 1; j < 8; j++ {
			if dp[i][j] != -1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Print(dp[7][7])
}
