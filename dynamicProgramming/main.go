package main

import "fmt"

func main() {
	cur := 2
	step := 4
	N := 4
	aim := 4
	fmt.Println(Process1(cur, step, aim, N))

	dp := [][]int{}
	for i := 0; i < N+1; i++ {
		dpi := make([]int, 0)
		dpj := make([]int, 0)
		for j := 0; j < step+1; j++ {
			dpj = append(dpj, -1)
		}
		dpi = append(dpi, dpj...)
		dp = append(dp, dpi)

	}
	fmt.Println(Process2(cur, step, aim, N, dp))

	fmt.Println(ways3(cur, step, aim, N))

	arr := []int{1, 100, 5, 20, 100, 20, 5}
	fmt.Println("win score=", win1(arr))
	fmt.Println("win score=", win2(arr))
	fmt.Println("win score=", win3(arr))
}
