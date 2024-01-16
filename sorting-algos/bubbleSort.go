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
