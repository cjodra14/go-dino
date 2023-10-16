package main

import (
	"fmt"
	"go-dino/models"
	"go-dino/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	playerIndex := 0
	enemyIndex := 0
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

	for i := 0; i < 1000000; i++ {
		screen, err := utils.GetDinoBoardScreen()
		if err != nil {
			logrus.Error(err)
		}

		imgMat, err := utils.RGBAImageToMat(screen)
		if err != nil {
			logrus.Error(err)
		}

		isWhiteDino := dinoWhite.FindObject(imgMat)
		fmt.Println(isWhiteDino)
		fmt.Println(dinoWhite.Location)

	}

}
