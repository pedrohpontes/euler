package multiples

func sumUpTo(upperBound int) int {
	if upperBound < 2 {
		return 0
	}
	return (upperBound * (upperBound - 1)) / 2
}

func sumOfMultiplesOf(divisor, upperBound int) int {
	return divisor * sumUpTo((upperBound-1)/divisor+1)
}

// SumOfMultiplesOf3and5 calculates the sum of all integers between 1 and
// upperBound, not including upperBound, that are multiples of either 3
// or 5.
func SumOfMultiplesOf3and5(upperBound int) int {
	sumOfMultiplesOf3 := sumOfMultiplesOf(3, upperBound)
	sumOfMultiplesOf5 := sumOfMultiplesOf(5, upperBound)
	sumOfMultiplesOf15 := sumOfMultiplesOf(15, upperBound)

	return sumOfMultiplesOf3 + sumOfMultiplesOf5 - sumOfMultiplesOf15
}
