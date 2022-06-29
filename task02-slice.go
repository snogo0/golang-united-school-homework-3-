package homework

func reverse(input []int64) (result []int64) {
	count := len(input)
	result = make([]int64, count, count)
	for _, val := range input {
		count = count - 1
		result[count] = val
	}
	return
}
