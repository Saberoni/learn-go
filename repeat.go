package iteration

const repeatedCount int = 5

func Repeat(character string) (repeated_text string) {
	for i := 0; i < repeatedCount; i++ {
		repeated_text += character
	}
	return
}

func Repeatv2(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
