
//堆排序
package sort

import "fmt"

func main()  {
	//arr:=[]int{3,5,3,0,8,6,1,5,8,6,2,4,9,4,7, 0, 1, 8, 9, 7, 3, 1, 2, 5, 9, 7, 4, 0, 2, 6 }
	arr:=[]int{20,50,20,40,90,70,10,80,30,60}
	//arr:=[]int{1,2,3,4,5,6,7,8,9}
	//len:=len(arr)
	//min_heap_sort(arr)
	max_heap_sort(arr)
	fmt.Println(arr)
	//min_heap_sort(arr)
	//fmt.Println(arr)

}

func max_heap_sort(arr []int)  {
	len:=len(arr)
	for i:=len/2-1;i>=0;i--{
		max_heapify(arr,i,len-1)
		fmt.Println(arr)
	}
	/**
	* 构建堆
	* 假如有N个节点，那么高度为H=logN，最后一层每个父节点最多只需要下调1次，
	倒数第二层最多只需要下调2次，顶点最多需要下调H次，
	而最后一层父节点共有2^(H-1)个,
	倒数第二层公有2^(H-2),顶点只有1(2^0)个，
	所以总共的时间复杂度为s = 1 * 2^(H-1) + 2 * 2^(H-2) + ... + (H-1) * 2^1 + H * 2^0
	将H代入后s= 2N - 2 - log2(N)，近似的时间复杂度就是O(N)。
	 */
	for i:=len-1;i>0;i-- {
		arr[0],arr[i]=arr[i],arr[0]
		max_heapify(arr,0,i-1)
	}

}
func min_heap_sort(arr []int)  {
	len:=len(arr)
	for i:=len/2-1;i>=0;i--{
		min_heapify(arr,i,len-1)
	}

	for i:=len-1;i>0;i--{
		arr[0],arr[i]=arr[i],arr[0]
		min_heapify(arr,0,i-1)
	}
}

func max_heapify(arr []int,start,end int)  {
	dad:=start
	son:=dad*2+1
	for son<=end{
		if son+1<=end&&arr[son]<arr[son+1]{ //先比较两个子节点大小，选择最大的
			son++     //变成右边节点
		}
		if arr[dad]>arr[son]{   //父节点与右边节点比较
			return
		}else{
			arr[dad],arr[son]=arr[son],arr[dad]
			dad=son
			son=dad*2+1
		}
	}
}



func min_heapify(arr []int,start,end int)  {
	dad:=start
	son:=dad*2+1
	for son<=end{
		if son+1<=end&&arr[son]>=arr[son+1]{
			son++
		}
		if arr[dad]<arr[son]{
			return
		}else{
			arr[dad],arr[son]=arr[son],arr[dad]
			dad=son
			son=dad*2+1
		}

	}
}

