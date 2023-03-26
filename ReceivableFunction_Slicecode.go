package main

import "fmt"

type deck []string
type book string

func main() {
	cards := deck{"Name1", "Name2", "Name3"}
	cards.print()

	var b book = "Holy Bible"
	b.print1()

}

func (b book) print1() {
	fmt.Println(b)
}

// Receivable Functions
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

