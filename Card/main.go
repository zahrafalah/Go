package main

import "fmt"

func main() {
	//var Card string = "Ace of spring"
	//Or
	Card := newCard
	fmt.Println(Card)
}

// a return function, when executed, it will return a value of type 'string'
func newCard() string {
	return "Ace of spring"
}
