package main

import "fmt"

func main()  {
	arr:=[]int{2, 5, 8, 9, 10, 4, 3, 16, 1, 7, 8}
	fmt.Println(mergerSort(arr))
}

var count=0

/**
* 归并排序时间复杂度O(n*log n)
 */
func mergerSort(arr []int)[]int  {
	len:=len(arr)
	count++
	fmt.Printf("第%d 次，len= %d\n",count,len)
	if len<=1{
		fmt.Printf("第%d 次，len= %v\n",count,arr)
		return arr
	}
	num:=len/2
	left:=mergerSort(arr[:num])
	right:=mergerSort(arr[num:])
	return merge(left,right)

}

func merge(left,right []int)(result []int)  {
	l,r:=0,0
	for l<len(left)&&r<len(right){
		if left[l]<right[r]{
			result=append(result,left[l])
			l++
		}else {
			result=append(result,right[r])
			r++
		}
	}
	result=append(result,left[l:]...)
	result=append(result,right[r:]...)
	fmt.Println(result)
	return
}