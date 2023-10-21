package models

import (
	"image"

	"gocv.io/x/gocv"
)

type Object struct {
	Name     string
	Width    int
	Height   int
	Image    gocv.Mat
	IsFound  bool
	Location image.Rectangle
}

func NewObject(imgPath, name string) *Object {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	return &Object{
		Name:   name,
		Image:  img,
		Height: img.Size()[0],
		Width:  img.Size()[1],
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

	location := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+object.Image.Cols(), maxLoc.Y+object.Image.Rows())

	if maxVal > 0.8 {
		object.Location = location
		object.IsFound = true
		return true
	}

	object.IsFound = false

	return false

}
