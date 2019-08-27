package main

import (
	"fmt"
)

//插入排序  降序
func insertSort(arr [10]int) [10]int {
	for i := 1;i < len(arr); i++ {
		//j 若采用递增的方式 则需要arr[j+1] 会导致数组越界
		for j := i; j > 0; j-- {
			if arr[j] > arr[j-1] {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}
		}
	}

	return arr
}

//选择排序
/*
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
*/
func selectSort(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i] , arr[j] = arr[j] , arr[i]
			}
		}
	}
	return arr
}

//冒泡排序
/*
它重复地走访过要排序的数列，一次比较两个元素，
如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，
也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
*/
func maopaoSort(arr [10]int) [10]int {
	//len(arr)-1 防止j+1越界
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j] , arr[j+1] = arr[j+1] , arr[j]
			}
		}
	}
	return arr
}

func main(){
	//
	//a := [10]int{12,3,6,78,34,23,43,56,32,46}
	var a [10]int = [10]int{12,3,6,78,34,23,43,56,32,46}
	fmt.Printf("sort before:%v \n",a)
	//ret := insertSort(a)
	//ret := selectSort(a)
	ret := maopaoSort(a)
	fmt.Printf("sort after:%v \n",ret)
}
