package main

import (
	"fmt"

	"github.com/ayoformayo/mini-ghost/Dictionary"
)

func main() {
	path := "./Dictionary/dictionary.txt"
	unfilteredWords, _ := Dictionary.ReadLines(path)
	filteredWords := Dictionary.FilterWords(unfilteredWords)
	fmt.Print(len(filteredWords))
}
