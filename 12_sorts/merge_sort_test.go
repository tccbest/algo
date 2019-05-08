package _2_sorts

import (
	"fmt"
	"math/rand"
	"testing"
)

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestMergeSort(t *testing.T) {
	arr := createRandomArr(10)
	fmt.Println("排序前:", arr)

	MergeSort(arr)
	fmt.Println("排序后:", arr)
}
