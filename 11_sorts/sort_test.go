package _1_sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 8, 5, 4, 7, 5, 9, 3, 1}
	fmt.Println("排序前:", arr)

	BubbleSort(arr)
	fmt.Println("排序后:", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 8, 5, 4, 7, 5, 9, 3, 1}
	fmt.Println("排序前:", arr)

	InsertionSort(arr)
	fmt.Println("排序后:", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{2, 8, 5, 4, 7, 5, 9, 3, 1}
	fmt.Println("排序前:", arr)

	SelectionSort(arr)
	fmt.Println("排序后:", arr)
}