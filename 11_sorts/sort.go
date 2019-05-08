package _1_sorts

//冒泡排序，时间复杂度O(n²)，原地排序，稳定
func BubbleSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		//设置提前退出标识
		flag := false
		for j := 0; j < length-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

				//有元素交换，证明无序，需要继续循环比较
				flag = true
			}
		}

		//一次循环没有元素交换，表示已经有序。循环退出
		if !flag {
			break
		}
	}
}

//插入排序，时间复杂度O(n²)，原地排序，稳定
func InsertionSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		value := a[i] //待插入的元素
		j := i - 1

		for ; j >= 0; j-- {
			//小于当前值，则表示有序，否则交换元素
			if a[j] < value {
				break
			}

			a[j+1] = a[j]
		}

		//最后一次未交换的位置，即为value应该在的位置
		a[j+1] = value
	}
}

//选择排序，时间复杂度O(n²)，原地排序，不稳定
func SelectionSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

//希尔排序，待完善
func ShellSort(a []int) {

}
