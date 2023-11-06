package ascii

import (
	"image/color"
	"testing"
)

// TestConvertToColorString tests the convertToColorString function.
func TestConvertToColorString(t *testing.T) {
	rgba := color.RGBA{R: 255, G: 0, B: 0, A: 255} // Pure red
	char := '#'
	result := convertToColorString(rgba, char)
	expected := "\x1b[38;2;255;0;0m#\x1b[0m"
	if result != expected {
		t.Errorf("convertToColorString() = %v, want %v", result, expected)
	}
}

// TestConvertToGrayString tests the convertToGrayString function.
func TestConvertToGrayString(t *testing.T) {
	rgba := color.RGBA{R: 255, G: 0, B: 0, A: 255} // Pure red
	char := '#'
	result := convertToGrayString(rgba, char)
	// Expected value for grayscale may need to be adjusted depending on the grayscale formula used.
	expected := "\x1b[38;2;76;76;76m#\x1b[0m"
	if result != expected {
		t.Errorf("convertToGrayString() = %v, want %v", result, expected)
	}
}
