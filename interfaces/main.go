package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string //Specifies any other type in the program that has the string getgreeting then that type is a bot. Hence English and Spanish Bots are now of type Bot interface
}

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

/*
Commented to write an interface
func printGreeting(eb englishBot) {
	//has type specific for english greeting
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	//has type specific for english greeting
	fmt.Println(sb.getGreeting())
}
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
