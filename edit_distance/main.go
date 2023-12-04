package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str1 string字符串
 * @param str2 string字符串
 * @return int整型
 */
func editDistance(str1 string, str2 string) int {
	// write code here

	m, n := len(str1)+1, len(str2)+1
	var dp [1000][1000]int

	for i := 0; i < m; i++ {
		dp[i][0] = i
	}

	for j := 0; j < n; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	} else if y < x && y < z {
		return y
	} else {
		return z
	}
}

func main() {
	str1 := "lrbb"
	str2 := "mqbhcda"

	n := editDistance(str1, str2)
	fmt.Println(n)
}
