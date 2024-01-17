package sortingAlgos

func BubbleSort(arr []int) {
	// [1, 5, 4, 2, 3]

	swappedItem := true
	length := len(arr)
	for swappedItem {
		swappedItem = false
		for i := 0; i < length; i++ {
			if i+1 < length && arr[i] > arr[i+1] {
				swappedItem = true
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
		length--
	}
}

func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			curr := i
			prev := i - 1
			for prev >= 0 && arr[prev] > arr[curr] {
				tmp := arr[prev]
				arr[prev] = arr[curr]
				arr[curr] = tmp
				curr--
				prev--
			}
		}
	}
}

func merge(a1, a2 []int) []int {
	var res []int
	l, r := 0, 0
	for l < len(a1) && r < len(a2) {
		if a1[l] <= a2[r] {
			res = append(res, a1[l])
			l++
		} else {
			res = append(res, a2[r])
			r++
		}
	}
	res = append(res, a1[l:]...)
	res = append(res, a2[r:]...)
	return res
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2

	left := MergeSort(arr[:m])
	right := MergeSort(arr[m:])

	return merge(left, right)
}
