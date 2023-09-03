package iteration

// const repeatCount = 5

// Make a repeated String from given input string
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {

		repeated += character
	}
	return repeated
}
