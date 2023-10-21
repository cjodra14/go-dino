package main

import (
	"fmt"
	"go-dino/models"
	"go-dino/utils"
	"image/color"

	// "time"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"

	"gocv.io/x/gocv"
)

func main() {
	playerIndex := 0
	enemyIndex := 0
	color := color.RGBA{}
	dinoWhite := models.NewObject("./objects/dino.png", false)
	dinoBlack := models.NewObject("./objects/dino_b.png", true)
	dinos := []models.Object{
		*dinoWhite,
		*dinoBlack,
	}
	fmt.Println(dinos)

	bird := models.NewObject("./objects/bird.png", false)
	birdBlack := models.NewObject("./objects/bird_b.png", true)
	cactus1 := models.NewObject("./objects/cact1.png", false)
	cactus1Black := models.NewObject("./objects/cact1_b.png", true)
	cactus2 := models.NewObject("./objects/cact2.png", false)
	cactus2Black := models.NewObject("./objects/cact2_b.png", true)

	enemies := []models.Object{
		*bird,
		*birdBlack,
		*cactus1,
		*cactus1Black,
		*cactus2,
		*cactus2Black,
	}
	fmt.Println(enemies)

	fmt.Println(playerIndex, " -- ", enemyIndex)

	// ticker := time.NewTicker(1 * time.Second) // Create a new ticker that ticks every 2 seconds

	// go func() { // Start a new Goroutine
	// 	for range ticker.C { // Loop indefinitely
	// 		Jump() // Call the Jump function
	// 	}
	// }()

	for {

		screen, err := utils.GetDinoBoardScreen()
		if err != nil {
			logrus.Error(err)
		}

		imgMat, err := utils.RGBAImageToMat(screen)
		if err != nil {
			logrus.Error(err)
		}

		isWhiteDino := dinos[0].FindObject(imgMat)
		isBlackDino := dinos[1].FindObject(imgMat)

		if isWhiteDino {
			playerIndex = 0
			color.B = 0
			color.R = 0
			color.G = 0
		}

		if isBlackDino {
			playerIndex = 1
			color.B = 255
			color.R = 255
			color.G = 255
		}

		gocv.Rectangle(&imgMat, dinos[playerIndex].Location, color, 2)
		window := gocv.NewWindow("Result")

		defer window.Close()

		window.IMShow(imgMat)

		if window.WaitKey(1) == 113 {
			logrus.Fatal("program exited")
		}

		fmt.Println("index", playerIndex)
		fmt.Println("is white", isWhiteDino)
		fmt.Println("is black", isBlackDino)
	}

}

func Jump() {
	robotgo.KeyTap("space")
}
