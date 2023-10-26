package ascii

import (
	"image"
	"image/color"
	"testing"
)

func TestColorASCIIString(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	col := color.RGBA{255, 0, 0, 255}
	img.Set(0, 0, col)
	img.Set(1, 1, col)

	result := colorASCIIString(img)

	if result == "" {
		t.Errorf("Expected a non-empty string, got an empty string")
	}
}
