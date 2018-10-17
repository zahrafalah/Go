package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck(same as class in OO approach)
// which is a slice of strings
type deck []string

// here we are adding a costume method to our type:

//(d deck) is considered as receiver
// d: the actual copy of the deck we are working with is
//available in the function as a variable called 'd'. In our case d is cards variable or this.
// deck: Every variable of type 'deck' can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create & return a list of playing cards.
//Essentially an array of strings
func newDeck() deck {
	cards := deck{}
	// slice of type string
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//two nested for loops to iterate (i,j) through the suites & cards and create new cards
	//instd of i and j  we put _ just to get rid of the err of not using them furthur in the code
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//func without receiver & deck is as an argument
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//this is going to take the slice of strings, joining all into 1 string separated with commas
	return strings.Join([]string(d), ",")

}

//this func is going to save a list of cards to a hard drive with a filename
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
