package src

import "strings"

// SplitInput handles literal "\n" strings from the terminal
func SplitInput(input string) []string {
if input == "" {
return []string{""}
}
// Replaces literal \n with a real newline and splits into a slice
return strings.Split(strings.ReplaceAll(input, "\\n", "\n"), "\n")
}
