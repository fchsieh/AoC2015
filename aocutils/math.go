package aocutils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinArray(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxArray(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
