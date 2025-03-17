package main

import "fmt"

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	japaneseHelloPrefix = "Konichiwa, "

	spanish  = "Spanish"
	japanese = "Japanese"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getHelloPrefix(language) + name
}

func getHelloPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case japanese:
		return japaneseHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
