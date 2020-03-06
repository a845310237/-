package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	i := 1
	j := 2
	sum := i + j
	var arr [][]int
	//arr := make([][]int, target/2)
	for i <= target/2 {
		if sum == target {
			var temp []int
			for k := i ;k <= j;k++ {
				temp = append(temp, k)
			}
			arr = append(arr, temp)
			sum -= i
			i++
		}
		if sum < target {
			j++
			sum += j
		}
		if sum > target {
			sum -= i
			i++
		}
	}
	return arr
}

func main() {
	var arr [][]int
	arr = findContinuousSequence(9)
	for _, v := range arr {
		fmt.Println(v)
	}
}