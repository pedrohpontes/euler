package multiples

// SumOfMultiplesOf3and5 calculates the sum of all integers between 1 and
// upperBound, not including upperBound, that are multiples of either 3
// or 5.
func SumOfMultiplesOf3and5(upperBound int) int {
	total := 0

	for i := 1; i < upperBound; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
