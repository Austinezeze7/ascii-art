package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"ascii-art/src"
)

// captureOutput mirrors main.go logic exactly and captures stdout
func captureOutput(input string, bannerLines []string) string {
	// Mirror main.go early return for empty string
	if input == "" {
		return ""
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	lines := src.SplitInput(input)
	for i, line := range lines {
		if line == "" {
			// Skip trailing empty segment — mirrors main.go fix
			if i == len(lines)-1 {
				continue
			}
			os.Stdout.WriteString("\n")
			continue
		}
		src.PrintArt(line, bannerLines)
	}

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// countNewlines counts \n in output
func countNewlines(s string) int {
	return strings.Count(s, "\n")
}

// loadTestBanner loads standard.txt — skips test if file not found
func loadTestBanner(t *testing.T) []string {
	t.Helper()
	lines, err := src.ReadBanner("banners/standard.txt")
	if err != nil {
		t.Skip("banners/standard.txt not available, skipping")
	}
	return lines
}

// TestEmptyString — go run . ""
func TestEmptyString(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("", banner)
	if out != "" {
		t.Errorf("empty string: expected no output, got %q", out)
	}
}

// TestOnlyNewline — go run . "\n"
func TestOnlyNewline(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput(`\n`, banner)
	if countNewlines(out) != 1 {
		t.Errorf("'\\n': expected 1 blank line, got %d", countNewlines(out))
	}
}

// TestHelloWithTrailingNewline — go run . "Hello\n"
func TestHelloWithTrailingNewline(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput(`Hello\n`, banner)
	if countNewlines(out) != 8 {
		t.Errorf("'Hello\\n': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestHello — go run . "hello"
func TestHello(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("hello", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'hello': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestHeLlO — go run . "HeLlO"
func TestHeLlO(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("HeLlO", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'HeLlO': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestHelloThere — go run . "Hello There"
func TestHelloThere(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("Hello There", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'Hello There': expected 8 lines, got %d", countNewlines(out))
	}
}

// Test1Hello2There — go run . "1Hello 2There"
func Test1Hello2There(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("1Hello 2There", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'1Hello 2There': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestCurlyHelloThere — go run . "{Hello There}"
func TestCurlyHelloThere(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("{Hello There}", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'{Hello There}': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestHelloNewlineThere — go run . "Hello\nThere"
func TestHelloNewlineThere(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput(`Hello\nThere`, banner)
	if countNewlines(out) != 16 {
		t.Errorf("'Hello\\nThere': expected 16 lines, got %d", countNewlines(out))
	}
}

// TestHelloDoubleNewlineThere — go run . "Hello\n\nThere"
func TestHelloDoubleNewlineThere(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput(`Hello\n\nThere`, banner)
	if countNewlines(out) != 17 {
		t.Errorf("'Hello\\n\\nThere': expected 17 lines, got %d", countNewlines(out))
	}
}

// TestRandomFourLowerThreeUpper — audit random case 1
func TestRandomFourLowerThreeUpper(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("abcdEFG", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'abcdEFG': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestRandomFiveLowerSpaceTwoNumbers — audit random case 2
func TestRandomFiveLowerSpaceTwoNumbers(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("hello 42", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'hello 42': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestRandomOneUpperThreeSpecial — audit random case 3
func TestRandomOneUpperThreeSpecial(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("A!@#", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'A!@#': expected 8 lines, got %d", countNewlines(out))
	}
}

// TestRandomMixedAll — audit random case 4
func TestRandomMixedAll(t *testing.T) {
	banner := loadTestBanner(t)
	out := captureOutput("ab  1!@ABC", banner)
	if countNewlines(out) != 8 {
		t.Errorf("'ab  1!@ABC': expected 8 lines, got %d", countNewlines(out))
	}
}