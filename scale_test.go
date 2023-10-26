package ascii

import (
	"image"
	"image/color"
	"testing"
)

func TestScaleImage(t *testing.T) {
	// Initialize a 2x2 test image
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	img.Set(0, 0, red)
	img.Set(1, 1, blue)

	targetWidth, targetHeight := 4, 4
	scaledImg := scaleImage(img, targetWidth, targetHeight)

	if scaledImg.Bounds().Dx() != targetWidth || scaledImg.Bounds().Dy() != targetHeight {
		t.Errorf("Scaled image dimensions do not match target dimensions")
	}

	// Verify pixel colors (assuming nearest-neighbor interpolation)
	if scaledImg.At(0, 0) != red || scaledImg.At(3, 3) != blue {
		t.Errorf("Scaled image colors do not match original image")
	}
}
