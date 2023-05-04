package iteration

func Repeat(character string, cycle int) (repeated string) {
	for	i := 0; i < cycle; i++ {
		repeated += character
	}
	return
}