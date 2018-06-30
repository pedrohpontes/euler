package multiples

func SumOfMultiplesOf3and5(upperBound int) int {
	total := 0

	for i := 1; i < upperBound; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
