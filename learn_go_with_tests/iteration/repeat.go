package iteration

const repeatTimes = 5

func Repeat(char string) string {
	var result string
	for i := 0; i < repeatTimes; i++ {
		result += char
	}
	return result
}
