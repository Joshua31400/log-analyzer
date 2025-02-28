package parser

import (
	"fmt"
	"regexp"
)

func ApachParseLine(log string) (string, string, string, string, error) {
	/*
		This function parses an Apache log line and returns the IP address, date, request, and status code.
	*/
	re := regexp.MustCompile(`^(\S+) - - \[([^\]]+)\] "(\S+) (\S+) HTTP/\S+" (\d{3}) \d+`)
	matches := re.FindStringSubmatch(log)

	if len(matches) == 0 {
		fmt.Println("Log dont fit with the file format:", log)
		return "", "", "", "", fmt.Errorf("Invalid apache log format")
	}

	ip := matches[1]
	date := matches[2]
	request := matches[3] + " " + matches[4]
	code := matches[5]

	return ip, date, request, code, nil
}
