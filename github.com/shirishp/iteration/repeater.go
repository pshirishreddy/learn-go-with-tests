package iteration
const repeatCount = 6

func Repeat(literal string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += literal
	}
	return repeated
}
