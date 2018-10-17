package main

import "fmt"

// Create a new type of deck(same as class in OO approach)
// which is a slice of strings
type deck []string

// here we are addind a costume methode to our type:

//(d deck) is considered as receiver
// d: the actual copy of the deck we are working with is
//available in the function as a variable called 'd'. In our case d is cards variable or this.
// deck: Every variable of type 'deck' can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
