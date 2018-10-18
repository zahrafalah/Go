package main

// func main() {
// 	//var Card string = "Ace of spring"
// 	//Or
// 	Card := newCard()
// 	fmt.Println(Card)
// }

// //new main function to Use Slice as an array
// func main() {
// 	//Declare a new variable by Slice
// 	cards := []string{"Ace of Diamonds", newCard()}
// 	//Add to slice
// 	cards = append(cards, "Six of Spade")

// 	//For loop to iterate through the slice
// 	//we need to redeclare or reinitialize card because with for loop
// 	//every single time that we step through the cards, we trough away
// 	//the previous card that had been declared
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}

// }

// //new main function to declare deck type
// func main() {

// 	cards := deck{"Ace of Diamonds", newCard()}

// 	cards = append(cards, "Six of Spade")

// 	cards.print()
// }

// // a return function, when executed, it will return a value of type 'string'
// func newCard() string {
// 	return "Ace of spring"
// }

// //main func to print the newDeck
// func main() {
// 	cards := newDeck()
// 	cards.print()
// }

//new main func to use deal func to split the cards in to two
//we do not need to import the deal function to get access to it since its declared in the same package

// func main() {
// 	cards := newDeck()

// 	//this is a little bit of syntax to capture two values that deal func returns
// 	hand, remaining := deal(cards, 5)
// 	//we can use print function on both of these since they are both type deck
// 	hand.print()
// 	remaining.print()
// }

// //new main func to utilize toString function
// //to get out comma separated list of cards string
// func main() {
// 	cards := newDeck()
// 	fmt.Println(cards.toString())
// }

// //new main func to create a new text file and save our cards inside of it
// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("My_cards")
// }

// //new main func to load a list of cards from the local machine(hard drive file)
// func main() {
// 	//newDeck of cards
// 	cards := newDeckFromFile("My_cards")
// 	cards.print()
// }

//new main fun for shuffle the deck of cards
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
