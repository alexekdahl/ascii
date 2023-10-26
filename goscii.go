package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
)

// colorASCIIString converts an RGBA image to its ASCII representation.
// Each pixel is converted to an ASCII character with color codes.
// The function assumes the image to be in RGBA format.
func colorASCIIString(img *image.RGBA) string {
	var buffer bytes.Buffer
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y))
			rgba := c.(color.RGBA)
			r, g, b, _ := rgba.RGBA()
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			buffer.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm#\x1b[0m", r8, g8, b8))
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// ToASCII converts raw image data to an ASCII art string.
// The image dimensions are scaled down to 'w' and 'h'.
func ToASCII(imgData []byte, w, h int) (string, error) {
	img, err := decodeImage(imgData)
	if err != nil {
		return "", err
	}

	scaledImg := scaleImage(img, w, h)
	ascii := colorASCIIString(scaledImg)

	return ascii, nil
}
