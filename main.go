package main

import (
	"fmt"
)

type Dimensions struct {
	length int
	width int
	height int
}

func (d Dimensions) Area() int {
	return d.height * d.width
}

func (d Dimensions) Volume() int {
	return d.height * d.width * d.length
}

// func dimensions(length, width, height int) (area, volume int) {
// 	area = length * height
// 	volume = length * height * width
// 	return
// }

func main(){
	// area, volume := dimensions(5, 5, 5)
	// fmt.Println("area: ", area, ", volume: ", volume)

	d := Dimensions{10, 5, 6}
	fmt.Println("Area: ",d.Area())
	fmt.Println("Volume: ",d.Volume())
}