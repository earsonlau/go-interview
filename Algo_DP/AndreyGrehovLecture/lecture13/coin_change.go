package lecture13

/*
Problem:
	Coin Change

	Given an unlimited supply of coins of given denominations,
	find the total number of ways to make a change of size n,
	by using exactly t coins.


	Transition function: f(i,t) = f(i-1, t-1) + (i-2, t-1) + f(i-3, t-1) + f(i-5, t-1)
	where 1, 2, 3, 5 are provided coin denomations.

*/


func coinChangeExactlyTCoins(n int, t int, coins []int) int{
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	dp[0][0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= t; j++ {
			if i > 0 && j == 0 {
				dp[i][j] = 0
				continue
			}
			for _, coin := range coins{
				if i - coin >= 0 {
					dp[i][j] += dp[i-coin][j-1]
				}
			}
		}
}
	return dp[n][t]
}