package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("V", ""))
}
