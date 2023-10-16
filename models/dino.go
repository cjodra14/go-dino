package models

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

type Object struct {
	Width     int
	Height    int
	Image     gocv.Mat
	DarkTheme bool
	Location  image.Rectangle
}

func NewObject(imgPath string, isDarkTheme bool) *Object {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	return &Object{
		Image:     img,
		Height:    img.Size()[0],
		Width:     img.Size()[1],
		DarkTheme: isDarkTheme,
	}
}

func (object *Object) FindObject(img gocv.Mat) bool {
	resultRows := img.Rows() - object.Image.Rows() + 1
	resultCols := img.Cols() - object.Image.Cols() + 1
	result := gocv.NewMatWithSize(resultRows, resultCols, gocv.MatTypeCV32F)
	defer result.Close()

	method := gocv.TmCcorrNormed
	gocv.MatchTemplate(img, object.Image, &result, method, gocv.NewMat())

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	fmt.Printf("Best match found at: %v\n", maxLoc)

	location := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+object.Image.Cols(), maxLoc.Y+object.Image.Rows())
	
	if maxVal > 0.8 {
		object.Location = location
		return true
	}

	return false

}
