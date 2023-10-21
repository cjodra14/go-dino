package utils

import (
	"image"

	"github.com/kbinani/screenshot"
)

func GetDinoBoardScreen() (*image.RGBA, error) {
	bounds := image.Rect(700, 100, 1300, 300)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return &image.RGBA{}, err
	}

	return img, nil
}
