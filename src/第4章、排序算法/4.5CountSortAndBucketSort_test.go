package test

import (
	"fmt"
	"testing"
)

//--------------计数排序--------------
func TestCountSort(t *testing.T) {
	array := []uint32{3, 2, 2, 3, 4, 62, 6, 5, 3, 6, 2, 45, 23, 65}
	fmt.Println(countSortV1(array))
}

func countSortV1(list []uint32) []uint32 {
	//1.寻找最大值
	var max uint32
	var min uint32
	length := len(list)
	for i := 0; i < length; i++ {
		if list[i] > max {
			max = list[i]
		} else if list[i] < min {
			min = list[i]
		}
	}
	diff := max - min

	//2.设置线型数组长度
	countArray := make([]uint32, diff+1)

	//3.遍历统计数组，填充统计数组
	for _, value := range list {
		countArray[value-min]++
	}

	//4.遍历统计数组，输出结果
	sortedArray := make([]uint32, 0, length)
	for i, v := range countArray {
		for j := uint32(0); j < v; j++ {
			sortedArray = append(sortedArray, uint32(i)+min)
		}
	}
	return sortedArray
}

//------------------------------------------
