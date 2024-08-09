package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs.")

	var fruitList[10] string

	fruitList[0]="Apple"
	fruitList[1]="Tomato"
	fruitList[9]="Peach"   

	fmt.Println("Fruit list is: ",fruitList)   //simply prints space for index 2 to 8
	fmt.Println("Fruit list is: ",len(fruitList)) //return length of an array

	var vegList=[3] string{"Potato","beans","mushroom"}
	fmt.Println("Vegy list is: ",vegList)
	fmt.Println("Vegy list is: ",len(vegList))
}
