package array_and_slice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	//lengthOfNumbers := len(numbers)
	//sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			// In our case, we are saying "take from 1 to the end" with numbers[1:]
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
