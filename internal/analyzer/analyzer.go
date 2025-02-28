package analyzer

import (
	"fmt"
	"log-analyzer/internal/parser"
	"log-analyzer/reports"
)

func Run() {
	fmt.Println("Welcome to Log Analyzer")
	stats := parser.TestStatisticsCounter()

	stats.DisplayStatistics()

	title := "Log Statistics Graph"
	filePath := "log_statistics_graph.png"
	reports.GenerateGraph(stats.HTTPCodeCount, title, filePath)
}
