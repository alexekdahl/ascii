package ascii

import "image"

// scaleImage scales the input image to target dimensions.
// The scaling is performed using simple nearest-neighbor interpolation.
func scaleImage(img image.Image, targetWidth, targetHeight int) *image.RGBA {
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// Determine the scaling factors for width and height
	scaleFactorWidth := float64(targetWidth) / float64(originalWidth)
	scaleFactorHeight := float64(targetHeight) / float64(originalHeight)

	// Calculate new dimensions
	newWidth := int(float64(originalWidth) * scaleFactorWidth)
	newHeight := int(float64(originalHeight) * scaleFactorHeight)

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * originalWidth / newWidth
			srcY := y * originalHeight / newHeight
			newImg.Set(x, y, img.At(srcX, srcY))
		}
	}

	return newImg
}
