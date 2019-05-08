package _2_sorts

//递推公式：
//quick_sort(p…r) = quick_sort(p…q-1) + quick_sort(q+1, r)
//终止条件：
//p >= r
func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	quickSort(arr, 0, length-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(arr, start, end)
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	//选择对比元素
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				//i,j交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}
