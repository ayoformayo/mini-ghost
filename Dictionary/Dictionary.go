package dictionary

import (
	"bufio"
	"fmt"
	"os"

	tree "github.com/ayoformayo/mini-ghost/Tree"
)

// Dictionary struct
type Dictionary struct {
	TotalWords []string
	WordTree   tree.WordTree
}

// LoadEligibleWords gets it in memory
func (thisDictionary *Dictionary) LoadEligibleWords() {
	lines, _ := ReadLines("Dictionary/EligibleDictionary.txt")
	thisDictionary.TotalWords = lines
}

// ReadLines opens dictionary file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeEligibleDictionary(lines []string) {
	f, _ := os.Create("Dictionary/EligibleDictionary.txt")
	// check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, word := range lines {
		withNewline := fmt.Sprintf("%s\n", word)
		writtenWord, _ := w.WriteString(withNewline)
		// w.WriteString("\n")
		fmt.Printf("wrote %d bytes\n", writtenWord)

	}
	w.Flush()
}
