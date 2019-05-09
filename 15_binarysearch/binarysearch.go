package _5_binarysearch

//二分查找循环实现
func BinarySearch(arr []int, val int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return -1
	}

	low := 0
	high := arrLen - 1

	for low <= high {
		//mid = (low + high) / 2 = low + (high - low) / 2 = low + ((high - low) >> 1)
		mid := low + ((high - low) >> 1)

		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//二分查找递归实现
func BinarySearchRecursive(arr []int, val int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return -1
	}

	return binarySearchRecursive(arr, val, 0, arrLen-1)
}

func binarySearchRecursive(arr []int, val, low, high int) int {
	mid := low + ((high - low) >> 1)
	if low > high {
		return -1
	}

	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		return binarySearchRecursive(arr, val, low, mid-1)
	} else {
		return binarySearchRecursive(arr, val, mid+1, high)
	}
}

//查找第一个值等于给定值的元素
func BinarySearchFirst(arr []int, val int) int {

}

//查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, val int) int {

}

//查找第一个大于等于给定值的元素
func BinarySearchLastGT(aarr []int, v int) int {
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, v int) int {
}
