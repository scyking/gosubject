package algorithms_test

import (
	"fmt"
	"gosubject/algorithms"
	"math/rand"
	"testing"
)

func arrayRandom() []int {
	array := rand.Perm(10)
	fmt.Printf("生成随机数组：%v \n", array)
	return array
}

func TestBubble(t *testing.T) {
	array := arrayRandom()
	algorithms.Bubble(array)
	fmt.Printf("冒泡排序结果：%v \n", array)
}

func TestSelection(t *testing.T) {
	array := arrayRandom()
	algorithms.Selection(array)
	fmt.Printf("选择排序结果：%v \n", array)
}

func TestInsertion(t *testing.T) {
	array := arrayRandom()
	algorithms.Insertion(array)
	fmt.Printf("插入排序结果：%v \n", array)
}

func TestShell(t *testing.T) {
	array := arrayRandom()
	algorithms.Shell(array)
	fmt.Printf("希尔排序结果：%v \n", array)
}

func TestMerge(t *testing.T) {
	array := arrayRandom()
	fmt.Printf("归并排序结果：%v \n", algorithms.Merge(array))
}

func TestQuick(t *testing.T) {
	array := arrayRandom()
	algorithms.Quick(array)
	fmt.Printf("快速排序结果：%v \n", array)
}

func TestHeap(t *testing.T) {
	array := arrayRandom()
	algorithms.Heap(array)
	fmt.Printf("堆排序结果：%v \n", array)
}
