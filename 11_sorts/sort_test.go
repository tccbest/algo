package _1_sorts

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

func TestBubbleSort(t *testing.T) {
	arr := createRandomArr(10)
	fmt.Println("排序前:", arr)

	BubbleSort(arr)
	fmt.Println("排序后:", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := createRandomArr(10)
	fmt.Println("排序前:", arr)

	InsertionSort(arr)
	fmt.Println("排序后:", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := createRandomArr(10)
	fmt.Println("排序前:", arr)

	SelectionSort(arr)
	fmt.Println("排序后:", arr)
}