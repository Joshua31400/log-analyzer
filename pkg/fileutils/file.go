package fileutils

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadLinesInRange(fileName string) ([]string, error) {
	/*
		This function reads the lines of a file and returns them as a slice of strings.
	*/
	filePath := filepath.Join("data", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
