package sort

import "fmt"

func main()  {
	arr:=[]int{4,1,3,0,2,7,5,9,8,6,7}
	QuickSort(arr,0,len(arr)-1)
}

func QuickSort(arr []int,left,right int)  {
	if left>=right{
		return
	}
	index:=PartSort(arr,left,right)
	QuickSort(arr,left,index-1)
	QuickSort(arr,index+1,right)
}

//挖坑法
func PartSort(arr []int,left,right int)int {
	key:=arr[right]
	fmt.Println(arr)
	for left<right{
		for left<right&&arr[left]<=key{
			left++
		}
		arr[right]=arr[left]
		for left<right&&arr[right]>key{
			right--
		}
		arr[left]=arr[right]
	}
	arr[right]=key
	fmt.Println(arr)
	return right
}

//交换法，这个方法还有点问题
func PartSort1(arr []int,left,right int)int  {
	fmt.Println("in partsort1")
	key:=arr[right]
	for left<right{
		fmt.Println("index",left,right)
		for left<right&&arr[left]<=key  {
			left++
		}
		for left<right&&arr[right]>key{
			right--
		}
		//fmt.Println(left,right)
		arr[left],arr[right]=arr[right],arr[left]
		fmt.Println(arr)
	}
	arr[left],key=key,arr[left]
	return left
}
