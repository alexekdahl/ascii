package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
)

// ArtStyle defines the ASCII art style.
type ArtStyle int

const (
	ColorASCII ArtStyle = iota
	GrayASCII
)

// ASCIIOptions holds configuration for ASCII generation.
type ASCIIOptions struct {
	Style  ArtStyle
	Char   rune
	Width  int
	Height int
}

// Option is a function that applies a given configuration to ASCIIOptions.
type Option func(*ASCIIOptions)

// WithChar sets the ASCII character to be used.
func WithChar(c rune) Option {
	return func(opts *ASCIIOptions) {
		opts.Char = c
	}
}

// WithStyle sets the style of the ASCII art (color or grayscale).
func WithStyle(s ArtStyle) Option {
	return func(opts *ASCIIOptions) {
		opts.Style = s
	}
}

// WithWidth sets the desired width of the ASCII art.
func WithWidth(w int) Option {
	return func(opts *ASCIIOptions) {
		opts.Width = w
	}
}

// WithHeight sets the desired height of the ASCII art.
func WithHeight(h int) Option {
	return func(opts *ASCIIOptions) {
		opts.Height = h
	}
}

// defaultOptions returns the default ASCIIOptions.
func defaultOptions() ASCIIOptions {
	return ASCIIOptions{
		Style:  ColorASCII,
		Char:   '#',
		Width:  50, // Default width
		Height: 25, // Default height
	}
}

// applyOptions applies the given options to ASCIIOptions.
func applyOptions(opts ...Option) ASCIIOptions {
	config := defaultOptions()
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

// toASCIIString generates the ASCII representation based on the given options.
func toASCIIString(img *image.RGBA, options ...Option) string {
	config := applyOptions(options...)
	bounds := img.Bounds()

	colorConverter := getColorConverter(config.Style)
	var buffer bytes.Buffer

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			buffer.WriteString(colorConverter(rgba, config.Char))
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// getColorConverter returns the appropriate color conversion function based on the style.
func getColorConverter(style ArtStyle) func(color.RGBA, rune) string {
	if style == GrayASCII {
		return convertToGrayString
	}
	return convertToColorString
}

// convertToColorString converts a color.RGBA to a colored ASCII string.
func convertToColorString(rgba color.RGBA, char rune) string {
	r, g, b, _ := rgba.RGBA()
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", uint8(r>>8), uint8(g>>8), uint8(b>>8), char)
}

// convertToGrayString converts a color.RGBA to a grayscale ASCII string.
func convertToGrayString(rgba color.RGBA, char rune) string {
	r, g, b, _ := rgba.RGBA()
	gray := uint8((0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256)
	return fmt.Sprintf("\x1b[38;2;%[1]d;%[1]d;%[1]dm%c\x1b[0m", gray, char)
}

// ToASCII converts raw image data to an ASCII art string with a specified style.
// The image dimensions are scaled down to 'w' and 'h'.
// ToASCII converts raw image data to an ASCII art string with specified options.
func ToASCII(imgData []byte, opts ...Option) (string, error) {
	img, err := decodeImage(imgData)
	if err != nil {
		return "", err
	}

	config := applyOptions(opts...)
	scaledImg := scaleImage(img, config.Width, config.Height)
	ascii := toASCIIString(scaledImg, opts...)

	return ascii, nil
}
