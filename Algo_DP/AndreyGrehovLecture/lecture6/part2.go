package lecture6

/*

Problem:
	Climbing Stairs (k steps, space optimized, skip red steps)
	You are climbing a stair case. It takes n steps to reach to the top.
	Each time you can either climb 1,...,k steps. You are not allowed to step on red stairs. 
	In how many distinct ways can you climb to the top?

Framework for Solving DP Problems:
	1. Define the objective function
		f(i) is the number of distinct ways to reach the i-th stair.
	2. Identify base cases
		f(0) = 1
		f(1) = 1
	3. Write down a recurrence relation for the optimized objective function
		f(n) = f(n-1) + f(n-2) + ... + f(n-k)
	4. What's the order of execution?
		bottom-up
	5. Where to look for the answer?
		f(n) 

*/

// Time comlexity: O(n)
// Space complexity: O(n)
// 优化原理：因为每次计算dp[i]只需要用到dp[i-1],dp[i-2],..,dp[i-k]共k个值
// 所以不需要dp数组的长度为n+1,dp数组的长度为k就足够过程中的运算了

func climbStairsksteps_so_sr(n int, k int, stairs []bool) int {
	dp := make([] int, k)
	dp[0] = 1
	for i := 1; i <= n; i++{
		for j:= 1; j < k; j++{
			if i-j<0 {
			continue
			}
			if stairs[i-1]{
				dp[i%k]=0
			}else{
				dp[i%k] +=dp[(i-j)%k]
			}
	}
}
	return dp[n%k]
}