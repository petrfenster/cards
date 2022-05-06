package cards

func cards() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
