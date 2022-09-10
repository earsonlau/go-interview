package lecture8

/*

Problem:
	Paid Staircase II
	You are climbing a paid staircase. It takes n steps to reach to the top and you have to
	pay p[i] to step on the i-th stair. Each time you can climb 1 or 2 steps.
	Return the cheapest path to the top of the staircase.


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

func climbPaidStairs_best_path_v2(n int,p []int) []int {
	dp := make([] int, n+1)
	from := make([] int,n+1)
	path := []int{}
	dp[0] = 0
	dp[1] = p[1]
	for i := 2; i <= n; i++{
		dp[i] = min(dp[i-1],dp[i-2]) + p[i]	
		if dp[i-1] < dp[i-2]{
			from[i] = i - 1
		}else{
			from[i] = i - 2
		}

}
	for curr := n; curr >= 0; curr = from[curr]{
		path = append(path, curr)
		if curr == 0{
			break
		}
	}	
return reverse(path)
}

// func min(a int, b int) int {
// 	if a > b{
// 		return b
// 	}else{
// 		return a
// 	}
// }

// func reverse(nums []int)[]int{
// 	for i,j :=0,len(nums)-1;i<j; i ,j = i+1, j- 1{
// 		nums[i],nums[j] = nums[j],nums[i]
// 	}
// 	return nums
// }