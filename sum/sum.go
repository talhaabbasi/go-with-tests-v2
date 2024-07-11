package sum

func Sum (numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll (numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}