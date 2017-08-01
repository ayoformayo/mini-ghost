package main

import (
	"fmt"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Game"
)

func main() {
	path := "./Dictionary/dictionary.txt"
	fmt.Println("Reading Dictionary Lines")
	unfilteredWords, _ := dictionary.ReadLines(path)
	fmt.Println("Lines Loaded")
	fmt.Println("Building Word Tree")
	wordTree := dictionary.BuildWordTree(unfilteredWords)
	fmt.Println("Word Tree built")
	dictionary := dictionary.Dictionary{WordTree: wordTree}
	// dictionary.LoadEligibleWords()
	game := Game.Game{Dictionary: dictionary}
	game.StartGame()

}
