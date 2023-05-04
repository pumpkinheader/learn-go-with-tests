package iteration

func Repeat(text string, cycle int) (repeated string) {
	for	i := 0; i < cycle; i++ {
		repeated += text
	}
	return
}