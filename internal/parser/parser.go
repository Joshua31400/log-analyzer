package parser

import "strings"

func SplitLines(content string) []string {
	/*
		This function splits the content into lines and returns a slice of lines.
	*/
	return strings.Split(strings.TrimSpace(content), "\n")
}
