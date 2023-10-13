package main // this program takes a long time to run because it goes through every pixel.

import (
	"fmt"

	"github.com/bglawson1001/img_module"
)

func main() {
	imageURL := "https://hgtvhome.sndimg.com/content/dam/images/hgtv/fullset/2011/7/18/0/HGTV_Color-Wheel-Full_s4x3.jpg.rend.hgtvcom.1280.1280.suffix/1400967008479.jpeg" // Replace with your actual image URL
	fileName := "downloaded_image.jpg"                                                                                                                                    // Specify the desired file name

	//will download a image file using the url it has. 
	err := img_module.Getpic(imageURL, fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	txt := "Color_Wheel"
	outputDestination := "hello.jpg" // I could not get this part to work with this module. 
	// I made a seperate module where I made Color_Wheel with Griffin's default starter code. 

	//is supposed to print out colored text to an output file. 
	err = img_module.Text(txt, outputDestination)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image saved as", outputDestination)

	err = img_module.Colors(outputDestination)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	output2 := "Gray_image.jpg"

	err = img_module.Gray_scale(fileName, output2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
