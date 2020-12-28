package sort

// 快速排序主要是通过基准数对原数组做分割
// 使pivot左边的数字都小于pivot pivot右边的数字都大于pivot
// 然后递归的走遍整个数组

func QuickSort(array []int, left, right int) []int {
	if left < right {
		index := partition(array, left, len(array)-1)
		QuickSort(array, left, index-1)
		QuickSort(array, index+1, right)
	}
	return array
}

func partition(array []int, left, right int) int {
	// 1.定义基准数
	pivot := left
	//
	index := pivot + 1

	for i := index; i <= right; i++ {
		if array[i] < array[pivot] {
			swap(array, i, index)
			index++
		}
	}
	swap(array, pivot, index-1)
	return index - 1
}

func swap(array []int, index1, index2 int) {
	array[index1], array[index2] = array[index2], array[index1]
}
