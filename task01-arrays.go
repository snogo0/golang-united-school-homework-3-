package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	for i := 0; i < 15; i = i + 1 {
		sum = sum + input[i]
	}
	return sum / 15
}
