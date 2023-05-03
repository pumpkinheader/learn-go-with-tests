package hello

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	defaultLanguage := "English"
	greetingPrefixMap := map[string] string {
		"English":"Hello, ",
		"French":"Bonjour, ",
		"Spanish":"Hola, ",
	}
	prefix, isThere := greetingPrefixMap[language]
	if !isThere {
		prefix = greetingPrefixMap[defaultLanguage]
	}
	return
}