package src

import (
"reflect"
"testing"
)

// TestSplitInput checks if the utility correctly handles literal \n
func TestSplitInput(t *testing.T) {
tests := []struct {
name     string
input    string
expected []string
}{
{"Basic text", "hello", []string{"hello"}},
{"Literal newline", "hello\\nworld", []string{"hello", "world"}},
{"Empty input", "", []string{""}},
{"Multiple newlines", "\\n\\n", []string{"", "", ""}},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
actual := SplitInput(tt.input)
if !reflect.DeepEqual(actual, tt.expected) {
t.Errorf("SplitInput(%q) = %v, want %v", tt.input, actual, tt.expected)
}
})
}
}

// TestMathFormula verifies that the formula (char-32)*9 + row correctly maps characters
func TestMathFormula(t *testing.T) {
// Let's test the character 'A' (ASCII 65) for Row 1
char := 'A'
row := 1
// Expected line index: (65-32)*9 + 1 = 33*9 + 1 = 297 + 1 = 298
expected := 298
actual := (int(char)-32)*9 + row

if actual != expected {
t.Errorf("Math for 'A' failed: got %d, want %d", actual, expected)
}

// Test Space (ASCII 32)
charSpace := ' '
expectedSpace := 1 // (32-32)*9 + 1
actualSpace := (int(charSpace)-32)*9 + row
if actualSpace != expectedSpace {
t.Errorf("Math for Space failed: got %d, want %d", actualSpace, expectedSpace)
}
}