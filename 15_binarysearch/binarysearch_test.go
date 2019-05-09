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
