package main

import (
	"fmt"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Game"
	tree "github.com/ayoformayo/mini-ghost/Tree"
)

func main() {
	path := "./Dictionary/dictionary.txt"
	fmt.Println("Reading Dictionary Lines")
	unfilteredWords, _ := dictionary.ReadLines(path)
	fmt.Println("Lines Loaded")
	fmt.Println("Building Word Tree")
	wordTree := tree.BuildWordTree(unfilteredWords)
	fmt.Println("Word Tree built")
	dictionary := dictionary.Dictionary{WordTree: wordTree}
	// dictionary.LoadEligibleWords()
	game := Game.Game{Dictionary: dictionary}
	game.StartGame()

}
