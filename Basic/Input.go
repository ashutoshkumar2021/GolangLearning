package main

import (
	"fmt"
	"os"
	"bufio"   //this helps to take input
)
func main(){
	welcome :="Welcome to goLang Journey"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)  //this line for taking input
	fmt.Println("Enter the rating for our Pizza:")
    
	// comma ok // error ok
    input, _ :=reader.ReadString('\n')   // in comma ok, it works like try and error block so, we have to write in this type syntax
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T, ", input)
}
