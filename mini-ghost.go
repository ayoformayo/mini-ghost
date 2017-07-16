package main

import (
	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Game"
)

func main() {
	// path := "./Dictionary/dictionary.txt"
	// unfilteredWords, _ := Dictionary.ReadLines(path)
	// filteredWords := Dictionary.FilterWords(unfilteredWords)
	dictionary := Dictionary.Dictionary{}
	dictionary.LoadEligibleWords()
	game := Game.Game{ActivePlayer: 0, Dictionary: dictionary}
	game.StartGame()
	// fmt.Print(len(filteredWords))

}
