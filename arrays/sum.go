package arrays


func Sum(numbers []int) (sum int ) {
	for _, number := range numbers {
		sum += number
	}
	return
}


func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range(numbersToSum) {
		sums = append(sums, Sum(numbers))
	}
	return
}


func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}