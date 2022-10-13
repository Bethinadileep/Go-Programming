// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type deck []string
//Custom Type Declartations

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "sixe of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
