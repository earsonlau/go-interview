package lecture5

import "AndreyGrehovLecture/internal/util"

/*

Problem:
	Climbing Stairs (space optimized)
	You are climbing a stair case. It takes n steps to reach to the top.
	Each time you can either climb 1 or 2 steps.
	In how many distinct ways can you climb to the top?

Framework for Solving DP Problems:
	1. Define the objective function
		f(i) is the number of distinct ways to reach the i-th stair.
	2. Identify base cases
		f(0) = 1
		f(1) = 1
	3. Write down a recurrence relation for the optimized objective function
		f(n) = f(n-1) + f(n-2)
	4. What's the order of execution?
		bottom-up
	5. Where to look for the answer?
		f(n) 

*/

// Time comlexity: O(n)
// Space complexity: O(n)
func climbStairs(n int) int {
	a := 1
	b := 1
	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c 
	}
	util.PrintMemory()
	// fmt.Printf("Allocated Memory123")
	return c
}