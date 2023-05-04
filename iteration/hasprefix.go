package iteration // Package名はいったん気にしないことにする

func HasPrefix (text string, prefix string) bool {
	split := len(prefix)
	return text[0:split] == prefix
}