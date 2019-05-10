package _5_binarysearch

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func createRandArr(n int) []int {
	r1 := rand.New(rand.NewSource(time.Now().Unix()))
	arr := make([]int, 0, 1000)
	i := r1.Intn(100)
	for ; i < n; {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		arr = append(arr, i)
		i += r.Intn(100)
	}

	return arr
}

func TestBinarySearch(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	arr := createRandArr(10000)

	fmt.Println("数组：", arr)

	val := r.Intn(10000)
	fmt.Println("查找元素：", val)

	resp := BinarySearch(arr, val)
	fmt.Println(resp)
}

func TestBinarySearchRecursive(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	arr := createRandArr(10000)

	fmt.Println("数组：", arr)

	val := r.Intn(10000)
	fmt.Println("查找元素：", val)

	resp := BinarySearchRecursive(arr, val)
	fmt.Println(resp)
}

func TestBinarySearchFirst(t *testing.T) {
	arr := []int{2, 3, 10, 10, 10, 23, 56, 30}
	val := 10

	fmt.Println("数组：", arr)
	fmt.Println("查找元素：", val)

	resp := BinarySearchFirst(arr, val)
	fmt.Println(resp)
}

func TestBinarySearchLast(t *testing.T) {
	arr := []int{2, 3, 10, 10, 10, 23, 56, 30}
	val := 10

	fmt.Println("数组：", arr)
	fmt.Println("查找元素：", val)

	resp := BinarySearchLast(arr, val)
	fmt.Println(resp)
}

func TestBinarySearchLastGT(t *testing.T) {
	arr := []int{2, 3, 10, 10, 10, 23, 56, 80}
	val := 90

	fmt.Println("数组：", arr)
	fmt.Println("查找元素：", val)

	resp := BinarySearchLastGT(arr, val)
	fmt.Println(resp)
}

func TestBinarySearchLastLT(t *testing.T) {
	arr := []int{2, 3, 10, 10, 10, 23, 56, 80}
	val := 22

	fmt.Println("数组：", arr)
	fmt.Println("查找元素：", val)

	resp := BinarySearchLastLT(arr, val)
	fmt.Println(resp)
}
