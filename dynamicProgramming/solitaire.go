package main

import (
	"fmt"
	"math"
)

func first1(arr []int, l, r int) int {
	if l > r || r > len(arr) {
		return -1
	}
	if l == r {
		return arr[l]
	}
	p1 := arr[l] + second1(arr, l+1, r)
	p2 := arr[r] + second1(arr, l, r-1)
	// fmt.Println("first ：", int(math.Max(float64(p1), float64(p2))))
	return int(math.Max(float64(p1), float64(p2)))
}

func second1(arr []int, l, r int) int {
	if l == r { //因为是后手，牌以已经被先手取走了，后手没牌可以拿
		return 0
	}
	p1 := first1(arr, l+1, r)
	p2 := first1(arr, l, r-1)
	// fmt.Println("second ：", int(math.Min(float64(p1), float64(p2))))
	return int(math.Min(float64(p1), float64(p2)))
}

func win1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	first := first1(arr, 0, len(arr)-1)
	second := second1(arr, 0, len(arr)-1)
	fmt.Printf("first=%d;second=%d\n", first, second)
	return int(math.Max(float64(first), float64(second)))
}

//缓存版本
func first2(arr []int, l, r int, fmap, smap [][]int) int {
	if fmap[l][r] != -1 {
		return fmap[l][r]
	}
	var ans int
	if l == r {
		ans = arr[l]
	} else {
		p1 := arr[l] + second2(arr, l+1, r, fmap, smap)
		p2 := arr[r] + second2(arr, l, r-1, fmap, smap)
		ans = int(math.Max(float64(p1), float64(p2)))
	}
	fmap[l][r] = ans
	return ans
}

func second2(arr []int, l, r int, fmap, smap [][]int) int {
	if smap[l][r] != -1 {
		return smap[l][r]
	}
	var ans int
	if l == r { //因为是后手，牌以已经被先手取走了，后手没牌可以拿
		ans = 0
	} else {
		p1 := first2(arr, l+1, r, fmap, smap)
		p2 := first2(arr, l, r-1, fmap, smap)
		ans = int(math.Min(float64(p1), float64(p2)))
	}

	smap[l][r] = ans
	return ans
}

func win2(arr []int) int {
	N := len(arr)
	if N == 0 {
		return 0
	}

	fmap := [][]int{}
	smap := [][]int{}
	for i := 0; i < N; i++ {
		fmap = append(fmap, make([]int, N))
		smap = append(smap, make([]int, N))
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmap[i][j] = -1
			smap[i][j] = -1
		}
	}
	first := first2(arr, 0, len(arr)-1, fmap, smap)
	second := second2(arr, 0, len(arr)-1, fmap, smap)
	// fmt.Printf("first=%d;second=%d\n", first, second)
	return int(math.Max(float64(first), float64(second)))
}

/*
牌数组：[1, 100, 5, 20, 100, 20, 5]

f[0][0]=1
f[0][1]=
fmap:
	-->		R
L	[1 0 0 0 0 0 0]							[1 100 6 120 106 140 111]
	[0 100 0 0 0 0 0]						[0 100 100 105 120 140 205]
	[0 0 5 0 0 0 0]							[0 0 5 20 105 105 45]
	[0 0 0 20 0 0 0]						[0 0 0 20 100 40 105]
	[0 0 0 0 100 0 0]						[0 0 0 0 100 100 105]
	[0 0 0 0 0 20 0]						[0 0 0 0 0 20 20]
	[0 0 0 0 0 0 5]							[0 0 0 0 0 0 5]
smap:
	[0 0 0 0 0 0 0]							[0 1 100 6 120 106 140]
	[0 0 0 0 0 0 0]							[0 0 5 20 105 105 45]
	[0 0 0 0 0 0 0]							[0 0 0 5 20 40 105]
	[0 0 0 0 0 0 0]							[0 0 0 0 20 100 40]
	[0 0 0 0 0 0 0]							[0 0 0 0 0 20 20]
	[0 0 0 0 0 0 0]							[0 0 0 0 0 0 5]
	[0 0 0 0 0 0 0]							[0 0 0 0 0 0 0]
*/
//动态规划版抽纸牌
func win3(arr []int) int {
	N := len(arr)
	if N == 0 {
		return 0
	}
	fmap := [][]int{}
	smap := [][]int{}
	for i := 0; i < N; i++ {
		fmap = append(fmap, make([]int, N))
		smap = append(smap, make([]int, N))
	}
	//如果只有一张牌，先手都取走，后手只能选择0
	for i := 0; i < N; i++ {
		fmap[i][i] = arr[i]
	}
	fmt.Println(fmap)
	fmt.Println(smap)
	//先手fmap
	// p1 := arr[l] + second1(arr, l+1, r)
	// p2 := arr[r] + second1(arr, l, r-1)
	// int(math.Max(float64(p1), float64(p2))
	for startCol := 1; startCol < N; startCol++ {
		L := 0
		R := startCol
		for R < N {
			fmap[L][R] = int(math.Max(float64(arr[L]+smap[L+1][R]), float64(arr[R]+smap[L][R-1])))
			smap[L][R] = int(math.Min(float64(fmap[L+1][R]), float64(fmap[L][R-1])))
			L++
			R++
		}
	}
	fmt.Println(fmap)
	fmt.Println(smap)
	return int(math.Max(float64(fmap[0][N-1]), float64(smap[0][N-1])))
}
