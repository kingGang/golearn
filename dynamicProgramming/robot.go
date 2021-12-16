package main

import "fmt"

//Process1 机器人走step步，来到aim位置，求有几种走法
//例子：N=[1,2,3,4],从2开始走，走4步，要到4那个位置，共8种走法；(Process1(2, 6, 4, 4))=8
//cur 当前位置
//step 一共走几次
//aim 目标位置 aim<=N
//一个有N个位置
func Process1(cur, step, aim, N int) int {

	if step == 0 {
		if cur == aim {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return Process1(2, step-1, aim, N)
	}
	if cur == N {
		return Process1(N-1, step-1, aim, N)
	}
	return Process1(cur-1, step-1, aim, N) + Process1(cur+1, step-1, aim, N)
}

//使用缓存
func Process2(cur, step, aim, N int, dp [][]int) int {

	if dp[cur][step] != -1 {
		return dp[cur][step]
	}
	var ans int
	if step == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = Process2(2, step-1, aim, N, dp)
	} else if cur == N {
		ans = Process2(N-1, step-1, aim, N, dp)
	} else {
		ans = Process2(cur-1, step-1, aim, N, dp) + Process2(cur+1, step-1, aim, N, dp)
	}
	dp[cur][step] = ans
	return ans
}

//动态规划版的 机器人走路
/**
---> step走几步
|	[0 0 0 0 0]
N	[0 0 0 1 0]
	[0 0 1 0 3]
	[0 1 0 2 0]
	[1 0 1 0 2]


	[0 0 0 0 0]
	[0 0 0 1 0]
	[0 0 1 0 4]
	[0 1 0 3 0]
	[1 0 2 0 5]
	[0 1 0 2 0]
*/
func ways3(cur, step, aim, N int) int {
	if N < 2 || cur < 1 || cur > N || aim < 1 || aim > N || step < 1 {
		return -1
	}
	dp := [][]int{}
	for i := 0; i < N+1; i++ {
		dpi := make([]int, step+1)
		dp = append(dp, dpi)

	}
	dp[aim][0] = 1
	for j := 1; j <= step; j++ {
		dp[1][j] = dp[2][j-1]
		for i := 2; i < N; i++ {
			dp[i][j] = dp[i-1][j-1] + dp[i+1][j-1]
		}
		dp[N][j] = dp[N-1][j-1]
	}
	fmt.Println(dp)
	return dp[cur][step]
}
