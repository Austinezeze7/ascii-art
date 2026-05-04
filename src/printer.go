package src

import "fmt"

// PrintArt draws the ASCII representation row-by-row for a single word.
func PrintArt(word string, bannerLines []string) {
if word == "" {
return
}

// Each ASCII character is 8 lines tall.
for row := 1; row <= 8; row++ {
for _, char := range word {
// (char - 32) * 9 + row finds the exact line in your banner data
lineIndex := (int(char)-32)*9 + row

if lineIndex < len(bannerLines) {
fmt.Print(bannerLines[lineIndex])
}
}
// Finish the current horizontal row and move to the next vertical line
fmt.Println()
}
}
