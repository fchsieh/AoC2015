package aocutils

import (
	"os"
	"strings"
)

func ReadPuzzle(file_name string) []string {
	raw_data, _ := os.ReadFile(file_name)
	// separate by new line
	data := strings.Split(strings.ReplaceAll(string(raw_data), "\r\n", "\n"), "\n")

	return data
}
