package main

//importing the Standared Go libraries
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new costume type of deck(same as class in OO approach)
// which is a slice of strings. The entire purpose of this file is
//to describe exactly what a deck is & how it behaves.
type deck []string

// here we are adding some costume methods to our type deck:

// (d deck) is considered as receiver
// d: the actual copy of the deck we are working with is
// available in the function as a variable called 'd'. In our case d is cards variable or this.
// deck: Every variable of type 'deck' can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create & return a list of playing cards.
// Essentially an array of strings
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
	// slices have this advance feature to select a range within slice
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//This is going to take the slice of strings, joining all into 1 string separated with commas
	return strings.Join([]string(d), ",")

}

//This func is going to save a list of cards to a hard drive with a filename
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
	// if there was no err we need to turn the bs into the newDeck
	// here we are using strings pkg from Go docs to use split func it
	// takes two args: 1st a string 2d a separator character that we want.
	// string(bs) is going to take the byte slice and turn it into a string.
	s := strings.Split(string(bs), ",")
	// we can return the s as the type of deck
	return deck(s)

}

// // function that randomizes the order of the deck. It needs a receiver of the deck
// // It does not need any return value within it
// func (d deck) shuffle() {
// 	//loop through the cards. we put _ instead of cards cz we will not use it furthur in the code
// 	for i, _ := range d {
// 		//generates a random number btw 0 & len(cards) - 1 by the rand pkg from Go docs
// 		newPosition := rand.Intn(len(d) - 1)
// 		//swapping the current card and the card at cards[randomNumber]
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

// every single time we shuffle our cards with the previous func, It shuffles in the exact
// same order/fashion which is not a perfect way of randomization.
// the rand.Intn that we used is a pseudo random generator that uses the exact same seeds
// every time by default(seed --> generator --> random numbers). In order to fix this problem
// we need to change this seed value every time. Check pkg rand docs/index/type Rand(a source of random nums).
// in this Rand func the source(is also a func) is the seed.
//
func (d deck) shuffle() {
	//lets make our seed value with help of another pkg called time. every single time this func will take the
	//current time as seed for generating newrandom nums. we wil pass this func as an arg for NewSource func
	source := rand.NewSource(time.Now().UnixNano())
	//create new rand obj
	r := rand.New(source)

	for i, _ := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
