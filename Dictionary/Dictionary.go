package Dictionary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

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
