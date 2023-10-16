package utils

import (
	"image"

	"gocv.io/x/gocv"
)

func RGBAImageToMat(img *image.RGBA) (gocv.Mat, error) {
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()
	// img.Pix contains the data in RGBA format, so we use gocv.MatTypeCV8UC4 (4 channels of uint8)
	mat, err := gocv.NewMatFromBytes(height, width, gocv.MatTypeCV8UC4, img.Pix)
	if err != nil {
		return gocv.Mat{}, err
	}
	
	gray := gocv.NewMat()
	gocv.CvtColor(mat, &gray, gocv.ColorBGRAToGray)

	return gray, nil
}
