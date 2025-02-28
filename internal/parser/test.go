package parser

import (
	"log"
	"log-analyzer/internal/stats"
	"log-analyzer/pkg/fileutils"
)

func TestStatisticsCounter() *stats.Statistics {
	filename := "test.txt"
	lines, err := fileutils.ReadLinesInRange(filename)
	if err != nil {
		log.Fatal(err)
	}

	stat := stats.NewStatistic()

	for _, line := range lines {
		ip, _, request, code, err := ApachParseLine(line)
		if err != nil {
			log.Println(err)
			continue
		}
		stat.AddStatistics(ip, code, request)
	}
	return stat
}
