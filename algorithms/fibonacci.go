package algorithms

// 递归实现斐波那契数列
func Recursion(n int) int {
	if n < 2 {
		return n
	}
	return Recursion(n-1) + Recursion(n-2)
}

// use loop,not recursion
func Loop(n int) int {
	array := [2]int{0, 1}
	if n < 2 {
		return n
	}
	index := 1
	for n >= 2 {
		index = 1 &^ index
		array[index] = array[0] + array[1]
		n--
	}
	return array[index]
}
