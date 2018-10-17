package main

import (
	"fmt"
)

// func main() {
// 	//var Card string = "Ace of spring"
// 	//Or
// 	Card := newCard()
// 	fmt.Println(Card)
// }

//new main function to Use Slice as an array
func main() {
	//Declare a new variable by Slice
	cards := []string{"Ace of Diamonds", newCard()}
	//Add to slice
	cards = append(cards, "Six of Spade")

	//For loop to iterate through the slice
	//we need to redeclare or reinitialize card because with for loop
	//every single time that we step through the cards, we trough away
	//the previous card that had been declared
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// a return function, when executed, it will return a value of type 'string'
func newCard() string {
	return "Ace of spring"
}
