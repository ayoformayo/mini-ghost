package Dictionary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Dictionary struct
type Dictionary struct {
	EligibleWords []string
	TotalWords    []string
}

// LoadEligibleWords gets it in memory
func (dictionary *Dictionary) LoadEligibleWords() {
	lines, _ := dictionary.ReadLines("Dictionary/EligibleDictionary.txt")
	dictionary.EligibleWords = lines
	dictionary.TotalWords = lines
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// FindEligibleFragment finds next possible letter
func (dictionary *Dictionary) FindEligibleFragment(fragment string) string {
	reg := fmt.Sprintf("^%s.+", fragment)
	letterPosition := len(fragment)
	newDictionary := []string{}
	newFragment := ""
	for _, word := range dictionary.EligibleWords {
		match, _ := regexp.MatchString(reg, word)
		if match {
			if newFragment == "" {
				newFragment = string(word[letterPosition])
			}
			newDictionary = append(newDictionary, word)
		}
	}
	dictionary.EligibleWords = newDictionary
	return newFragment
}

// FragmentIsWord sees if this exists
func (dictionary *Dictionary) FragmentIsWord(fragment string) bool {
	return contains(dictionary.EligibleWords, fragment)
}

// ResetDictionary sets eligiblewords to all words again
func (dictionary *Dictionary) ResetDictionary() {
	dictionary.EligibleWords = dictionary.TotalWords
}

// ReadLines opens dictionary file
func (dictionary *Dictionary) ReadLines(path string) ([]string, error) {
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

// FilterWords chooses only winning words
func FilterWords(lines []string) []string {
	var filteredWords []string
	prevValidWord := ""
	for _, v := range lines {
		if len(v) > 3 {
			reg := fmt.Sprintf("^%s", prevValidWord)
			match, _ := regexp.MatchString(reg, v)
			if len(prevValidWord) == 0 || !match {
				prevValidWord = v
				filteredWords = append(filteredWords, v)
				fmt.Println(prevValidWord)
			}
		}
	}
	writeEligibleDictionary(filteredWords)
	return filteredWords
}

// func CheckWord(fragment []string) {
//
// }

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
