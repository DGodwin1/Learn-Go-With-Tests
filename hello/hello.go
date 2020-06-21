package main

import "fmt"

func main() {
	fmt.Println(Hello("David", ""))
}

var prefixMap = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
	"Rude":    "Leave me alone, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix, ok := prefixMap[language]
	if !ok {
		prefix = prefixMap["English"]
	}

	return prefix + name
}
