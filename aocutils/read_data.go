package aocutils

import (
	"os"
	"strings"
)

func ReadPuzzle(file_name string) []string {
	raw_data, _ := os.ReadFile(file_name)
	data := strings.Fields(string(raw_data))
	return data
}
