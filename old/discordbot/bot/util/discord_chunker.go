package util

import "strings"

var maxChunkSize = 1900

// Chunk splits an input string into chunks no longer than 1900 characters, using line endings as pivots
func Chunk(input string) []string {
	chunkedLines := ChunkLines(strings.Split(input, "\n"))
	chunks := []string{}
	for _, lines := range chunkedLines {
		chunks = append(chunks, strings.Join(lines, "\n"))
	}
	return chunks
}

// ChunkLines splits an array of strings into groups of strings, with each group no larger than 1900 characters
func ChunkLines(lines []string) [][]string {
	chunks := [][]string{}
	currentChunk := []string{}
	currentChunkLen := 0

	for _, line := range lines {
		if currentChunkLen+len(line) > maxChunkSize {
			chunks = append(chunks, currentChunk)
			currentChunk = []string{}
			currentChunkLen = 0
		}
		currentChunk = append(currentChunk, line)
		currentChunkLen += len(line)
	}
	chunks = append(chunks, currentChunk)
	return chunks
}
