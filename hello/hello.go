package main

import "fmt"

func main(){
	fmt.Println(Hello("David", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const rudePrefix = "Leave me alone, "

const spanish = "Spanish"
const french = "French"
const rude = "Rude"

func Hello(name, language string) string{
	if name == ""{
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {

	case spanish :
		prefix = spanishHelloPrefix

	case french :
		prefix = frenchHelloPrefix

	case rude :
		prefix = rudePrefix
	}

	return prefix + name
}