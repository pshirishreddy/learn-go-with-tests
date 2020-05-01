package slices

func Sum(numbers [] int) int  {
	var sum int
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}
