package lecture10

/*
Problem:
	Paint Fence with K Colors

	There is a fence ring with n posts, each post can be painted with different kind of color.
	You have to paint all the posts such that no adjacent fence posts have the same color.
	The begin fence and the end fence are adjacent too, so they do not have the same color.
	Return the total number of ways you can paint the fence.
	相邻均不同色， 且第一个和最后一个也不同色。
Objective funtion:
	f(i, k) = f(i - 1, k - 2) + f(i - 2, k - 1)
	分类讨论：
	1. 倒数第二块和第一块颜色不同的时候， 此时f[n]包含的总数为：f[n - 1] * (k - 2) 
		因为最后一块的颜色不能和第一块与倒数第二块相同，最后一块可以选择的颜色有k-2种
  2. 倒数第二块和第一块颜色相同的时候，此时为：f[n - 2] * (k - 1)
		因为最后一块的颜色选择不能和第一块和倒数第二块相同, 最后一块可以选择的颜色有k-1种
	*/

func paintFenceWithKColors(n int, k int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = k
	dp[2] = k*(k-1)
	dp[3] = k*(k-1)*(k-2)
	for i := 4; i <= n; i++ {
			dp[i] = dp[i-1]*(k-2) + dp[i-2]*(k-1)
	}
	return dp[n]
}