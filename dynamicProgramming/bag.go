package main

import (
	"container/list"
	"fmt"
	"math"
)

func process1(w, v []int, index, bag int) int {
	lenW := len(w)
	if bag < 0 {
		return -1
	}
	if index == lenW {
		return 0
	}
	p1 := process1(w, v, index+1, bag)
	next := process1(w, v, index+1, bag-w[index])
	p2 := 0
	if next != -1 {
		p2 = v[index] + next
	}
	return int(math.Max(float64(p1), float64(p2)))
}

func BagMaxValue() {
	weights := []int{3, 2, 4, 7, 3, 1, 7}
	values := []int{5, 6, 3, 19, 12, 4, 2}
	bag := 15
	fmt.Println(process1(weights, values, 0, bag))
	fmt.Printf("动态规划版本背包问题：%d\n", dp(weights, values, 0, bag))
	// list := list.New()
	dp := []map[string]*list.List{}
	// N := len(weights)
	for i := 0; i < 3+1; i++ {
		dpi := make(map[string]*list.List)
		dp = append(dp, dpi)

	}
	list := Process([]byte("abc"), 0, "", dp)
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/*
------>col 背包容量
      0 1 2 3  4   5  6  7  8  9 10 11 12 13 14 15
  0  [0 4 6 12 16 18 22 22 23 27 31 35 37 41 41 42]
r 1  [0 4 6 12 16 18 22 22 23 25 31 35 37 41 41 41]
o 2  [0 4 4 12 16 16 16 19 23 23 31 35 35 35 35 38]
w 3  [0 4 4 12 16 16 16 19 23 23 31 35 35 35 35 35]
  4  [0 4 4 12 16 16 16 16 16 16 16 18 18 18 18 18]
  5  [0 4 4  4 	4  4  4  4  6  6  6  6  6  6  6  6]
  6  [0 0 0  0 	0  0  0  2  2  2  2  2  2  2  2  2]
  7  [0 0 0  0 	0  0  0  0  0  0  0  0  0  0  0  0]
货物重量
weights := []int{3, 2, 4, 7, 3, 1, 7}
货物价值
values := []int{5, 6, 3, 19, 12, 4, 2}
row 从第几个货物开始
col 背包容量
dp[row][col] 第row个货物背包容量为Col时候背包中的价值
*/
func dp(w, v []int, index, bag int) int {
	N := len(w)
	dp := [][]int{}
	for i := 0; i < N+1; i++ {
		dpi := make([]int, bag+1)
		dp = append(dp, dpi)

	}

	//dp[N][...]=0
	//row 第几个货物
	// col 背包剩余的容量
	for row := N - 1; row >= 0; row-- {
		for col := 0; col <= bag; col++ {
			// 		p1 := process1(w, v, index+1, bag)
			// next := process1(w, v, index+1, bag-w[index])
			// p2 := 0
			// if next != -1 {
			// 	p2 = v[index] + next
			// }
			p1 := dp[row+1][col]
			p2 := 0
			next := col - w[row]
			if next < 0 {
				next = -1
			} else {
				next = dp[row+1][col-w[row]]
			}
			if next != -1 {
				p2 = v[row] + next
			}
			dp[row][col] = int(math.Max(float64(p1), float64(p2)))
		}
	}
	// for _, item := range dp {
	// 	fmt.Println(item)
	// }
	return dp[0][bag]
}

func Process(str []byte, index int, path string, dp []map[string]*list.List) (ans *list.List) {
	if ans == nil {
		ans = list.New()
	}
	// if len(path) >= 2 {

	// 	ans.PushBack(path)
	// 	return
	// }
	if dp[index] != nil {
		if _, ok := dp[index][path]; ok {
			return dp[index][path]
		}
	}
	if index == len(str) {

		ans.PushBack(path)
		mp := make(map[string]*list.List)
		mp[path] = ans
		dp[index] = mp
		return
	}
	ans.PushBackList(Process(str, index+1, path, dp))
	yes := path + string(str[index])

	ans.PushBackList(Process(str, index+1, yes, dp))
	mp := make(map[string]*list.List)
	mp[yes] = ans
	dp[index] = mp
	return
}

// func dpProcess(str []byte, index int, path string) (ans *list.List) {
// 	N := len(str)
// 	dp := []map[string]*list.List{}
// 	for i := 0; i < N+1; i++ {
// 		dpi := make(map[string]*list.List)
// 		dp = append(dp, dpi)

// 	}
// 	for row := N - 1; row >= 0; row-- {
// 		for col := 0; col <= 3; col++ {

// 		}
// 	}
// }
