package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		// Remove punctuation and convert to lowercase
		cleanedWord := strings.Trim(strings.ToLower(word), ".,!?()\"'")
		wordCount[cleanedWord]++
	}

	return wordCount
}

func main() {
	// Read the content of the text file in the same directory
	inputFilePath := "input.txt"

	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string

	for scanner.Scan() {
		content += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	// Count word frequency
	wordFrequency := countWords(content)

	// Display word frequency
	fmt.Println("Word Frequency:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
