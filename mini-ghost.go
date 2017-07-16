package main

import (
	"fmt"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Game"
)

func main() {
	path := "./Dictionary/dictionary.txt"
	unfilteredWords, _ := Dictionary.ReadLines(path)
	filteredWords := Dictionary.FilterWords(unfilteredWords)
	game := Game.Game{}
	game.StartGame()
	fmt.Print(len(filteredWords))

}
