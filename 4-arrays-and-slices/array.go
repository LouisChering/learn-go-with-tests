package arraysandslices

// Returns a sum of all numbers in the given array
func Sum(numbers []int) int {
	var sum int = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Returns an array of sums for any number for multiple arrays
func SumAll(numberstoSum ...[]int) []int {
	sums := make([]int, len(numberstoSum))
	for i, numbers := range numberstoSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// Returns an array of sums for any number for multiple arrays, excluding the first number of each array
func SumAllTails(numberstoSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberstoSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
