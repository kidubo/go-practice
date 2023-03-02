package main

import (
	"fmt"
)


// Receiver Methods 
type Dimensions struct {
	length int
	width int
	height int
}

// Reciver Methods
func (d *Dimensions) Area() int {
	// Struct pointer
	d.height = 8
	return d.height * d.width
}

// Reciver Methods
func (d Dimensions) Volume() int {
	d.height = 6
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
	fmt.Println("Dimensions",d)
	fmt.Println("Volume: ",d.Volume())
	fmt.Println("Dimensions",d)


	// x, y := 5, 10
	// n := &x
	// fmt.Println(*n, y)
	// // Using equal b'cs we are setting the value not initialize the value
	// *n = 50  
	// fmt.Println(x, y)

  // t := &y
	// *t = y + 30
	// fmt.Println("address:", t, "value:", *t)
}