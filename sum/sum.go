package sum

func Sum (numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAllTails (numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func SumAllHeads (numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		head := numbers[:1]
		sums = append(sums, Sum(head))
	}

	return sums
}