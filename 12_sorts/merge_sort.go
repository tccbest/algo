package _2_sorts

//递推公式：
//merge_sort(p…r) = merge(merge_sort(p…q), merge_sort(q+1…r))
//终止条件：
//p >= r 不用再继续分解
func MergeSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	mergeSort(arr, 0, length-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmp)
}
