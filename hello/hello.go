package hello

# TODO: [ドメインプリミティブへの置き換え](https://qiita.com/takokun778/items/4185cda3f6e6b395ed68)

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