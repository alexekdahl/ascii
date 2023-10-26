package ascii

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	img := createTestImage()
	var imgBuffer bytes.Buffer
	err := png.Encode(&imgBuffer, img)
	if err != nil {
		t.Fatalf("Failed to encode test image: %v", err)
	}

	decodedImg, err := decodeImage(imgBuffer.Bytes())
	if err != nil {
		t.Fatalf("decodeImage failed: %v", err)
	}

	if decodedImg.Bounds().Dx() != img.Bounds().Dx() || decodedImg.Bounds().Dy() != img.Bounds().Dy() {
		t.Errorf("Decoded image dimensions do not match original image")
	}
}

func createTestImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	red := color.RGBA{255, 0, 0, 255}
	img.Set(0, 0, red)
	return img
}
