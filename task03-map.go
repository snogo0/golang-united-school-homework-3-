package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)
	for key, _ := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	
	result = make([]string, 0)
	for key := range keys {
		result = append(result, input[key])
	}
	return result
}
