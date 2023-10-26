package ascii

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// decodeImage decodes raw image data to an image.Image object.
func decodeImage(imgData []byte) (image.Image, error) {
	img, format, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}
	if format != "jpeg" && format != "png" && format != "gif" {
		return nil, errors.New("unsupported image format")
	}
	return img, nil
}
