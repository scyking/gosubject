package algorithms

import (
	"fmt"
	. "sort"
)

func Bubble(arr IntSlice) {
	length := arr.Len()

	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			// 比较相邻元素
			if arr.Less(j+1, j) {
				// tmp := arr[j+1]
				// arr[j+1] = arr[j]
				// arr[j] = tmp
				// swap语法糖
				// arr[j+1],arr[j] = arr[j],arr[j+1]
				arr.Swap(j, j+1)
			}
		}
		fmt.Printf("第%d轮循环，当前数组状态：%v，已排序位置：%v \n", i+1, arr, arr[length-1-i:])
	}
}

func Selection(arr IntSlice) {
	length := arr.Len()

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr.Less(j, minIndex) {
				minIndex = j
			}
		}
		arr.Swap(i, minIndex)
		fmt.Printf("第%d轮循环，当前数组状态：%v，已排序位置：%v \n", i+1, arr, arr[:i+1])
	}
}

func Insertion(arr []int) {
	length := len(arr)

	fmt.Printf("初始状态，当前数组状态：%v，已排序位置：%v \n", arr, arr[:1])

	for i := 0; i < length-1; i++ {
		index := i
		current := arr[i+1]

		for index >= 0 && arr[index] > current {
			// 元素后移
			arr[index+1] = arr[index]
			index--
		}
		// 确定元素插入位置
		arr[index+1] = current
		fmt.Printf("第%d轮循环，当前数组状态：%v，已排序位置：%v \n", i+1, arr, arr[:i+1+1])
	}
}

func Shell(arr []int) {
	length := len(arr)

	for gap := length >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < length; i++ {
			current := arr[i]
			index := i - gap
			for index >= 0 && arr[index] > current {
				// 元素后移
				arr[index+gap] = arr[index]
				index -= gap
			}
			arr[index+gap] = current
		}
		fmt.Printf("增量 gap：%d，插入排序结果：%v \n", gap, arr)
	}
}

func Merge(arr []int) []int {
	length := len(arr)

	if length < 2 {
		return arr
	}
	middle := length >> 1
	return merge(Merge(arr[:middle]), Merge(arr[middle:]))
}

func merge(left, right []int) []int {
	length := len(left) + len(right)
	result := make([]int, length)

	fmt.Printf("归并排序，left：%v，right：%v \n", left, right)

	for index, i, j := 0, 0, 0; index < length; index++ {
		if i >= len(left) {
			result[index] = right[j]
			j++
		} else if j >= len(right) {
			result[index] = left[i]
			i++
		} else if left[i] > right[j] {
			result[index] = right[j]
			j++
		} else {
			result[index] = left[i]
			i++
		}
	}

	fmt.Printf("归并排序结果：%v \n", result)

	return result
}

func Quick(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, start, end int) {
	if start < end {
		index := partition(arr, start, end)
		fmt.Printf("快排分区结果：%v，基准index：%d \n", arr, index)
		quick(arr, start, index-1)
		quick(arr, index+1, end)
	}
}

// 分区操作
func partition(arr IntSlice, start, end int) int {
	// 设定基准值
	pivot := start
	index := pivot + 1

	for i := index; i <= end; i++ {
		if arr.Less(i, pivot) {
			arr.Swap(i, index)
			index++
		}
	}
	index--
	arr.Swap(pivot, index)

	return index
}

func Heap(arr IntSlice) {
	buildMaxHeap(arr)
	fmt.Printf("构建最大堆：%v \n\n", arr)

	length := arr.Len()
	// 将最大堆首位调整到尾部
	for i := length - 1; i > 0; i-- {
		arr.Swap(0, i)
		heapify(arr, 0, i)
	}
}

func buildMaxHeap(arr []int) {
	length := len(arr)
	// 从最后一个非叶子节点开始向上构造最大堆
	for i := length>>1 - 1; i >= 0; i-- {
		heapify(arr, i, length)
	}
}

// 堆调整
func heapify(arr IntSlice, i, length int) {
	left := i<<1 + 1
	right := i<<1 + 2
	largest := i

	// 如果有左子树，且左子树大于父节点，则将最大指针指向左子树
	if left < length && arr.Less(largest, left) {
		largest = left
	}

	// 如果有右子树，且右子树大于父节点，则将最大指针指向右子树
	if right < length && arr.Less(largest, right) {
		largest = right
	}

	// 如果父节点不是最大值，则将父节点与最大值交换，并且递归调整与父节点交换的位置。
	if largest != i {
		fmt.Printf("父节点index：%d，不是最大值，进行堆调整！", i)
		arr.Swap(i, largest)
		heapify(arr, largest, length)
	}
	fmt.Printf("堆调整结果：%v，节点index：%d \n", arr, i)
}

func Counting(arr []int) {
	if len(arr) == 0 {
		return
	}
	min := arr[0]
	max := arr[0]
	// 找出待排序的数组中最大和最小的元素
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	bias := 0 - min
	bucket := make([]int, max-min+1)

	for i := 0; i < len(arr); i++ {
		bucket[arr[i]+bias]++
	}
	fmt.Printf("计数排序，构建桶结果：%v \n", bucket)
	for index, i := 0, 0; index < len(arr); {
		if bucket[i] != 0 {
			arr[index] = i - bias
			bucket[i]--
			index++
		} else {
			i++
		}
	}
}

func Bucket(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 设置桶的默认数量为5
	bucket(arr, 5)
}

func bucket(arr []int, size int) {
	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	counts := (max-min)/size + 1
	buckets := make([][]int, counts)

	for i := 0; i < len(arr); i++ {
		index := (arr[i] - min) / size
		buckets[index] = append(buckets[index], arr[i])
	}
	fmt.Printf("构建桶：%v \n\n", buckets)
	index := 0
	for i := 0; i < len(buckets); i++ {
		// 对每个桶进行排序，使用其他排序方法
		Insertion(buckets[i])
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
}

func Radix(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 取得数组中的最大数，并取得位数
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	digit := 0
	for max != 0 {
		max /= 10
		digit++
	}
	fmt.Printf("最大位数：%d \n", digit)
	radix(arr, digit)
}

func radix(arr []int, digit int) {
	mod := 10
	dev := 1

	radix := make([][]int, 10)

	for i := 0; i < digit; i++ {
		for j := 0; j < len(arr); j++ {
			bucket := int((arr[j] % mod) / dev)
			radix[bucket] = append(radix[bucket], arr[j])
		}
		fmt.Printf("位数：%d；基数数组：%v \n", i+1, radix)
		index := 0
		for j := 0; j < len(radix); j++ {
			for k := 0; k < len(radix[j]); k++ {
				arr[index] = radix[j][k]
				index++
			}
			radix[j] = []int{}
		}
		mod *= 10
		dev *= 10
	}

}
