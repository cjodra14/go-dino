package main

import (
	"fmt"
	"go-dino/models"
	"go-dino/utils"
	"image/color"
	"time"

	// "time"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"

	"gocv.io/x/gocv"
)

func main() {
	playerIndex := 0
	enemyIndex := 0
	color := color.RGBA{}
	dinoWhite := models.NewObject("./objects/dino.png", "White Dino")
	dinoBlack := models.NewObject("./objects/dino_b.png", "Black Dino")
	dinos := []models.Object{
		*dinoWhite,
		*dinoBlack,
	}
	fmt.Println(dinos)

	bird := models.NewObject("./objects/bird.png", "Bird")
	birdBlack := models.NewObject("./objects/bird_b.png", "Black bird")
	cactus1 := models.NewObject("./objects/cact1.png", "cactus 1")
	cactus1Black := models.NewObject("./objects/cact1_b.png", "cactus 1 black")
	cactus2 := models.NewObject("./objects/cact2.png", "cactus 2")
	cactus2Black := models.NewObject("./objects/cact2_b.png", "cactus 2 black")

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

		for i, dino := range dinos {
			dino.FindObject(imgMat)
			if dino.IsFound {
				playerIndex = i
				setColor(playerIndex, &color)
				gocv.Rectangle(&imgMat, dino.Location, color, 5)
				fmt.Println(dino.Name, " at  ", dino.Location)
			}
		}

		for _, enemy := range enemies {
			enemy.FindObject(imgMat)
			if enemy.IsFound {
				fmt.Println(enemy.Name, " at  ", enemy.Location)
				gocv.Rectangle(&imgMat, enemy.Location, color, 2)
			}
		}

		window := gocv.NewWindow("Result")

		defer window.Close()

		window.IMShow(imgMat)

		if window.WaitKey(1) == 113 {
			logrus.Fatal("program exited")
		}

		fmt.Println("index", playerIndex)
		time.Sleep(time.Millisecond * 10)
	}

}

func setColor(playerIndex int, color *color.RGBA) {
	if playerIndex == 0 {
		color.B = 0
		color.R = 0
		color.G = 0
	}

	if playerIndex == 1 {
		color.B = 255
		color.R = 255
		color.G = 255
	}
}

func Jump() {
	robotgo.KeyTap("space")
}
