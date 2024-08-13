package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance in golang; No super or parent

	Ashutosh := User{"Ashutosh", "ashu@gmail.co", true, 20}
	fmt.Println(Ashutosh)
	fmt.Printf("Ashutosh Details are: %+v\n",Ashutosh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
