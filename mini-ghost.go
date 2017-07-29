package main

import (
	"fmt"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Game"
)

func main() {
	path := "./Dictionary/dictionary.txt"
	fmt.Println("Reading Dictionary Lines")
	unfilteredWords, _ := Dictionary.ReadLines(path)
	fmt.Println("Lines Loaded")
	fmt.Println("Building Word Tree")
	wordTree := Dictionary.BuildWordTree(unfilteredWords)
	fmt.Println("Word Tree built")
	dictionary := Dictionary.Dictionary{WordTree: wordTree}
	// dictionary.LoadEligibleWords()
	game := Game.Game{Dictionary: dictionary}
	game.StartGame()

}
