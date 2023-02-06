package homework4

import (
	"fmt"
	"strings"
)

func BubbleSort(slice []int) []int {
	for iter := range slice {
		for i := 0; i < len(slice)-iter; i++ {
			if i == len(slice)-1 {
				break
			}
			target := slice[i]
			if target > slice[i+1] {
				slice[i] = slice[i+1]
				slice[i+1] = target
			}
		}
	}
	return slice
}

func StupidSort(slice []int) []int {
	for range slice {
		for iter, val := range slice {
			if iter == len(slice)-1 {
				break
			}
			if slice[iter] > slice[iter+1] {
				slice[iter] = slice[iter+1]
				slice[iter+1] = val
			}
		}
	}
	return slice
}

func BinarySearch(slice []int, elementToSearch int) int {
	firstIndex, lastIndex := 0, len(slice)-1
	for firstIndex <= lastIndex {
		targetIndex := (firstIndex + lastIndex) / 2
		if slice[targetIndex] == elementToSearch {
			return targetIndex
		}
		if slice[targetIndex] < elementToSearch {
			firstIndex = targetIndex + 1
		} else {
			lastIndex = targetIndex - 1
		}
	}
	return -1
}

func CompareFirstImpl(firstString string, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}
	newMap := make(map[rune]int)
	for _, val := range firstString {
		_, ok := newMap[val]
		if ok {
			i := newMap[val]
			i++
			newMap[val] = i
		} else {
			newMap[val] = 1
		}
	}
	for _, val := range secondString {
		_, ok := newMap[val]
		if ok {
			i := newMap[val]
			i--
			if i == 0 {
				delete(newMap, val)
			} else {
				newMap[val] = i
			}
		} else {
			return false
		}
	}
	return true
}

func CompareSecondImpl(firstString string, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}
	for _, val := range firstString {
		for iter, secondVal := range secondString {
			if val == secondVal {
				secondString = strings.Replace(secondString, string(val), "", iter+1)
				break
			}
		}
	}
	if len(secondString) == 0 {
		return true
	}
	return false
}

func ShowStairsByNumber(n int) {
	stairs := getStairs(n)
	for i := 0; i < len(stairs); i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func getStairs(n int) [][]int {
	globalSlice := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		innerSlice := make([]int, 0)
		for j := 0; j < n; j++ {
			innerSlice = append(innerSlice, j)
		}
		globalSlice = append(globalSlice, innerSlice)
	}
	return globalSlice
}
