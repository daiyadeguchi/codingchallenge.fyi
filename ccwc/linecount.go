package main

import "os"

func getLineCount(filepath string) int {
	dat, _ := os.ReadFile(filepath)
	content := string(dat)
	lineCount := 0
	for _, line := range content {
		if line == '\n' {
			lineCount++
		}
	}
	return lineCount
}
