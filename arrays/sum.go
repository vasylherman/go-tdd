package arrays

func Sum(numbers []int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNums := len(numbersToSum)
	sums := make([]int, lengthOfNums)
	for i, num := range numbersToSum {
		sums[i] = Sum(num)
	}
	return sums
}
