package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"os"
)

type calculateMeanAverageColourWithDeltaResult struct {
	averageRed   uint8
	averageGreen uint8
	averageBlue  uint8

	deltaRed   uint8
	deltaGreen uint8
	deltaBlue  uint8
}

func calculateMeanAverageColourWithDelta(img image.Image) (result calculateMeanAverageColourWithDeltaResult) {
	imgSize := img.Bounds().Size()

	var redSum float64
	var greenSum float64
	var blueSum float64

	for x := 0; x <= imgSize.X; x++ {
		for y := 0; y <= imgSize.Y; y++ {
			pixel := img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)

			redSum += float64(col.R)
			greenSum += float64(col.G)
			blueSum += float64(col.B)
		}
	}

	imgArea := float64(imgSize.X * imgSize.Y)

	result.averageRed = uint8(math.Round(redSum / imgArea))
	result.averageGreen = uint8(math.Round(greenSum / imgArea))
	result.averageBlue = uint8(math.Round(blueSum / imgArea))

	redSum = 0
	greenSum = 0
	blueSum = 0

	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			pixel := img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)

			redSum += math.Abs(float64(result.averageRed) - float64(col.R))
			greenSum += math.Abs(float64(result.averageGreen) - float64(col.G))
			blueSum += math.Abs(float64(result.averageBlue) - float64(col.B))
		}
	}

	result.deltaRed = uint8(math.Round(redSum / imgArea))
	result.deltaGreen = uint8(math.Round(greenSum / imgArea))
	result.deltaBlue = uint8(math.Round(blueSum / imgArea))

	return
}

func main() {
	file, err := os.Open("12.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	result := calculateMeanAverageColourWithDelta(img)

	fmt.Printf(
		"rgb(%d, %d, %d)",
		result.averageRed,
		result.averageGreen,
		result.averageBlue,
	)
}