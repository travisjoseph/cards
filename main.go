package main

func main() {
	cards := newDeckFromfile("myCards")
	cards.shuffle()
	cards.print()
}
