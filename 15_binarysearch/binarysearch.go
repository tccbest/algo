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
	arrLen := len(arr)
	if arrLen == 0 {
		return -1
	}

	low := 0
	high := arrLen - 1

	for low <= high {
		mid := low + ((high - low) >> 1)

		if arr[mid] == val {
			if mid == 0 || arr[mid-1] != val {
				return mid
			}

			high = mid - 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, val int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return -1
	}

	low := 0
	high := arrLen - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == val {
			if mid == arrLen-1 || arr[mid+1] != val {
				return mid
			}

			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchLastGT(arr []int, val int) int {
	arrLen := len(arr)
	if arrLen == 0 || arr[arrLen - 1] < val {
		return -1
	}

	low := 0
	high := arrLen - 1

	tag := -1 //记录最后一次比较操作，1表示arr[mid] > val，-1相反，0相等
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == val {
			tag = 0
			if mid != arrLen-1 && arr[mid+1] > val {
				return mid + 1
			}

			low = mid + 1
		} else if arr[mid] > val {
			tag = 1
			high = mid - 1
		} else {
			tag = -1
			low = mid + 1
		}

		if low > high {
			switch tag {
			case 1:
				return mid
			case -1:
				return mid + 1
			case 0:
				return mid + 1
			}
		}
	}

	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, val int) int {
	arrLen := len(arr)
	if arrLen == 0 || arr[arrLen - 1] < val {
		return -1
	}

	low := 0
	high := arrLen - 1

	tag := -1 //记录最后一次比较操作，1表示arr[mid] > val，-1相反，0相等
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == val {
			tag = 0
			if mid != 0 && arr[mid-1] < val {
				return mid - 1
			}

			low = mid + 1
		} else if arr[mid] > val {
			tag = 1
			high = mid - 1
		} else {
			tag = -1
			low = mid + 1
		}

		if low > high {
			switch tag {
			case 1:
				return mid
			case -1:
				return mid
			case 0:
				return mid
			}
		}
	}

	return -1
}
