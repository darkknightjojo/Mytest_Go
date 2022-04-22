package image

import (
	"fmt"
)

func CreateImage(name string, height int, weight int) MyImage {
	newImage := image{
		Height: height,
		Weight: weight,
		Name:   name,
	}
	return MyImage{myImage: newImage}
}

func DisplayImage(myImage MyImage) {
	fmt.Println("打印了一张照片:" + myImage.myImage.Name)
	fmt.Println("-----------------image end---------------")
}
