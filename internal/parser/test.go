package parser

import (
	"fmt"
	"log"
	"log-analyzer/pkg/fileutils"
)

func TestFile() {
	filename := "test.txt"
	lines, err := fileutils.ReadLinesInRange(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
