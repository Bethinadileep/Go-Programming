package main

import "fmt"

func main() {
    cards := []string{"Ace of Diamonds", newCard()};
    cards = append (cards, "siz of spades");
    
    for i,card := range cards {
        fmt.Println(i,card);
    }
    
}

func newCard() string{
    return "five Diamonds";
}
