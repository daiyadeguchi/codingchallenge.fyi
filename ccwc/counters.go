package main

import (
	"strings"
)

func getLineCount(dat []byte) int {
	content := string(dat)
	lineCount := 0
	for _, line := range content {
		if line == '\n' {
			lineCount++
		}
	}
	return lineCount
}

func getWordCount(dat []byte) int {
	content := strings.Fields(string(dat))
	wordCount := len(content)
	return wordCount
}
