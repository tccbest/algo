package _2_sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := createRandomArr(10)
	fmt.Println("排序前:", arr)

	QuickSort(arr)
	fmt.Println("排序后:", arr)
}
