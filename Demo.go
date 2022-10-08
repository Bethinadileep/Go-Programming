//Introduction
package main

import "fmt"

func main() {
    //var card string = "Ace of Spades"
    //Initializing in another way
    card := "Ace of Spades"
    card = "Five of Diamonds"
    fmt.Println(card)
}

//functions
package main

import "fmt"

func main() {
    card := MyCard()

    fmt.Println(card)
}

func MyCard() string {
    return "Five Diamonds"
}

