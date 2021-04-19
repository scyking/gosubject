package algorithms_test

import (
	"fmt"
	"gosubject/algorithms"
	"math/rand"
	"testing"
)

func arrayRandom(n int) []int {
	array := rand.Perm(n)
	fmt.Printf("生成随机数组：%v \n", array)
	return array
}

func TestBubble(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Bubble(array)
	fmt.Printf("冒泡排序结果：%v \n", array)
}

func TestSelection(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Selection(array)
	fmt.Printf("选择排序结果：%v \n", array)
}

func TestInsertion(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Insertion(array)
	fmt.Printf("插入排序结果：%v \n", array)
}

func TestShell(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Shell(array)
	fmt.Printf("希尔排序结果：%v \n", array)
}

func TestMerge(t *testing.T) {
	array := arrayRandom(10)
	fmt.Printf("归并排序结果：%v \n", algorithms.Merge(array))
}

func TestQuick(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Quick(array)
	fmt.Printf("快速排序结果：%v \n", array)
}

func TestHeap(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Heap(array)
	fmt.Printf("堆排序结果：%v \n", array)
}

func TestCounting(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Counting(array)
	fmt.Printf("插入排序结果：%v \n", array)
}

func TestBucket(t *testing.T) {
	array := arrayRandom(10)
	algorithms.Bucket(array)
	fmt.Printf("桶排序结果：%v \n", array)
}

func TestRadix(t *testing.T) {
	array := arrayRandom(30)
	algorithms.Radix(array)
	fmt.Printf("基数排序结果：%v \n", array)

	arr := []int{1, 2, 44, 56, 23, 99, 4, 101, 33}
	algorithms.Radix(arr)
	fmt.Printf("基数排序结果：%v \n", arr)
}
