package Dictionary

import (
	"bufio"
	"fmt"
	"os"
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
	for _, v := range lines {
		if len(v) > 3 {
			filteredWords = append(filteredWords, v)
		}
	}
	fmt.Print(filteredWords)
	return lines
}
