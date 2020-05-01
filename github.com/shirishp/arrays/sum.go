package arrays

func Sum(numbers [5]int) int{
	var sum int
	for _, number := range numbers{
		sum = sum + number
	}
	return sum
}
