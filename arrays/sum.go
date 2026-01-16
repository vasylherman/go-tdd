package arrays

func Sum(numbers []int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
