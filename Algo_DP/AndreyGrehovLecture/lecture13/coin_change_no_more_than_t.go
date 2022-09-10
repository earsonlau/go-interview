 package lecture13


 func coinChangeNoMoreTCoins(n int, t int, coins []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	} 
	dp[0][0] = 1
	for i := 0; i <= n; i++ {
		for j:= 0; j <= t; j++{
			if i > 0 && j == 0{
				dp[i][j] = 0
				continue
			}
			if i == 0 && j > 0{
				dp[i][j] = 1
				continue
			}
			for _, coin := range coins{
				if i-coin >= 0 {
					dp[i][j] += dp[i-coin][j-1]
				}
			}

		} 
	}
	return dp[n][t]

 }