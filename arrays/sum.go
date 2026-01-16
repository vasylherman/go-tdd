package arrays

func Sum(numbers [3]int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
