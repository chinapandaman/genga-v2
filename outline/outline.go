package outline

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"math"
	"os"
)

func getEightBitRGB(r16 uint32, g16 uint32, b16 uint32) (uint8, uint8, uint8) {
	r8 := uint8(r16 >> 8)
	g8 := uint8(g16 >> 8)
	b8 := uint8(b16 >> 8)

	return r8, g8, b8
}

func OutlineImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newBounds := image.Rect(0, 0, width, height)
	newImg := image.NewRGBA(newBounds)

	for x := 1; x <= width; x++ {
		for y := 1; y <= height; y++ {
			r16, g16, b16, _ := img.At(x, y).RGBA()
			r8, g8, b8 := getEightBitRGB(r16, g16, b16)

			r16, g16, b16, _ = img.At(x-1, y-1).RGBA()
			r82, g82, b82 := getEightBitRGB(r16, g16, b16)

			diff := math.Sqrt(float64((r8-r82)*(r8-r82) + (g8-g82)*(g8-g82) + (b8-b82)*(b8-b82)))
			if diff > 5 {
				black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
				newImg.Set(x, y, black)
			}
		}
	}

	output, err := os.Create("./images/output.png")
	if err != nil {
		log.Fatalf("Failed to create outlined image file: %v", err)
	}
	defer output.Close()

	if err := png.Encode(output, newImg); err != nil {
		log.Fatalf("Failed to save outlined image: %v", err)
	}
}
