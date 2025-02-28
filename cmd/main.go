package main

import (
	"fmt"
	"log-analyzer/internal/parser"
)

func main() {
	fmt.Println("Welcome to Log Analyzer")
	parser.TestReaderFile()
	parser.TestStatisticsCounter().DisplayStatistics()
}
