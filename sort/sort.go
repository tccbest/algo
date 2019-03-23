package sort

func BubbleSort(a []int) {
    length := len(a)
    if length <= 1 {
        return
    }

    for i := 0; i < length; i++ {
        flag := false
        for j := 0; j < length-i-1; j++ {
            if a[i] < a[j+1] {
                a[i], a[j+1] = a[j+1], a[i]
                flag = true
            }

            if !flag {
                break
            }
        }
    }
}
