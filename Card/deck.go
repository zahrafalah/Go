package main

//importing the Str Go libraries
import (
	"fmt"
	"io/ioutil"
	"os"
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

//function to read from the hard drive file and giving us a new deck from a file
//we dont need a receiver since we already dont have a deck. we need to pass the
//file name as an argument and annotate it as a string.after callin the filename
//the expectation is to receive a deck(the return type would be a type deck).
func newDeckFromFile(filename string) deck {
	//bs short for byteslice & err are variable initialization here.
	bs, err := ioutil.ReadFile(filename)
	//err a value of type 'error'. If nothing went wrong, it will have a value of 'nil: no value'.
	//if err print err:
	if err != nil {
		//Option #1 -log the error and return a call to newDeck()
		//Option #2 -log the error and entirely quit the program
		fmt.Println("Error:", err)
		//OS standard library package
		os.Exit(1)
	}
	// if there was no err we need to  to turn the bs in to the newDeck
	// here we are using strings pkg from Go docs to use split func it
	//takes to args: 1st a string 2d a separate character that we want.
	// string(bs) is going to take the byte slice and turn it into a string.
	s := strings.Split(string(bs), ",")
	// we can return the s as the type of deck
	return deck(s)

}
