package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of deck which is a slice of string

type deck []string

//Create and return a list of new cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

//new function that belowngs to deck, which prints each element in the deck

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//return an error
func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fileBytes, errors := ioutil.ReadFile(filename)
	if errors != nil {
		// option 1 log the error and return call to newDeck()

		// option 2 : log error and terminate the program
		fmt.Println("Error is ", errors)
		os.Exit(1)
	}
	fileContent := string(fileBytes)
	cards := strings.Split(fileContent, ",")
	return deck(cards)
}
