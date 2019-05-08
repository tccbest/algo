package _2_sorts

import (
	"fmt"
	"math/rand"
	"testing"
)

func createRandomArr2(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	arr := createRandomArr2(10)
	fmt.Println("排序前:", arr)

	QuickSort(arr)
	fmt.Println("排序后:", arr)
}
