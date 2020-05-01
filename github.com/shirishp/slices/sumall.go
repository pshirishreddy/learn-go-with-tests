package slices

func SumAll(numbersToSum ... []int) (sums [] int ) {
	//lengthOfNumbersToSum := len(numbersToSum)
	//sums = make([] int, lengthOfNumbersToSum)
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ... []int) (tailSums [] int){
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tailSums = append(tailSums, Sum(numbers[1:]))
		}
	}
	return
}
