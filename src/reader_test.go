package src

import (
	"testing"
	"fmt"
)

// In Go tests, we use TestName(t *testing.T) instead of func main()
func TestReadBannerLogic(t *testing.T) {
	// We use ../ because the test runs inside the src folder, 
	// so it needs to go up one level to find the banners folder.
	lines, err := ReadBanner("../banners/standard.txt") 
	if err != nil {
		t.Fatalf("❌ ERROR: Could not read file: %v", err)
	}

	fmt.Printf("✅ Total lines read: %d\n", len(lines))

	if len(lines) < 855 {
		t.Errorf("❌ ERROR: File too short! Got %d lines", len(lines))
	}
}
