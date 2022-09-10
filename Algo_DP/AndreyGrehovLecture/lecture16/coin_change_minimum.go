package lecture16

import (
	"math"
	"AndreyGrehovLecture/internal/util"

)

/*

Problem:
	Given an unlimited supply of coins of given denominations,
	what is the minimum number of coins needed to make a change of size n?

	coins = {1, 3, 5}

*/

func coinChangeMinimum(n int, coins []int) int{
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++{
		dp[i] = math.MaxInt64
		for j := range coins{
			// j 为coins数组的下标
			if i-coins[j] < 0 || 1 + dp[i-coins[j]] == math.MinInt64{
				continue
			}
			// 用第j个硬币的方案数为1+dp[i-coins[j]]
			// 不用第j个硬币的方案数为dp[i](如果dp[i]还没计算过，默认是math.MaxInt64)
			dp[i] = util.Min(dp[i], 1 + dp[i-coins[j]])
		}
		//if i >= 1 {
		//	dp[i] = util.Min(dp[i], 1+dp[i-1])
		//}
		//if i >= 3 {
		//	dp[i] = util.Min(dp[i], 1+dp[i-1], 1+dp[i-3])
		//}
		//if i >= 5 {
		//	dp[i] = util.Min(dp[i], 1+dp[i-1], 1+dp[i-3], 1+dp[i-5])
		//}

		if dp[n] == math.MaxInt64{
			return -1
		}

	}
	return dp[n]

}