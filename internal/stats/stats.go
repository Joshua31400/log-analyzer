package stats

import (
	"fmt"
)

type Statistics struct {
	/*
		This struct contains the statistics of the log file.
	*/
	IPCount       map[string]int
	HTTPCodeCount map[string]int
	RequestCount  map[string]int
}

func NewStatistic() *Statistics {
	/*
		This function creates a new Statistics object and initializes the maps.
	*/
	return &Statistics{
		IPCount:       make(map[string]int),
		HTTPCodeCount: make(map[string]int),
		RequestCount:  make(map[string]int),
	}
}

func (s *Statistics) AddStatistics(ip, httpCode, request string) {
	/*
		This function adds the statistics of the log file.
	*/
	s.IPCount[ip]++
	s.HTTPCodeCount[httpCode]++
	s.RequestCount[request]++
}

func (s *Statistics) DisplayStatistics() {
	/*
		This function displays the statistics of the log file.
	*/
	fmt.Println("Log Statistics:")

	fmt.Println("\nOccurrences of IP addresses:")
	for ip, count := range s.IPCount {
		fmt.Printf("%s : %d\n", ip, count)
	}

	fmt.Println("\nOccurrences of HTTP codes:")
	for code, count := range s.HTTPCodeCount {
		fmt.Printf("%s : %d\n", code, count)
	}

	fmt.Println("\nOccurrences of HTTP requests:")
	for request, count := range s.RequestCount {
		fmt.Printf("%s : %d\n", request, count)
	}
}
