package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter text (english characters only): ")
	scanner.Scan()
	text := scanner.Text()

	text = strings.ToLower(text)

	reg := regexp.MustCompile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Println(word, ":", count)
	}
}
