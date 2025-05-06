package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/truetype"
)

func main() {
	// Load the background image
	bgImgFile, err := os.Open("background.jpg") // Replace with your image path
	if err != nil {
		log.Fatal(err)
	}
	defer bgImgFile.Close()

	bgImg, _, err := image.Decode(bgImgFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new image with the same dimensions as the background
	upImg := image.NewRGBA(bgImg.Bounds())
	draw.Draw(upImg, upImg.Bounds(), &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	draw.Draw(upImg, upImg.Bounds(), bgImg, image.ZP, draw.Over)

	// Load the font
	fontFile, err := os.Open("font.ttf") // Replace with your font path
	if err != nil {
		log.Fatal(err)
	}
	defer fontFile.Close()

	fontBytes, err := os.ReadFile(fontFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Create a drawing context
	pointSize := fixed.I(36) // Font size
	face := truetype.NewFace(f, &truetype.Options{Size: float64(pointSize)})
	drawer := &font.Drawer{
		Dst:  upImg,
		Src:  image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}), // Text color
		Face: face,
	}

	// Draw the text
	x := fixed.I(100)
	y := fixed.I(100)
	drawer.DrawString("Hello, World!", fixed.P(x, y))

	// Save the resulting image
	outFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, upImg, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image with text created successfully!")
}
