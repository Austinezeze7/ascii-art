package main

import (
	"fmt"
	"os"

	"ascii-art/src"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <string>")
		os.Exit(1)
	}

	input := os.Args[1]

	if input == "" {
		return
	}

	bannerLines, err := src.ReadBanner("banners/standard.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not load banner: %v\n", err)
		os.Exit(1)
	}

	lines := src.SplitInput(input)

	for i, line := range lines {
		if line == "" {
			// Skip trailing empty segment from a trailing \n
			if i == len(lines)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		src.PrintArt(line, bannerLines)
	}
}