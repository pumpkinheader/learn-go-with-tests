package iteration // Package名はいったん気にしないことにする

func HasPrefix (text string, prefix string) bool {
	return text[0:len(prefix)] == prefix
}