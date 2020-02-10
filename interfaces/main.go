package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//has type specific for english greeting
	return "Hello!"
}
func (spanishBot) getGreeting() string {
	//has type specific for english greeting
	return "Hola!"
}

func printGreeting(eb englishBot) {
	//has type specific for english greeting
	fmt.Println(eb.getGreeting())
}
func printGreeting(sb spanishBot) {
	//has type specific for english greeting
	fmt.Println(sb.getGreeting())
}
