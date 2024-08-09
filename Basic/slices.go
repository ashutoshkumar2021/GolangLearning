package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on slices")

	var fruitList=[]string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitList is %T\n",fruitList)
	fmt.Println(fruitList);

	fruitList=append(fruitList,"Mango","Banana")
	fmt.Println(fruitList);

	fruitList=append(fruitList[1:3])
	fmt.Println(fruitList);
}
