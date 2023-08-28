/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
)

func main() {
	sliceAndSaveImages("img.jpeg", 2, 2)
}

func sliceAndSaveImages(imagePath string, rows, cols int) {
	img, err := loadImage(imagePath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	width := img.Bounds().Dx() 	// Dx() returns the width of the image, img.Bounds() returns the rectangle that contains the image
	height := img.Bounds().Dy() // Dx and Dy are the width and height of the rectangle, respectively

	fmt.Println("image width", img.Bounds().Dx())
	fmt.Println("image height:", img.Bounds().Dy())

	sliceWidth := width / cols
	sliceHeight := height / rows

	// Create an "output" folder if it doesn't exist
	err = os.Mkdir("output", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating 'output' folder:", err)
		return
	}

//dc is a drawing context, which is used to draw images

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			dc := gg.NewContext(sliceWidth, sliceHeight) // NewContext(width, height) creates a new context with the specified width and height
			dc.DrawImage(img, -col*sliceWidth, -row*sliceHeight) // DrawImage(img, x, y) draws the image at the specified coordinates

			outputFilePath := filepath.Join("output", fmt.Sprintf("output_%d_%d.png", row, col))
			err := dc.SavePNG(outputFilePath)
			if err != nil {
				fmt.Println("Error saving image:", err)
			}
		}
	}
}

func loadImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Close the file when the function returns, defer is used to delay the execution of a function until the surrounding function returns

	img, _, err := image.Decode(file) // Decode decodes an image that has been encoded in a registered format
	if err != nil {
		return nil, err
	}

	return img, nil
}