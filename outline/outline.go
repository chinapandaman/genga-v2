package outline

import (
	"fmt"
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

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

	x, y := 10, 20
	r16, g16, b16, _ := img.At(x, y).RGBA()

	// 5. Convert 16-bit channels to standard 8-bit values (0-255)
	r8 := uint8(r16 >> 8)
	g8 := uint8(g16 >> 8)
	b8 := uint8(b16 >> 8)
	fmt.Println(r8, g8, b8)
}
