package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("myCards")

	fmt.Println(newDeckFromFile("hello"))
}
