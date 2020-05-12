package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		s := make([]int, n+1)
		dp := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Scan(&s[i])
			dp[i] = 1
		}

		max := 0
		for i := 1; i <= n; i++ {
			for j := i * 2; j <= n; j += i {
				if s[j] > s[i] && dp[j] < dp[i]+1 {
					dp[j] = dp[i] + 1
				}
			}
			if max < dp[i] {
				max = dp[i]
			}
		}
		fmt.Println(dp)
		fmt.Println(max)
	}
}
