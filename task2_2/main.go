package task2_2

import (
	"fmt"
	"sort"
)

func CreateThreeValuesToOneSum(data []int, sum int) [][]int {
	if len(data) < 3 {
		return [][]int{}
	}
	var (
		tempResult = make(map[string][]int)
		numbers    = make(map[int][]int)
		dataLen    = len(data)
		result     [][]int
	)
	sort.Ints(data)
	for i := 0; i < dataLen; i++ {
		numbers[data[i]] = append(numbers[data[i]], i)
	}
	for i := 0; i < dataLen; i++ {
		for j := i + 1; j < dataLen; j++ {
			tempSum := sum - data[i] - data[j]
			if slice, ok := numbers[tempSum]; ok {
			SliceCheck:
				for _, val := range slice {
					if val != i && val != j {
						toSlice := []int{data[i], data[j], tempSum}
						sort.Ints(toSlice)
						tempResult[fmt.Sprintf("%v_%v_%v", toSlice[0], toSlice[1], toSlice[2])] = toSlice
						break SliceCheck
					}
				}
			}
		}
	}
	for _, toSlice := range tempResult {
		result = append(result, toSlice)
	}
	return result
}


