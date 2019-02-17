//This application takes in an image and decodes it into rgb values. These values are summed and averaged. The value that is highest (rgb) is then printed as the most occuring color. In addition, if all colors are close to each other, the console prints out that it may be a document since document's rgb values are typically very close to each other. An image up to 4k resolution can be used.

package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sort"
)

//struct containing image name, highest color, value of color
type imageInfo struct {
	name  string
	color string
	val   int
}

func main() {

	directory := "./images/"

	fmt.Println("Which color would you like to look for: \n1)Red color\n2)Blue color\n3)Green color\n4)Documents")

	var color int

	fmt.Scanf("%d", &color)

	var images []imageInfo

	//decifer input
	if color == 1 {
		images = parseDir(directory, "red")
	} else if color == 2 {
		images = parseDir(directory, "blue")
	} else if color == 3 {
		images = parseDir(directory, "green")
	} else if color == 4 {
		images = parseDir(directory, "doc")
	} else {
		fmt.Printf("Incorrect Input\nInput: %d", color)
		return
	}

	//output sorted array of images
	fmt.Println("\n\nIn order from highest color value to least")
	for _, image := range images {
		fmt.Println(image.name)
	}

}

//parses all files in the given directory
func parseDir(dirName string, desiredColor string) []imageInfo {

	dir, err := os.Open(dirName)

	if err != nil {
		fmt.Println(err)
	}

	files, err := dir.Readdir(-1)

	images := make([]imageInfo, 0)

	//check each file to be an image and have corect color
	for _, image := range files {
		imageName := image.Name()
		imageExt := imageName[len(imageName)-4:]
		if imageExt == ".png" || imageExt == ".jpg" {
			fmt.Printf("Scanning: %s\n", imageName)
			color, size := parseImage(dirName + imageName)

			if desiredColor == color {
				images = append(images, imageInfo{imageName, color, size})
			}

		}
	}

	//use the implemented sort interface
	sort.Sort(ByColorValue(images))

	return images
}

func parseImage(imageName string) (string, int) {

	//Take in a file
	imageFile, err := os.Open(imageName)

	if err != nil {
		fmt.Println(err)
	}

	//Get Image data and type of file
	imageData, _, err := image.Decode(imageFile)

	if err != nil {
		fmt.Println(err)
	}

	//get bounds
	bounds := imageData.Bounds()

	//first iteration
	r, g, b, _ := imageData.At(1, 1).RGBA()

	//converting uint32 to values 0-255
	redAvg := uint64(r >> 8)
	greenAvg := uint64(g >> 8)
	blueAvg := uint64(b >> 8)

	//iterate over all the data
	for x := bounds.Min.X + 1; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y + 1; y < bounds.Max.Y-8; y++ {

			//grab rgb values, convert them, and add them to the running total.
			r, g, b, _ = imageData.At(x, y).RGBA()
			redAvg += uint64(r >> 8)
			greenAvg += uint64(g >> 8)
			blueAvg += uint64(b >> 8)

		}
	}

	//fmt.Printf("Before division!\nRed: %d, Green: %d, Blue: %d\n", redAvg, greenAvg, blueAvg)

	//get iteration count and calculate average
	iterations := uint64((bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y))
	redAvg /= iterations
	greenAvg /= iterations
	blueAvg /= iterations

	//return which color occurs the most
	if isDoc(greenAvg, redAvg, blueAvg) {
		return "doc", 0
	}

	if redAvg >= greenAvg && redAvg >= blueAvg {

		return "red", int(redAvg)

	}
	if greenAvg >= redAvg && greenAvg >= blueAvg {

		return "green", int(greenAvg)

	}
	if blueAvg >= greenAvg && blueAvg >= redAvg {

		return "blue", int(blueAvg)

	}

	//or else return N/A and 0
	return "N/A", 0
}

//Documents are usually within 5 rgb values of each other color and have value > 170 meaning black text takes up almost half the white image.
func isDoc(greenAvg uint64, redAvg uint64, blueAvg uint64) bool {

	if int(greenAvg-blueAvg) <= 5 && int(greenAvg-redAvg) <= 5 && int(redAvg-blueAvg) <= 5 && greenAvg >= 170 {
		return true
	} else {
		return false
	}
}

//implementation of the sort interface
type ByColorValue []imageInfo

func (img ByColorValue) Len() int           { return len(img) }
func (img ByColorValue) Swap(i, j int)      { img[i], img[j] = img[j], img[i] }
func (img ByColorValue) Less(i, j int) bool { return img[i].val > img[j].val }
