package main

import "fmt"

const (
	english            = "English"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	var helloPrefix string
	switch lang {
	case spanish:
		helloPrefix = spanishHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	return fmt.Sprintf("%s, %s!", helloPrefix, name)
}

func main() {
}
