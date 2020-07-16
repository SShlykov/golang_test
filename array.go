package main

import (
	"log"
)

func mult(nums []int) int {
	if len(nums) <= 1 {
		log.Fatalf("cant multiply %d number(s)", len(nums))
	} else if len(nums) > 2 {
		newArr := nums[1:len(nums)]
		return nums[0] * mult(newArr)
	} else {
		return nums[0] * nums[1]
	}
	return 0
}
func cycle(from int, to int, f func(a int)) {
	for i := from; i < to; i++ {
		f(i+1)
	}
}
func bubbleSort(arr []int) {
	for i:=1; i < len(arr); i++ {
		for j:=0; j < len(arr) - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}