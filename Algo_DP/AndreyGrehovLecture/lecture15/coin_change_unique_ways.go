package lecture15

/* 
Problem:
	Coin change

	Given an unlimited supply of coins of given denominations,
	find the unique number of ways to make a change of size n.

	Denominations:
	coins = [1, 2, 3, 5]

	Transition function:
	
	f[i][t] is the total numebr of unique ways to make a change of size i,
	when the last coin is <= t	

	i >= 1: f[i][1] = f[i-1][1]
	i >= 2: f[i][2] = f[i-1][1] + f[i-2][2]
	i >= 3: f[i][3] = f[i-1][1] + f[i-2][2] + f[i-3][3]
	i >= 5: f[i][5] = f[i-1][1] + f[i-2][2] + f[i-3][3]+ f[i-5][5]
	
	n = 3 coins = 1, 2, 3
	
	No duplicates

	for_, coin := range coins{
		for i:= 1; i <= n; i++ {
			...
		}
	}

	With duplicates

	
	for i:= 1; i <= n; i++ {
		for_, coin := range coins{
			...
		}
	}

	No duplicates(n = 3 coins = 1,2,3)

	coin = [1]
			(1)      (1)      (1)
	3 ------ 2 ------ 1 ------ 0

	coin = [1,2]
			(1)      (1)     (1)
	/------ 2 ------ 1 ------ 0
	3
	\------- 1 ------ 0
			(2)     (1)

	coin = [1,2,3]

			(1)     (1)      (1)
	/------ 2 ------ 1 ------ 0
	|
	|  (3)
	3 ------ 0
	|
	|
	\------- 1 ------ 0
			(2)       (1)

	Answer: 3

	With duplicates

	coins = [1]
			(1)
	1 ------- 0

	
	coins = [1,2]

			(1)       (1)
	/------- 1 ------- 0
	2
	\--------0
			(2)

	coins = [1,2,3]

							(1)      (1)
		 (1)    /------- 1 ------- 0
	 /------ 2
	3         \--------0
	|              (2)
	|  (2)     (1)
	|----- 1 ----- 0
	|
	|  (3)
	|----- 0

	Answer: 4

	*/

/*

	思路一：
	交换lecture12的内外层循环，外层遍历不同面值的硬币，内层遍历不同金额
	从而可以先算出只用某个币能找零的金额，比如只用2元只能找零(2,4,6,8...,2K)的金额。
	从而单币种方案都先算好的情况下，(1,1,2)和(1,2,1)和(2,1,1)就不会再被视为三种方案了
	因为lecture12的(1,1,2)和(1,2,1)和(2,1,1)均被视为变成了(1,1)和2的搭配
	而(1,1)是单币种只能当成一个方案，所以(1,1)和2的搭配也只能是一个方案

*/
	func coinChange(n int, coins []int) int {
		dp := make([]int, n+1)
		dp[0] = 1 
		for _, coin := range coins{
			for i := 1; i <= n; i++ {
				if i - coin >=0 {
					dp += dp[i-coin]
				}
			}
		}
		return dp[n]
	}

/*
	思路二：
	lecture12的内外层循环不变，但是再加一个循环，只保留非递增的情况
	非递增排列要求任意两个相邻数字不能有递增，如9，8，8，7，6，5，,2，2，1
	(1,1,2)和(1,2,1)和(2,1,1)就不会再被视为三种方案了
	因为只保留(2,1,1)一种非递增排列方案
	f[i][t] is the total numebr of unique ways to make a change of size i,
	when the last coin is <= t	
	为什么再加一个循环就能只保留非递增呢，因为这个循环对币种的遍历是有序的，
	比如币种只有{1,2,3,5}，那就是对所有的找零额度的所有币，
	依次计算最后一个用到的硬币面值<=1的方案数，<=2的方案数，<=3的方案数，<=5的方案数
	请注意，此处要保证依次计算，因为前一次计算的结果会影响后面的计算，保证了非递增，故不能并行计算
*/
	func coinChangeUniqueWays(n int, coins []int) int {
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, len(coins))
		}

		for i := range coins{
			dp[0][i] = 1
		}
		// i >= 1: f[i][1] = f[i-1][1]
		// i >= 2: f[i][2] = f[i-2][2] + f[i-1][1] 
		// i >= 3: f[i][3] = f[i-1][1] + f[i-2][2] + f[i-3][3]
		// i >= 5: f[i][5] = f[i-1][1] + f[i-2][2] + f[i-3][3]+ f[i-5][5]
		for i := 0; i <= n; i++{
			for j := range coins{
				for k := 0; k <= j; k++{
					if i-coins[k] < 0 {
						continue
					}
					dp[i][j] += dp[i-coins[k]][k]

				}
			}
				//if i >= 1 {
		//	dp[i][0] = dp[i-1][0]
		//	dp[i][1] = dp[i-1][0]
		//	dp[i][2] = dp[i-1][0]
		//	dp[i][3] = dp[i-1][0]
		//}
		//
		//if i >= 2 {
		//	dp[i][1] += dp[i-2][1]
		//	dp[i][2] += dp[i-2][1]
		//	dp[i][3] += dp[i-2][1]
		//}
		//
		//if i >= 3 {
		//	dp[i][2] += dp[i-3][2]
		//	dp[i][3] += dp[i-3][2]
		//}
		//
		//if i >= 5 {
		//	dp[i][3] += dp[i-5][3]
		//}
		}
		return dp[n][len(coins) - 1]
	}